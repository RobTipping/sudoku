package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"

	solver "github.com/RobTipping/sudoku/internal"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var normStyle1 = lipgloss.NewStyle().
	//Bold(true).
	Foreground(lipgloss.Color("#000000")).
	Background(lipgloss.Color("#C0CFB2")).
	PaddingLeft(1).
	PaddingRight(1)

var normStyle2 = lipgloss.NewStyle().
	//Bold(true).
	Foreground(lipgloss.Color("#000000")).
	Background(lipgloss.Color("#8BBF88")).
	PaddingLeft(1).
	PaddingRight(1)

var styles = [6]lipgloss.Style{normStyle1, normStyle2, invalidStyle1, invalidStyle2, styleSelected, invalidSelected}

var styleSelected = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#000000")).
	Background(lipgloss.Color("#8BA888")).
	PaddingLeft(1).
	PaddingRight(1)

var invalidStyle1 = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#F70D1A")).
	Background(lipgloss.Color("#FFCCCB")).
	PaddingLeft(1).
	PaddingRight(1)

var invalidStyle2 = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#F70D1A")).
	Background(lipgloss.Color("#F1807E")).
	PaddingLeft(1).
	PaddingRight(1)

var invalidSelected = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#F70D1A")).
	Background(lipgloss.Color("#8BA888")).
	PaddingLeft(1).
	PaddingRight(1)

var styleBorder = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFFFFF")).
	Background(lipgloss.Color("#44624A"))

type model struct {
	grid       [9][9]int
	solvedGrid [9][9]int
	cursor     [2]int
	tickCount  int
	solved     bool
}

func initialModel() model {
	return model{
		grid: [9][9]int{
			// Uncomment and use as templete grid values
			{0, 0, 0, 0, 0, 0, 0, 0, 1},
			{0, 0, 2, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 3, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		solved: false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.tickCount = 0
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "up", "k":
			if m.cursor[0] > 0 {
				m.cursor[0]--
			}
		case "down", "j":
			if m.cursor[0] < len(m.grid)-1 {
				m.cursor[0]++
			}
		case "left", "h":
			if m.cursor[1] > 0 {
				m.cursor[1]--
			}
		case "right", "l":
			if m.cursor[1] < len(m.grid[m.cursor[0]])-1 {
				m.cursor[1]++
			}
		case "s":
			m.solvedGrid, m.solved = solver.SolveGrid(m.grid)
			if m.solved == true {
				m.grid = m.solvedGrid
			}
		case "enter", " ":
			if m.cursor[1] < len(m.grid[m.cursor[0]])-1 {
				m.cursor[1]++
			} else {
				if m.cursor[0] < len(m.grid)-1 {
					m.cursor[0]++
					m.cursor[1] = 0
				}
			}
		case "c":
			for i := range m.grid {
				for j := range m.grid[i] {
					m.grid[i][j] = 0
				}
			}
		}

		if len(msg.Runes) != 0 {
			if unicode.IsDigit(msg.Runes[0]) == true {
				if value, err := strconv.Atoi(msg.String()); err == nil {
					m.grid[m.cursor[0]][m.cursor[1]] = value
					if m.cursor[1] < len(m.grid[m.cursor[0]])-1 {
						m.cursor[1]++
					} else {
						if m.cursor[0] < len(m.grid)-1 {
							m.cursor[0]++
							m.cursor[1] = 0
						}
					}
				}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Enter Grid To Solve\nPress \"S\" to find the solution\n"
	for i := range m.grid {
		if i == 3 || i == 6 {
			s += styleBorder.Render("                             ")
			s += "\n"
		}
		for j := range m.grid[i] {
			if j == 3 || j == 6 {
				s += styleBorder.Render(" ")
			}
			squareStyle := 0
			if m.cursor[0] == i && m.cursor[1] == j {
				if m.grid[i][j] != 0 && solver.CheckSquare(m.grid, i, j) == false {
					squareStyle = 5
				} else {
					squareStyle = 4
				}
			} else {
				squareStyle = ((i * len(m.grid[i])) + j) % 2
				if m.grid[i][j] != 0 && solver.CheckSquare(m.grid, i, j) == false {
					squareStyle += 2
				}
			}
			// 2 strings here one for number one for 0
			if m.grid[i][j] != 0 {
				s += styles[squareStyle].Render(fmt.Sprintf("%d", m.grid[i][j]))
			} else {
				s += styles[squareStyle].Render(" ")
			}

		}
		s += "\n"
	}
	s += "\n"
	s += fmt.Sprintf("\n%v\n", m.solved)
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error, %v", err)
		os.Exit(1)
	}
}
