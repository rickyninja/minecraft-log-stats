package minelog

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"net/textproto"
	"testing"
	"time"
)

func TestWhitelist(t *testing.T) {
	data := []byte(`[
  {
    "uuid": "74c82b84-9ea7-4330-b5c0-0d705656eb6e",
    "name": "Sanrei_"
  },
  {
    "uuid": "f6e1e899-4389-4539-a3cf-1cb00cd89e67",
    "name": "Sanji15"
  },
  {
    "uuid": "2ca2332f-429c-46b9-926f-03f9d7dc8049",
    "name": "ricky_ninja"
  }
]`)
	var got Players
	err := json.Unmarshal(data, &got)
	if err != nil {
		t.Fatal(err)
	}
	want := Players{
		Player{
			Name: "Sanrei_",
			UUID: "74c82b84-9ea7-4330-b5c0-0d705656eb6e",
		},
		Player{
			Name: "Sanji15",
			UUID: "f6e1e899-4389-4539-a3cf-1cb00cd89e67",
		},
		Player{
			Name: "ricky_ninja",
			UUID: "2ca2332f-429c-46b9-926f-03f9d7dc8049",
		},
	}
	if len(got) != len(want) {
		t.Fatalf("wrong len for Players got %d want %d", len(got), len(want))
	}
	for i, w := range want {
		g := got[i]
		if g.Name != w.Name {
			t.Errorf("wrong Name got %s want %s", g.Name, w.Name)
		}
		if g.UUID != w.UUID {
			t.Errorf("wrong UUID got %s want %s", g.UUID, w.UUID)
		}
	}
}

func TestGetTimeFromFilename(t *testing.T) {
	now := time.Now().In(time.UTC)
	cases := []struct {
		filename string
		want     time.Time
	}{
		{"2019-09-18-1.log.gz", time.Date(2019, 9, 18, 0, 0, 0, 0, time.UTC)},
		// TODO I'm not certain time.Now() is an appropriate assumption for latest.log.
		// Will days rollover in one logfile?  See about putting better date string in logs.
		{"latest.log", time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)},
	}
	for _, tc := range cases {
		got, err := getTimeFromFilename(tc.filename)
		if err != nil {
			t.Fatal(err)
		}
		if got.Year() != tc.want.Year() {
			t.Errorf("wrong year got %d want %d", got.Year(), tc.want.Year())
		}
		if got.Month() != tc.want.Month() {
			t.Errorf("wrong month got %d want %d", got.Month(), tc.want.Month())
		}
		if got.Day() != tc.want.Day() {
			t.Errorf("wrong day got %d want %d", got.Day(), tc.want.Day())
		}
		if got.Location() != tc.want.Location() {
			t.Errorf("wrong location got %s want %s", got.Location(), tc.want.Location())
		}
	}
}

