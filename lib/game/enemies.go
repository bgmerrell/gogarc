package game

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

const enemiesPath = "content/enemies/enemies.json"

func LoadEnemies() ([]Being, error) {
	enemies := []Being{}
	f, err := os.Open(enemiesPath)
	if err != nil {
		return enemies, errors.New("Failed to load enemies file: " + err.Error())
	}
	dec := json.NewDecoder(f)

	var b Being
	for {
		if err = dec.Decode(&b); err == io.EOF {
			err = nil
			break
		} else if err != nil {
			return enemies, errors.New(
				"Failed to decode enemies file: " + err.Error())
		}
		enemies = append(enemies, b)
	}
	return enemies, err
}
