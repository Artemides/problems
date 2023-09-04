package dp

import (
	"container/list"
	"fmt"
	"reflect"
)

type Cell struct {
	cell  int
	right int
	left  int
	top   int
	down  int
}

func traverseMaze(paths []Cell, visited *list.List, path *list.List, start Cell, end Cell) bool {
	fmt.Println("start ", start)
	if start == end {
		path.PushFront(start.cell)
		return true
	}

	visited.PushBack(start.cell)
	options := reflect.ValueOf(start)
	for i := 0; i < options.NumField(); i++ {
		option := int(options.Field(i).Int())

		if option == start.cell || option == 0 {
			continue
		}
		tried := findVisited(*visited, option)
		fmt.Println("tried ", option, tried)

		if tried {
			continue
		}

		cell := findCell(paths, option)
		fmt.Println(cell)
		if traverseMaze(paths, visited, path, cell, end) {
			path.PushFront(start.cell)
			return true
		}

	}
	return false

}

func findVisited(visited list.List, val int) bool {

	for element := visited.Front(); element != nil; element = element.Next() {
		value := element.Value
		if value == val {
			return true
		}
	}
	return false
}

func findCell(paths []Cell, cell int) Cell {
	for _, v := range paths {
		if v.cell == cell {
			return v
		}
	}
	return Cell{}
}

func Maze() {
	one := Cell{cell: 1}
	one.right = 2
	two := Cell{cell: 2}
	two.left = 1
	two.down = 6
	three := Cell{cell: 3}
	three.down = 7
	four := Cell{cell: 4}
	four.down = 8
	five := Cell{cell: 5}
	five.right = 6
	six := Cell{cell: 6}
	six.left = 5
	six.top = 2
	six.down = 10
	seven := Cell{cell: 7}
	seven.right = 8
	seven.top = 3
	seven.down = 11
	eight := Cell{cell: 8}
	eight.left = 7
	eight.top = 4
	eight.down = 12
	nine := Cell{cell: 9}
	nine.right = 10
	nine.down = 13
	ten := Cell{cell: 10}
	ten.right = 11
	ten.left = 9
	ten.top = 6
	ten.down = 14
	eleven := Cell{cell: 11}
	eleven.left = 10
	eleven.top = 7
	eleven.down = 15
	twelve := Cell{cell: 12}
	twelve.top = 8
	twelve.down = 16
	thirteen := Cell{cell: 13}
	thirteen.top = 9
	fourteen := Cell{cell: 14}
	fourteen.top = 10
	fifteen := Cell{cell: 15}
	fifteen.top = 11
	sixteen := Cell{cell: 16}
	sixteen.top = 12
	maze := []Cell{
		one,
		two,
		three,
		four,
		five,
		six,
		seven,
		eight,
		nine, ten, eleven, twelve, thirteen, fourteen, fifteen, sixteen}

	visited := list.New()
	path := list.New()

	scaped := traverseMaze(maze, visited, path, five, eight)

	fmt.Println("scaped ", scaped)
	for element := path.Front(); element != nil; element = element.Next() {
		value := element.Value
		fmt.Print("=>>", value)

	}
	fmt.Println("")
	for element := visited.Front(); element != nil; element = element.Next() {
		value := element.Value
		fmt.Print("=>>", value)

	}
}