func TestLogLine_Parse(t *testing.T) {
	data := []byte(`[01:01:52] [User Authenticator #21/INFO]: UUID of player Sanrei_ is 74c82b84-9ea7-4330-b5c0-0d705656eb6e
[01:01:52] [Server thread/INFO]: Sanrei_[/98.171.94.68:51425] logged in with entity id 276786 at (-175.17652997589863, 56.5, -335.5669264472643)
[01:01:52] [Server thread/INFO]: Sanrei_ joined the game
[01:03:01] [Server thread/WARN]: Fetching packet for removed entity atx['Air'/276938, l='rickyninja', x=-123.80, y=59.00, z=-355.80]
[01:04:28] [Server thread/WARN]: Fetching packet for removed entity atx['Air'/277121, l='rickyninja', x=-123.80, y=59.00, z=-355.80]
[01:17:58] [User Authenticator #22/INFO]: UUID of player ricky_ninja is 2ca2332f-429c-46b9-926f-03f9d7dc8049
[01:17:58] [Server thread/INFO]: ricky_ninja[/98.161.238.54:59754] logged in with entity id 284180 at (-49.82148007647506, 45.0, -432.4264969504812)
[01:17:58] [Server thread/INFO]: ricky_ninja joined the game
[01:24:03] [Server thread/WARN]: Fetching packet for removed entity atx['Air'/286440, l='rickyninja', x=-123.80, y=59.00, z=-355.20]
[01:24:23] [Server thread/WARN]: Fetching packet for removed entity atx['Air'/286508, l='rickyninja', x=-123.80, y=59.00, z=-355.80]
[01:32:55] [Server thread/WARN]: Fetching packet for removed entity atx['Air'/291045, l='rickyninja', x=-123.80, y=59.00, z=-355.20]
[01:48:33] [Server thread/WARN]: Fetching packet for removed entity atx['Air'/298013, l='rickyninja', x=-123.80, y=59.00, z=-355.20]
[01:48:33] [Server thread/WARN]: Fetching packet for removed entity atx['Air'/298013, l='rickyninja', x=-123.80, y=59.00, z=-355.20]
[02:08:48] [Server thread/WARN]: Fetching packet for removed entity atx['Air'/305999, l='rickyninja', x=-123.80, y=59.00, z=-355.80]
[02:13:42] [Server thread/INFO]: Gamerule keepInventory is now set to: true
[02:15:03] [Server thread/INFO]: ricky_ninja fell from a high place
[02:15:06] [Server thread/WARN]: ricky_ninja moved too quickly! 8.533917584639994,-8.499999799999998,13.815646564463066
[02:27:56] [Server thread/WARN]: Fetching packet for removed entity atx['Air'/313266, l='rickyninja', x=-123.20, y=59.00, z=-355.80]
[02:37:47] [Server thread/WARN]: Fetching packet for removed entity atx['Air'/318141, l='rickyninja', x=-123.80, y=59.00, z=-355.20]
[02:37:47] [Server thread/WARN]: Fetching packet for removed entity atx['Air'/318141, l='rickyninja', x=-123.80, y=59.00, z=-355.20]
[02:56:34] [Server thread/WARN]: Can't keep up! Is the server overloaded? Running 14705ms or 294 ticks behind
[03:20:02] [Server thread/WARN]: Can't keep up! Is the server overloaded? Running 2891ms or 57 ticks behind
[03:25:05] [Server thread/WARN]: ricky_ninja moved too quickly! -1.7976588995634302,0.0,10.005226292941984
[03:25:05] [Server thread/WARN]: Can't keep up! Is the server overloaded? Running 3260ms or 65 ticks behind
[04:21:26] [Server thread/INFO]: Sanrei_ lost connection: Disconnected
[04:21:26] [Server thread/INFO]: Sanrei_ left the game
[04:39:00] [Server thread/INFO]: ricky_ninja lost connection: Disconnected
[04:39:00] [Server thread/INFO]: ricky_ninja left the game
[05:33:32] [User Authenticator #23/INFO]: UUID of player ricky_ninja is 2ca2332f-429c-46b9-926f-03f9d7dc8049
[05:33:32] [Server thread/INFO]: ricky_ninja[/98.161.238.54:60032] logged in with entity id 365869 at (2789.762499988079, 63.0, -345.69999998807907)
[05:33:32] [Server thread/INFO]: ricky_ninja joined the game
[05:40:34] [User Authenticator #24/INFO]: UUID of player Sanrei_ is 74c82b84-9ea7-4330-b5c0-0d705656eb6e
[05:40:34] [Server thread/INFO]: Sanrei_[/98.171.94.68:52440] logged in with entity id 369256 at (2788.1443483583316, 63.0, -342.02148361508006)
[05:40:34] [Server thread/INFO]: Sanrei_ joined the game
[05:54:08] [Server thread/INFO]: Sanrei_ lost connection: Timed out
[05:54:08] [Server thread/INFO]: Sanrei_ left the game
[06:12:21] [User Authenticator #25/INFO]: UUID of player Sanrei_ is 74c82b84-9ea7-4330-b5c0-0d705656eb6e
[06:12:21] [Server thread/INFO]: Sanrei_[/98.171.94.68:52679] logged in with entity id 379659 at (2892.8133002489567, 70.0, -529.2102536253246)
[06:12:21] [Server thread/INFO]: Sanrei_ joined the game
[06:39:56] [Server thread/INFO]: Sanrei_ lost connection: Disconnected
[06:39:56] [Server thread/INFO]: Sanrei_ left the game
[06:46:40] [User Authenticator #26/INFO]: UUID of player Sanrei_ is 74c82b84-9ea7-4330-b5c0-0d705656eb6e
[06:46:40] [Server thread/INFO]: Sanrei_[/98.171.94.68:52745] logged in with entity id 392481 at (2801.218979518151, 14.0, -500.72136006553615)
[06:46:40] [Server thread/INFO]: Sanrei_ joined the game
[07:27:30] [Server thread/INFO]: Sanrei_ lost connection: Disconnected
[07:27:30] [Server thread/INFO]: Sanrei_ left the game
[07:30:34] [User Authenticator #27/INFO]: UUID of player Sanrei_ is 74c82b84-9ea7-4330-b5c0-0d705656eb6e
[07:30:34] [Server thread/INFO]: Sanrei_[/98.171.94.68:52949] logged in with entity id 406496 at (2763.025204821351, 41.0, -570.8728403995634)
[07:30:34] [Server thread/INFO]: Sanrei_ joined the game
[07:30:52] [Server thread/INFO]: Sanrei_ lost connection: Timed out
[07:30:52] [Server thread/INFO]: Sanrei_ left the game
[07:36:47] [User Authenticator #28/INFO]: UUID of player Sanrei_ is 74c82b84-9ea7-4330-b5c0-0d705656eb6e
[07:36:47] [Server thread/INFO]: Sanrei_[/98.171.94.68:53023] logged in with entity id 407228 at (2763.633089508713, 41.0, -575.6262156324581)
[07:36:47] [Server thread/INFO]: Sanrei_ joined the game
[07:37:12] [Server thread/INFO]: Sanrei_ lost connection: Timed out
[07:37:12] [Server thread/INFO]: Sanrei_ left the game
[07:45:31] [User Authenticator #29/INFO]: UUID of player Sanrei_ is 74c82b84-9ea7-4330-b5c0-0d705656eb6e
[07:45:31] [Server thread/INFO]: Sanrei_[/98.171.94.68:53134] logged in with entity id 410241 at (2763.633089508713, 41.0, -575.6262156324581)
[07:45:31] [Server thread/INFO]: Sanrei_ joined the game
[08:41:36] [Server thread/INFO]: ricky_ninja lost connection: Disconnected
[08:41:36] [Server thread/INFO]: ricky_ninja left the game
[09:35:45] [Server thread/WARN]: Fetching packet for removed entity atx['Air'/434667, l='rickyninja', x=-123.20, y=59.00, z=-355.80]
[09:36:25] [Server thread/INFO]: Sanrei_ lost connection: Disconnected
[09:36:25] [Server thread/INFO]: Sanrei_ left the game
[23:20:01] [User Authenticator #30/INFO]: UUID of player ricky_ninja is 2ca2332f-429c-46b9-926f-03f9d7dc8049
[23:20:01] [Server thread/INFO]: ricky_ninja[/98.161.238.54:32974] logged in with entity id 434745 at (2788.488739127381, 63.0, -342.2679383036058)
[23:20:01] [Server thread/INFO]: ricky_ninja joined the game
[23:30:58] [Server thread/INFO]: ricky_ninja lost connection: Disconnected
[23:30:58] [Server thread/INFO]: ricky_ninja left the game`)
	whitelist := &Whitelist{
		Players: Players{
			Player{Name: "ricky_ninja"},
			Player{Name: "Sanrei_"},
		},
	}
	filename := "testdata/2019-09-18-1.log.gz"
	r := textproto.NewReader(bufio.NewReader(bytes.NewReader(data)))
	for {
		line, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatal(err)
		}
		ll, err := NewLogLine(filename, whitelist)
		if err != nil {
			t.Fatal(err)
		}
		err = ll.Parse(line)
		if err != nil {
			t.Fatal(err)
		}
		const (
			wantYear  = 2019
			wantMonth = 9
			wantDay   = 18
		)
		if ll.Time.Year() != wantYear {
			t.Errorf("expected year %d got %d", wantYear, ll.Time.Year())
		}
		if ll.Time.Month() != wantMonth {
			t.Errorf("expected month %d got %d", wantMonth, ll.Time.Month())
		}
		if ll.Time.Day() != wantDay {
			t.Errorf("expected day %d got %d", wantDay, ll.Time.Day())
		}
		if ll.Level > 1 {
			t.Errorf("unexpected level %d", ll.Level)
		}
		if ll.Line == "" {
			t.Errorf("found empty log line")
		}
	}
}

