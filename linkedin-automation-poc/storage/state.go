package storage

import (
	"encoding/json"
	"os"
)

type State struct {
	Actions int `json:"actions"`
}

func Save(state State) {
	data, _ := json.MarshalIndent(state, "", "  ")
	_ = os.WriteFile("state.json", data, 0644)
}
