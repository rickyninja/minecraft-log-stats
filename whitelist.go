package minelog

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

var ErrPlayerNotFound = errors.New("player not found")

type Whitelist struct {
	Players Players
}

func (w Whitelist) GetName(name string) (Player, error) {
	for _, p := range w.Players {
		if p.Name == name {
			return p, nil
		}
	}
	return Player{}, ErrPlayerNotFound
}

func (w Whitelist) GetUUID(id string) (Player, error) {
	for _, p := range w.Players {
		if p.UUID == id {
			return p, nil
		}
	}
	return Player{}, ErrPlayerNotFound
}

func (w Whitelist) ContainsPlayer(player string) bool {
	for _, p := range w.Players {
		if p.Name == player {
			return true
		}
	}
	return false
}

func LoadWhitelistFile(filename string) (*Whitelist, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return LoadWhitelist(fd)
}

func LoadWhitelist(r io.Reader) (*Whitelist, error) {
	var p Players
	err := json.NewDecoder(r).Decode(&p)
	if err != nil {
		return nil, err
	}
	return &Whitelist{Players: p}, nil
}
