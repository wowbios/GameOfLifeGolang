package Components

import (
	"container/list"
)

type Game struct {
	Settings   GameSettings
	Field      *[][]bool
	Generation int64
}

func NewGame(settings GameSettings) *Game{
	var field = make([][]bool, settings.Width)
	for i := range field {
		field[i] = make([]bool, settings.Height)
	}
	return &Game {settings, &field, 0}
}

func (game *Game) readEvents() []ChangeEvent{
	var events []ChangeEvent = nil
	for i, row := range *game.Field {
		for j, col := range row {
			if events == nil {
				events = make([]ChangeEvent, len(*game.Field) * len(row))
			}

			events[i*j + j] = ChangeEvent{i,j, col}
		}
	}
	return events
}

func (game *Game) Render(events []ChangeEvent) {
	game.Settings.Render.Render(&GameState{game.Generation, events})
}

func (game *Game) Prepare () {
	game.Settings.Preset.InitializeField(game.Field)

	game.Render(game.readEvents())
}

func (game *Game) MakeNextGeneration(){
	newField := *game.Field
	events := list.New()

	for i, xi := range *game.Field {
		for j, yj := range xi {
			func(x, y int) {
				isAlive := game.Settings.Ruleset.IsAlive(game.Field, x, y)
				if yj != isAlive {
					events.PushBack(ChangeEvent{x, y, isAlive})
					newField[x][y] = isAlive
				}
			} (i, j)
		}
	}

	game.Generation++
	eventsArr := make([]ChangeEvent, events.Len())
	var i = 0
	for  e := events.Front(); e != nil; e = e.Next() {
		eventsArr[i] = e.Value.(ChangeEvent)
		i++
	}
	if events.Len() > 0 {
		game.Render(eventsArr)
	}
}

