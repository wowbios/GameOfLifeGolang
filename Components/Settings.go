package Components

import (
	"GameOfLifeGolang/Preset"
	"GameOfLifeGolang/Ruleset"
)

type GameSettings struct {
	Width, Height int
	Preset        Preset.Preset
	Render        Render
	Ruleset       Ruleset.Ruleset
}
