package main

import (
	"bufio"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"

	"github.com/wcharczuk/go-chart"
)

var (
	logdir    string
	whitelist string
)

func init() {
	flag.StringVar(&logdir, "logdir", "", "the directory with log files")
	flag.StringVar(&whitelist, "whitelist", "", "the whitelist.json file")
}

func main() {
	flag.Parse()
	err := validateFlags()
	if err != nil {
		log.Fatal(err)
	}
	players, err := GetPlayers(whitelist)
	if err != nil {
		log.Fatal(err)
	}
	deaths, err := processLogDir(logdir, players)
	if err != nil {
		log.Fatal(err)
	}
	deathCountByVictim := make(map[string]int)
	deathTypeByVictim := make(map[string]map[string]int)
	for _, death := range deaths {
		deathCountByVictim[death.Victim]++
		deathType, ok := deathTypeByVictim[death.Victim]
		if !ok {
			deathType = make(map[string]int)
			deathTypeByVictim[death.Victim] = deathType
		}
		if death.Type == DeathMob {
			deathType[death.Killer]++
		} else {
			deathType[death.Type.String()]++
		}
		//fmt.Printf("%s -> %s\n", death, death.LogLine.Line)
	}
	whoDiedMost := mostDeaths(deathCountByVictim)
	whoDiedMostWays := mostVaried(deathTypeByVictim)
	fmt.Printf("%s has died more than anyone else with %d deaths\n", whoDiedMost, deathCountByVictim[whoDiedMost])
	dtypes := deathTypeByVictim[whoDiedMostWays]
	ways := make([]string, 0)
	pie := chart.PieChart{
		Width:  512,
		Height: 512,
	}
	for k, i := range dtypes {
		pie.Values = append(pie.Values, chart.Value{Value: float64(i), Label: k})
		ways = append(ways, fmt.Sprintf("%s: %d", k, i))
	}
	fmt.Printf("%s has died more ways than anyone else: %s\n", whoDiedMost, strings.Join(ways, ", "))

	f, err := os.Create("/tmp/output.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	pie.Render(chart.PNG, f)
}

func mostVaried(m map[string]map[string]int) string {
	var (
		victim string
		count  int
	)
	for v, typeCount := range m {
		distinctDeathTypes := len(typeCount)
		if distinctDeathTypes > count {
			victim = v
			count = distinctDeathTypes
		}
	}
	return victim
}

func mostDeaths(m map[string]int) string {
	var (
		victim string
		count  int
	)
	for v, c := range m {
		if c > count {
			victim = v
			count = c
		}
	}
	return victim
}

func processLogDir(logdir string, players Players) ([]Death, error) {
	deaths := make([]Death, 0)
	files, err := ioutil.ReadDir(logdir)
	if err != nil {
		return deaths, err
	}
	for _, fi := range files {
		if !strings.HasSuffix(fi.Name(), ".gz") && !strings.HasSuffix(fi.Name(), ".log") {
			continue
		}
		filename := filepath.Join(logdir, fi.Name())
		tr, err := GetTextReader(filename)
		if err != nil {
			return deaths, err
		}
		for {
			line, err := tr.ReadLine()
			if err != nil {
				if err == io.EOF {
					break
				}
				return deaths, err
			}
			ll, err := NewLogLine(filename, players)
			if err != nil {
				return deaths, err
			}
			err = ll.Parse(line)
			if err != nil {
				return deaths, err
			}
			death := ll.GetDeath()
			//fmt.Println(ll.Line)
			toks := strings.Fields(ll.Line)
			//fmt.Printf("isplayer %s: %t\n", toks[0], ll.IsPlayer(toks[0]))
			if death.Type == DeathNone && ll.PlayerInWhiteList(toks[0]) && !KnownEvent(line) {
				log.Printf("unrecognized death: %s", ll.Line)
			}
			if death.Type == DeathNone {
				continue
			}
			deaths = append(deaths, death)
		}
	}
	return deaths, nil
}

// TODO probably belongs in lib
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

// TODO probably belongs in lib
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

func validateFlags() error {
	if logdir == "" {
		return fmt.Errorf("logdir is required")
	}
	return nil
}
