package preview

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	errMsg error
)

const (
	name = iota
	username
	password
)

const (
	hotPink  = lipgloss.Color("#FF06B7")
	darkGray = lipgloss.Color("#767676")
)

type Model struct {
	inputs  []textinput.Model
	focused int
	err     error
}

func InitalModel() Model {
	var inputs []textinput.Model = make([]textinput.Model, 3)
	inputs[name] = textinput.New()
	inputs[name].Placeholder = "name"
	// inputs[name].Focus()
	inputs[name].CharLimit = 20
	inputs[name].Width = 30
	inputs[name].Prompt = ""

	inputs[username] = textinput.New()
	inputs[username].Placeholder = "username"
	inputs[username].CharLimit = 5
	inputs[username].Width = 15
	inputs[username].Prompt = ""

	inputs[password] = textinput.New()
	inputs[password].Placeholder = "password"
	inputs[password].CharLimit = 3
	inputs[password].Width = 15
	inputs[password].Prompt = ""

	return Model{
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
				return m, tea.Quit
			}
			m.nextInput()
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.nextInput()
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

var (
	labelStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFA500"))
	inputBoxStyle = lipgloss.NewStyle().Padding(0, 1).BorderForeground(lipgloss.Color("#008080")).Width(20)
	windowStyle   = lipgloss.NewStyle().
			Padding(1, 2).
			Margin(2, 0).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#5F5FD7"))
)

func (m Model) View() string {
	return windowStyle.Render(fmt.Sprintf(
		`
 %s %s

 %s %s
 %s %s
`,
		labelStyle.Render("Name:"), m.inputs[name].View(),
		labelStyle.Render("Username:"), inputBoxStyle.Render(m.inputs[username].View()),
		labelStyle.Render("Password:"), inputBoxStyle.Render(m.inputs[password].View()),
	))
}

// nextInput focuses the next input field
func (m *Model) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

// prevInput focuses the previous input field
func (m *Model) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
