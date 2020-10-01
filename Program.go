package main

import (
	"GameOfLifeGolang/Components"
	"GameOfLifeGolang/Preset"
	"GameOfLifeGolang/Render"
	"GameOfLifeGolang/Ruleset"
	"fmt"
)

func main() {
	settings := Components.GameSettings{
		Width:   10,
		Height:  10,
		Preset:  new (Preset.GliderAtTheMiddlePreset),
		Render:  new (Render.ConsoleRender),
		Ruleset: new (Ruleset.ConwaysRuleset),
	}
	field := make([][]bool, settings.Width)
	for i, _ := range field {
		field[i] = make([]bool, settings.Height)
	}

	var game = Components.Game{
		Settings: settings,
		Field: &field,
	}

	game.Prepare()

	go cycle(&game)

	var input string
	fmt.Scanln(&input)
}

func cycle(game *Components.Game){
	game.MakeNextGeneration()
	// time.Sleep(100 * time.Millisecond)
	go cycle(game)
}