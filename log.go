package minelog

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type LogLevel int

const (
	LogInfo LogLevel = iota
	LogWarn
)

type LogLine struct {
	Time      time.Time
	Level     LogLevel
	Line      string
	whitelist *Whitelist
}

func NewLogLine(filename string, wl *Whitelist) (*LogLine, error) {
	ts, err := getTimeFromFilename(filename)
	if err != nil {
		return nil, err
	}
	l := &LogLine{
		Time:      ts,
		whitelist: wl,
	}
	return l, nil
}

func getTimeFromFilename(filename string) (time.Time, error) {
	// testdata/2019-09-18-1.log.gz
	// testdata/latest.log
	filename = filepath.Base(filename)
	if filename == "latest.log" {
		return time.Now().In(time.UTC), nil
	}
	e := strings.Index(filename, ".")
	filetime, err := time.Parse("2006-01-02", filename[:e-2])
	if err != nil {
		return filetime, err
	}
	return filetime, nil
}

func getPlayerFromLine(line string) string {
	toks := strings.Fields(line)
	return toks[0]
}

type DeathType int

func (d DeathType) String() string {
	switch d {
	case DeathNone:
		return "None"
	case DeathFire:
		return "Fire"
	case DeathLava:
		return "Lava"
	case DeathInWall:
		return "InWall"
	case DeathDrown:
		return "Drown"
	case DeathStarve:
		return "Starve"
	case DeathCactus:
		return "Cactus"
	case DeathFall:
		return "Fall"
	case DeathOutOfWorld:
		return "OutOfWorld"
	case DeathGeneric:
		return "Generic"
	case DeathExplosion:
		return "Explosion"
	case DeathMagic:
		return "Magic"
	case DeathMob:
		return "Mob"
	case DeathPlayer:
		return "Player"
	case DeathArrow:
		return "Arrow"
	case DeathFireball:
		return "Fireball"
	case DeathThrown:
		return "Thrown"
	case DeathIndirectMagic:
		return "IndirectMagic"
	case DeathSquish:
		return "Squish"
	}
	return "None"
}

const (
	DeathNone DeathType = iota
	DeathFire
	DeathLava
	DeathInWall
	DeathDrown
	DeathStarve
	DeathCactus
	DeathFall
	DeathOutOfWorld
	DeathGeneric
	DeathExplosion
	DeathMagic
	DeathMob
	DeathPlayer
	DeathSquish
	DeathArrow
	DeathFireball
	DeathThrown
	DeathIndirectMagic
)

type Death struct {
	Type    DeathType
	Victim  string
	Killer  string
	LogLine LogLine
}

func (d Death) String() string {
	s := fmt.Sprintf("%s %s", d.Victim, d.Type)
	if d.Killer != "" {
		s += fmt.Sprintf(" by %s", d.Killer)
	}
	return s
}

