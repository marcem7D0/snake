package snake

import "fmt"

type Block byte

var (
	Edge       Block = '#'
	PlayerHead Block = '@'
	PlayerBody Block = '*'
	Apple      Block = '%'
	Empty      Block = ' '
)

type Map struct {
	Width  int
	Height int

	cells [][]Block
}

func NewMap(width, height int) Map {
	height = height - 1
	cells := make([][]Block, height)

	for i := range height {
		cells[i] = make([]Block, width)

		for j := range width {
			if i == 0 || i == height-1 || j == 0 || j == width-1 {
				cells[i][j] = Edge
				continue
			}

			cells[i][j] = Empty
		}
	}

	return Map{Width: width, Height: height, cells: cells}
}

func (m Map) Draw() {
	fmt.Print("\033[H")

	for i := range m.Height {
		for j := range m.Width {
			fmt.Printf("%s", string(m.cells[i][j]))
		}
		fmt.Println()
	}
}
