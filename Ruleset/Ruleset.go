package Ruleset

type Ruleset interface {
	IsAlive(field *[][]bool, x int, y int) bool
}
