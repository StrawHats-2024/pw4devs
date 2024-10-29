package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "d": // delete item
			currItem := m.list.SelectedItem()
			value, ok := currItem.(item)
			if !ok {
				return m, m.list.ToggleSpinner()
			}
			return m, m.list.NewStatusMessage(fmt.Sprintf("Selected %d: ", value.id))
		case "u": // copy username
			currItem := m.list.SelectedItem()
			value, ok := currItem.(item)
			if !ok {
				return m, m.list.NewStatusMessage("Error occured")
			}
			return decryptedCredentails(value, &m, "username")
		case "p": // copy password
			currItem := m.list.SelectedItem()
			value, ok := currItem.(item)
			if !ok {
				return m, m.list.NewStatusMessage("Error occured")
			}
			return decryptedCredentails(value, &m, "password")
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}
