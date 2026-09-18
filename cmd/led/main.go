package main

import (
	"fmt"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
)

var g_height int = 32
var g_width int = 32

type tickMsg time.Time

type Cell struct {
	isAlive bool
}

type model struct {
	matrix [][]Cell
	gameState int // 1 = running, 0 = stopped. potentially other states to be added
}

func (m model) matrixAsString() string {
	retVal := ""
	for _, row := range m.matrix {
		for _, cell := range row {
			if cell.isAlive{ 
				retVal += "0 "
			} else { 
				retVal += ". "
			}
		}
		retVal += "\r\n"
	}
	return retVal
}

func initialModel() model {
	startingCells := [][]int{{1,1},{2,2},{2,3},{3,1},{3,2}}

	matrix := make([][]Cell, g_height)

	for i := 0; i < g_height; i++ {
		matrix[i] = make([]Cell, g_width)
		for j := 0; j < g_width; j++ {
			matrix[i][j] = Cell{ false }
		}
	}

	for _, pos := range startingCells {
		matrix[pos[0]][pos[1]].isAlive = true
	}

	return model{ matrix, 1 }
}

func (m model) Init() tea.Cmd {
    return tickCmd()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyPressMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            return m, tea.Quit
		}
	case tickMsg:
		m.ProgressCells()
		return m, tickCmd()
    }
	return m, nil
}

func (m model) outOfBounds(y, x int) bool {
	if y < 0 { return true }
	if x < 0 { return true }
	if y >= len(m.matrix) { return true }
	if x >= len(m.matrix[y]) { return true }
	return false
}

func (m model) CountLivingNeighbours(y, x int) int {
	n := 0
	directions := [][]int{{-1,-1},{-1,0},{-1,1},{0,-1},{0,1},{1,-1},{1,0},{1,1}}

	for _, pos := range directions {
		n_y := y + pos[0]
		n_x := x + pos[1]
		if m.outOfBounds(n_y, n_x) {
			continue
		}
		if m.matrix[n_y][n_x].isAlive {
			n++
		}
	}

	return n
}

func (m *model) ProgressCells() {
	newMatrix := make([][]Cell, g_height)
	m.gameState = 0

	for i := 0; i < len(m.matrix); i++ {
		newMatrix[i] = make([]Cell, g_width)

		for j := 0; j < len(m.matrix[i]); j++ {
			neighbours := m.CountLivingNeighbours(i, j)

			if (neighbours == 2 && m.matrix[i][j].isAlive) || neighbours == 3 {
				newMatrix[i][j] = Cell{ true }
				m.gameState = 1

			} else {
				newMatrix[i][j] = Cell{ false }
			}
		}
	}
	m.matrix = newMatrix
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond * 100, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) View() tea.View {
    s := "Conways Game of Life Simulation\n\n"

	s += m.matrixAsString()

    s += "\nPress q to quit.\n"

    return tea.NewView(s)
}

func main() {
	p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}