func TestLogLine_GetDeath(t *testing.T) {
	whitelist := &Whitelist{
		Players: Players{
			Player{Name: "ricky_ninja"},
			Player{Name: "Sanrei_"},
			Player{Name: "ChadMan9010"},
		},
	}
	cases := []struct {
		name    string
		logLine LogLine
		want    Death
	}{
		{
			"chat message",
			LogLine{
				Time:      time.Now(),
				Line:      "<ricky_ninja> blew up your building?",
				whitelist: whitelist,
			},
			Death{
				Type:   DeathNone,
				Victim: "",
				Killer: "",
			},
		},
		{
			"death by explosion",
			LogLine{
				Time:      time.Now(),
				Line:      "ricky_ninja blew up",
				whitelist: whitelist,
			},
			Death{
				Type:   DeathExplosion,
				Victim: "ricky_ninja",
				Killer: "",
			},
		},
		{
			"death by falling",
			LogLine{
				Time:      time.Now(),
				Line:      "ricky_ninja fell from a high place",
				whitelist: whitelist,
			},
			Death{
				Type:   DeathFall,
				Victim: "ricky_ninja",
				Killer: "",
			},
		},
		{
			"death by squishing",
			LogLine{
				Time:      time.Now(),
				Line:      "Sanrei_ was squished too much",
				whitelist: whitelist,
			},
			Death{
				Type:   DeathSquish,
				Victim: "Sanrei_",
				Killer: "",
			},
		},
		{
			"death by drowning",
			LogLine{
				Time:      time.Now(),
				Line:      "Sanrei_ drowned",
				whitelist: whitelist,
			},
			Death{
				Type:   DeathDrown,
				Victim: "Sanrei_",
				Killer: "",
			},
		},
		{
			"death by creeper",
			LogLine{
				Time:      time.Now(),
				Line:      "ricky_ninja was blown up by Creeper",
				whitelist: whitelist,
			},
			Death{
				Type:   DeathExplosion,
				Victim: "ricky_ninja",
				Killer: "Creeper",
			},
		},
		{
			"burnt by ghast",
			LogLine{
				Time:      time.Now(),
				Line:      "ChadMan9010 was burnt to a crisp whilst fighting Ghast",
				whitelist: whitelist,
			},
			Death{
				Type:   DeathFire,
				Victim: "ChadMan9010",
				Killer: "Ghast",
			},
		},
	}
	for _, tc := range cases {
		got := tc.logLine.GetDeath()
		want := tc.want
		if got.Type != want.Type {
			t.Errorf("case name: [%s] wrong death type got %s want %s", tc.name, got.Type, want.Type)
		}
		if got.Victim != want.Victim {
			t.Errorf("case name: [%s] wrong victim got %s want %s", tc.name, got.Victim, want.Victim)
		}
		if got.Killer != want.Killer {
			t.Errorf("case name: [%s] wrong killer got %s want %s", tc.name, got.Killer, want.Killer)
		}
	}
}