// TODO Needs tweaks, these are from buckit server, and don't align with offical server.
func (l *LogLine) GetDeath() Death {
	toks := strings.Fields(l.Line)
	d := Death{LogLine: *l}
	if !l.whitelist.ContainsPlayer(toks[0]) {
		d.Type = DeathNone
		return d
	}
	// TODO found a complete list of deaths: https://minecraft.gamepedia.com/Death_messages
	// TODO break this up by death category, perhaps via functions.
	if strings.Contains(l.Line, "went up in flames") {
		d.Type = DeathFire
		d.Victim = toks[0]
	} else if strings.Contains(l.Line, "burned to death") {
		d.Type = DeathFire
		d.Victim = toks[0]
	} else if strings.Contains(l.Line, "was burnt to a crisp whilst fighting") {
		d.Type = DeathFire
		d.Victim = toks[0]
		d.Killer = toks[len(toks)-1]
	} else if strings.Contains(l.Line, "hit the ground too hard whilst trying to escape") {
		d.Type = DeathFall
		d.Victim = toks[0]
		d.Killer = toks[len(toks)-1]
	} else if strings.Contains(l.Line, "hit the ground too hard") {
		d.Type = DeathFall
		d.Victim = toks[0]
	} else if strings.Contains(l.Line, "fell from a high place") {
		d.Type = DeathFall
		d.Victim = toks[0]
	} else if strings.Contains(l.Line, "tried to swim in lava") {
		d.Type = DeathLava
		d.Victim = toks[0]
	} else if strings.Contains(l.Line, "discovered the floor was lava") {
		d.Type = DeathLava
		d.Victim = toks[0]
	} else if strings.Contains(l.Line, "suffocated in a wall") {
		d.Type = DeathInWall
		d.Victim = toks[0]
	} else if strings.Contains(l.Line, "drowned") {
		d.Type = DeathDrown
		d.Victim = toks[0]
	} else if strings.Contains(l.Line, "starved to death") {
		d.Type = DeathStarve
		d.Victim = toks[0]
	} else if strings.Contains(l.Line, "was pricked to death") {
		d.Type = DeathCactus
		d.Victim = toks[0]
	} else if strings.Contains(l.Line, "fell out of the world") {
		d.Type = DeathOutOfWorld
		d.Victim = toks[0]
	} else if strings.Contains(l.Line, "died") {
		d.Type = DeathGeneric
		d.Victim = toks[0]
	} else if strings.Contains(l.Line, "blew up") {
		d.Type = DeathExplosion
		d.Victim = toks[0]
	} else if strings.Contains(l.Line, "blown up by") {
		d.Type = DeathExplosion
		d.Victim = toks[0]
		d.Killer = toks[len(toks)-1]
	} else if strings.Contains(l.Line, "was squished too much") {
		d.Type = DeathSquish
		d.Victim = toks[0]
	} else if strings.Contains(l.Line, "was killed by magic") {
		d.Type = DeathMagic
		d.Victim = toks[0]
	} else if strings.Contains(l.Line, "was slain by") {
		d.Type = DeathMob
		d.Victim = toks[0]
		d.Killer = toks[len(toks)-1]
	} else if strings.Contains(l.Line, "was shot by") {
		d.Type = DeathArrow
		d.Victim = toks[0]
		d.Killer = toks[len(toks)-1]
	} else if strings.Contains(l.Line, "was fireballed by") {
		d.Type = DeathFireball
		d.Victim = toks[0]
		d.Killer = toks[len(toks)-1]
	} else if strings.Contains(l.Line, "was pummeled by") {
		d.Type = DeathThrown
		d.Victim = toks[0]
		d.Killer = toks[len(toks)-1]
		// ricky_ninja was killed by Guardian using magic
	} else if strings.Contains(l.Line, "was killed by") {
		d.Type = DeathIndirectMagic
		d.Victim = toks[0]
		d.Killer = toks[4] //toks[len(toks)-1]
	} else {
		d.Type = DeathNone
	}
	return d
}

func (l *LogLine) Parse(line string) error {
	// [07:03:55] [Server thread/INFO]: ricky_ninja fell from a high place
	parts := strings.SplitAfterN(line, "]: ", 2)
	if len(parts) != 2 {
		return fmt.Errorf("can't parse: %s", line)
	}
	timeloglevel := strings.SplitAfterN(parts[0], " ", 2)
	if len(timeloglevel) != 2 {
		return fmt.Errorf("can't parse time: %s", parts[0])
	}
	timestr := timeloglevel[0]
	timestr = strings.TrimPrefix(timestr, "[")
	timestr = strings.TrimSuffix(timestr, "] ")
	ts, err := time.Parse("15:04:05", timestr)
	if err != nil {
		return err
	}
	if strings.Contains(timeloglevel[1], "INFO") {
		l.Level = LogInfo
	} else if strings.Contains(timeloglevel[1], "WARN") {
		l.Level = LogWarn
	}
	l.Line = parts[1]
	l.Time = time.Date(l.Time.Year(), l.Time.Month(), l.Time.Day(), ts.Hour(), ts.Minute(), ts.Second(), ts.Nanosecond(), time.UTC)
	return nil
}

func KnownEvent(line string) bool {
	known := []string{"joined", "left", "lost connection", "moved too quickly!",
		"has made the advancement", "completed the challenge", "moved wrongly"}
	for _, k := range known {
		if strings.Contains(line, k) {
			return true
		}
	}
	return false
}

func GetTextReader(filename string) (*textproto.Reader, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	var r io.Reader
	r = fd
	if strings.HasSuffix(filename, ".gz") {
		gr, err := gzip.NewReader(fd)
		if err != nil {
			return nil, err
		}
		r = gr
	}
	tr := textproto.NewReader(bufio.NewReader(r))
	return tr, nil
}
