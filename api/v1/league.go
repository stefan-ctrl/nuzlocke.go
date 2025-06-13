package v1

import (
	"strings"
)

type StarterType string

type League struct {
	Bosses map[string]Boss `json:",inline"`
}

type Boss struct {
	Name       string    `json:"name"`
	Speciality string    `json:"speciality"`
	Img        string    `json:"img"`
	Pokemon    []Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name      string   `json:"name"`
	Level     string   `json:"level"`
	Moves     []string `json:"moves" validate:"max=4"`
	Ability   string   `json:"ability"`
	Abilities []string `json:"abilities"`
	Held      string   `json:"held,omitempty"`
	Starter   string   `json:"starter,omitempty"` // Starter type of the player on the line to separate the Pokémons
	TeraType  string   `json:"tera_type,omitempty"`
}

// Helper to parse boss header fields into a Boss struct
func parseBossHeader(fields []string) Boss {
	boss := Boss{}
	if len(fields) > 1 {
		boss.Name = fields[1]
	}
	if len(fields) > 2 {
		boss.Speciality = fields[2]
	}
	if len(fields) > 3 {
		boss.Img = fields[3]
	}
	return boss
}

// Helper to parse a Pokémon line into a Pokemon struct
func parsePokemon(fields []string) Pokemon {
	poke := Pokemon{
		Name:    fields[0],
		Level:   fields[1],
		Moves:   strings.Split(fields[2], ","),
		Ability: fields[3],
	}
	if len(fields) > 4 {
		poke.Held = fields[4]
	}
	if len(fields) > 5 {
		poke.Starter = fields[5]
	}
	if len(fields) > 6 {
		poke.TeraType = fields[6]
	}
	poke.Abilities = []string{poke.Ability}
	return poke
}

func NewLeague(str string) *League {
	league := &League{
		Bosses: make(map[string]Boss),
	}
	lines := strings.Split(str, "\n")
	var currentBoss *Boss
	var currentKey string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "--") {
			fields := strings.SplitN(line[2:], "|", 5)
			if len(fields) < 1 {
				continue
			}
			currentKey = fields[0]
			boss := parseBossHeader(fields)
			currentBoss = &boss
			league.Bosses[currentKey] = *currentBoss
			continue
		}
		if currentBoss == nil {
			continue
		}
		fields := strings.Split(line, "|")
		if len(fields) < 4 {
			continue
		}
		poke := parsePokemon(fields)
		boss := league.Bosses[currentKey]
		boss.Pokemon = append(boss.Pokemon, poke)
		league.Bosses[currentKey] = boss
	}
	return league
}

func NewLeagueWithStarterSplit(str string) map[StarterType]*League {
	starters := []StarterType{"grass", "fire", "water"}
	leagues := make(map[StarterType]*League)
	for _, starter := range starters {
		leagues[starter] = &League{Bosses: make(map[string]Boss)}
	}
	lines := strings.Split(str, "\n")
	currentBossHeaders := make(map[StarterType]*Boss)
	var currentKey string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "--") {
			fields := strings.SplitN(line[2:], "|", 5)
			if len(fields) < 1 {
				continue
			}
			currentKey = fields[0]
			for _, starter := range starters {
				boss := parseBossHeader(fields)
				currentBossHeaders[starter] = &boss
			}
			continue
		}
		fields := strings.Split(line, "|")
		if len(fields) < 4 {
			continue
		}
		poke := parsePokemon(fields)
		for _, starter := range starters {
			if poke.Starter == "" || poke.Starter == string(starter) {
				boss := currentBossHeaders[starter]
				boss.Pokemon = append(boss.Pokemon, poke)
				currentBossHeaders[starter] = boss
			}
		}
		for _, starter := range starters {
			boss := currentBossHeaders[starter]
			if boss != nil && len(boss.Pokemon) > 0 {
				leagues[starter].Bosses[currentKey] = *boss
			}
		}
	}
	return leagues
}
