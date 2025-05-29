package v1

import (
	"strings"
)

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
	Starter   string   `json:"starter,omitempty"`
	TeraType  string   `json:"tera_type,omitempty"`
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
			// Boss header
			fields := strings.SplitN(line[2:], "|", 5)
			if len(fields) < 1 {
				continue
			}
			currentKey = fields[0]
			currentBoss = &Boss{}
			if len(fields) > 1 {
				currentBoss.Name = fields[1]
			}
			if len(fields) > 2 {
				currentBoss.Speciality = fields[2]
			}
			if len(fields) > 3 {
				currentBoss.Img = fields[3]
			}
			league.Bosses[currentKey] = *currentBoss
			continue
		}
		// Pokemon line
		if currentBoss == nil {
			continue
		}
		fields := strings.Split(line, "|")
		if len(fields) < 4 {
			continue
		}
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
		// Add ability to Abilities slice for compatibility
		poke.Abilities = []string{poke.Ability}
		boss := league.Bosses[currentKey]
		boss.Pokemon = append(boss.Pokemon, poke)
		league.Bosses[currentKey] = boss
	}
	return league
}
