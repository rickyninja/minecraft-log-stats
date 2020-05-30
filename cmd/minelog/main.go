package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/rickyninja/minelog"
	"github.com/wcharczuk/go-chart"
)

type arg struct {
	logdir    string
	whitelist string
}

func (a *arg) Validate() error {
	if a.logdir == "" {
		return fmt.Errorf("logdir is required")
	}
	if a.whitelist == "" {
		return fmt.Errorf("whitelist is required")
	}
	return nil
}

func (a *arg) Parse() {
	flag.StringVar(&a.logdir, "logdir", "", "the directory with log files")
	flag.StringVar(&a.whitelist, "whitelist", "", "the whitelist.json file")
	flag.Parse()
}

func main() {
	cli := new(arg)
	cli.Parse()
	if err := cli.Validate(); err != nil {
		log.Fatal(err)
	}
	whitelist, err := minelog.LoadWhitelistFile(cli.whitelist)
	deaths, err := processLogDir(cli.logdir, whitelist)
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
		if death.Type == minelog.DeathMob {
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

func processLogDir(logdir string, whitelist *minelog.Whitelist) ([]minelog.Death, error) {
	deaths := make([]minelog.Death, 0)
	files, err := ioutil.ReadDir(logdir)
	if err != nil {
		return deaths, err
	}
	for _, fi := range files {
		if !strings.HasSuffix(fi.Name(), ".gz") && !strings.HasSuffix(fi.Name(), ".log") {
			continue
		}
		filename := filepath.Join(logdir, fi.Name())
		tr, err := minelog.GetTextReader(filename)
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
			ll, err := minelog.NewLogLine(filename, whitelist)
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
			if death.Type == minelog.DeathNone && whitelist.ContainsPlayer(toks[0]) && !minelog.KnownEvent(line) {
				log.Printf("unrecognized death: %s", ll.Line)
			}
			if death.Type == minelog.DeathNone {
				continue
			}
			deaths = append(deaths, death)
		}
	}
	return deaths, nil
}
