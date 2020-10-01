package Render

import (
	"GameOfLifeGolang/Components"
	"fmt"
)

type ConsoleRender struct {

}

func (render *ConsoleRender) Render(state *Components.GameState){
	// setCursorPosition(0, 0)
	fmt.Print(state.Generation)
	for _, event := range state.Events {
		// setCursorPosition(event.X, event.Y)
		fmt.Printf("Event %d %d %v\n", event.X, event.Y, event.IsAlive)
	/*	if event.IsAlive{
			fmt.Print("█")
		} else {
			fmt.Print(" ")
		}*/
	}
}

func setCursorPosition(line, col int) {
	fmt.Print(MoveTo(line, col))
}

func MoveTo(line, col int) string {
	return escape("[%d;%dH", line, col)
}

var Esc = "\x1b"

func escape(format string, args ...interface{}) string {
	return fmt.Sprintf("%s%s", Esc, fmt.Sprintf(format, args...))
}