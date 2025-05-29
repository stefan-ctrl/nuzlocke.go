package v1_test

import (
	"github.com/go-json-experiment/json"
	"github.com/stefan-ctrl/nuzlocke.go/api/v1"
	"github.com/stefan-ctrl/nuzlocke.go/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

const gotFile = ".txt"
const nuzlockeGotDataPath = "../../nuzlocke.data/leagues/"

const wantFile = ".json"
const nuzlockeWantDataPath = "./test/"

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
