package v1_test

import (
	"github.com/go-json-experiment/json"
	"github.com/kinbiko/jsonassert"
	"github.com/stefan-ctrl/nuzlocke.go/api/v1"
	"github.com/stefan-ctrl/nuzlocke.go/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

const gotFile = ".txt"
const nuzlockeGotDataPath = "../../nuzlocke.data/leagues/"

const wantFile = ".json"
const nuzlockeWantDataPath = "./test/leagues/"

func TestNewLeague(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "bw",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			bytes, err := util.ReadFile(nuzlockeGotDataPath + tt.name + gotFile)
			if err != nil {
				t.Fatalf("Failed to read file %s: %v", tt.name, err)
			}
			got := v1.NewLeague(string(*bytes))

			bytes, err = util.ReadFile(nuzlockeWantDataPath + tt.name + wantFile)
			if err != nil {
				t.Fatalf("Failed to read file %s: %v", tt.name, err)
			}

			// Unmarshal the expected data
			want := &v1.League{}
			err = json.Unmarshal(*bytes, want)
			if err != nil {
				t.Fatalf("Failed to unmarshal expected data: %v", err)
			}

			assert.Equal(t, want, got)
		})
	}

}

func TestNewLeagueWithStarterSplit(t *testing.T) {
	tests := []struct {
		name     string
		starters []string
	}{
		{
			name:     "renplat",
			starters: []string{"grass", "fire", "water"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			bytes, err := util.ReadFile(nuzlockeGotDataPath + tt.name + gotFile)
			if err != nil {
				t.Fatalf("Failed to read file %s: %v", tt.name, err)
			}
			got := v1.NewLeagueWithStarterSplit(string(*bytes))

			for _, starter := range tt.starters {
				if _, ok := got[v1.StarterType(starter)]; !ok {
					assert.Failf(t, "Failed to find starter \"%s\"", starter)
				}
				pokemonTypeInFile := "." + starter
				bytes, err = util.ReadFile(nuzlockeWantDataPath + tt.name + pokemonTypeInFile + wantFile)
				if err != nil {
					t.Fatalf("Failed to read file %s: %v", tt.name, err)
				}

				gotAsJson, err := json.Marshal(got[v1.StarterType(starter)])
				if err != nil {
					t.Fatalf("Failed to unmarshal expected data: %v", err)
				}

				// Unmarshal the expected data
				ja := jsonassert.New(t)
				ja.Assert(string(gotAsJson), string(*bytes))
			}

		})
	}

}
