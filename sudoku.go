package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    grid    [][]int
    cursor  [2]int
}

func initialModel() model {
    return model{
        grid:[][]int{{1, 2, 3},{4, 5, 6},{7, 8, 9},{0, 0, 0},},
    }
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
        case tea.KeyMsg:
            switch msg.String() {
                case "ctrl+c", "q":
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
            }
    }
    return m, nil
}

func (m model) View() string {
    s := fmt.Sprintf("%d%d\n", m.cursor[0],m.cursor[1])
    for i, _ := range m.grid {
        cusor := [3]string{" ", " ", " "}
        if m.cursor[0] == i {
            for j, _ := range m.grid[i] {
                if m.cursor[1] == j {
                    cusor[j] = ">"
                }
            }
        }
        s += fmt.Sprintf("[%s][%s][%s]\n",cusor[0], cusor[1], cusor[2])
    }
    return s
}

func main() {
    fmt.Println("this will be a shitty little sudoku solver in go")
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Error, %v", err)
        os.Exit(1)
    }
}
