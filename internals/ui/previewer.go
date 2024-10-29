package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

type (
	errMsg error
)

const (
	nameInput = iota
	usernameInput
	passwordInput
)

const (
	hotPink  = lipgloss.Color("#FF06B7")
	darkGray = lipgloss.Color("#767676")
)

func (m *model) InitalPreviewModel() {
	var inputs []textinput.Model = make([]textinput.Model, 3)
	inputs[nameInput] = textinput.New()
	inputs[nameInput].Placeholder = "name"
	inputs[nameInput].SetValue("test name")
	// inputs[name].Focus()
	inputs[nameInput].CharLimit = 20
	inputs[nameInput].Width = 30
	inputs[nameInput].Prompt = ""

	inputs[usernameInput] = textinput.New()
	inputs[usernameInput].Placeholder = "username"
	inputs[usernameInput].SetValue("test username")
	inputs[usernameInput].CharLimit = 5
	inputs[usernameInput].Width = 15
	inputs[usernameInput].Prompt = ""

	inputs[passwordInput] = textinput.New()
	inputs[passwordInput].Placeholder = "password"
	inputs[passwordInput].SetValue("test password")
	inputs[passwordInput].CharLimit = 3
	inputs[passwordInput].Width = 15
	inputs[passwordInput].Prompt = ""
	m.inputs = inputs
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

func (m model) viewPreviewer() string {
	return windowStyle.Render(fmt.Sprintf(
		`
 %s %s

 %s %s
 %s %s
`,
		labelStyle.Render("Name:"), m.inputs[nameInput].View(),
		labelStyle.Render("Username:"), inputBoxStyle.Render(m.inputs[usernameInput].View()),
		labelStyle.Render("Password:"), inputBoxStyle.Render(maskPassword(m.inputs[passwordInput].View())),
	))
}
