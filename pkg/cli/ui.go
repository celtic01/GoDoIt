package cli

import (
	"fmt"
	"os"

	"github.com/celtic01/GoDoIt/internal/todo"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor    int
	TodoList  *todo.TodoList
	textInput textinput.Model
	state     string
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

type TodosListMsg []string

func Exec() {

	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	var initModel = model{
		cursor:    0,
		TodoList:  todo.NewTodoList(),
		textInput: ti,
		state:     "list",
	}

	p := tea.NewProgram(initModel)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if m.state == "input" {
		m.textInput, cmd = m.textInput.Update(msg)
	}

	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.TodoList.Todos)-1 {
				m.cursor++
			}

		case "enter":
			if m.state == "input" {
				m.TodoList.Add(*todo.NewTodo(m.textInput.Value(), ""))
				m.TodoList.SaveTodo()
				m.state = "list"
				m.textInput.SetValue("")
			} else {
				m.TodoList.Todos[m.cursor].Completed = !m.TodoList.Todos[m.cursor].Completed
				m.TodoList.SaveTodo()
			}

		case "i":
			m.state = "input"

		case "d":
			if len(m.TodoList.Todos) > 0 {
				m.TodoList.Remove(m.TodoList.Todos[m.cursor].ID)
			}
		}

	}

	return m, cmd
}

func (m model) View() string {
	s := ""
	if m.state == "input" {
		s = m.textInput.View()
	} else {
		for i, title := range m.TodoList.GetAllTitles() {

			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			completed := " "
			if m.TodoList.Todos[i].Completed {
				completed = "x"
			}
			s += fmt.Sprintf("%s [%s] %s \n", cursor, completed, title)
		}
	}
	return s
}
