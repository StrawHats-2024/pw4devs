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
			return decryptedCredentailsWithCopy(value, &m, "username")
		case "p": // copy password
			currItem := m.list.SelectedItem()
			value, ok := currItem.(item)
			if !ok {
				return m, m.list.NewStatusMessage("Error occured")
			}
			return decryptedCredentailsWithCopy(value, &m, "password")
		case "enter":
			visibleItems := m.list.VisibleItems()
			if len(visibleItems) > 0 {
				currItem := m.list.SelectedItem()
				value, ok := currItem.(item)
				if !ok {
					return m, m.list.NewStatusMessage("Error occurred")
				}
				return m, func() tea.Msg {
					data, err := decryptedData(value)
					if err != nil {
						return fmt.Errorf("Error while decryptedData")
					}
					return data
				}
			}

		}
	case DecryptionResultMsg:
		m.inputs[nameInput].SetValue(msg.title)
		m.inputs[usernameInput].SetValue(msg.username)
		m.inputs[passwordInput].SetValue(msg.password)

	case errMsg:
		m.list.NewStatusMessage(msg.Error())

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

type DecryptionResultMsg struct {
	title    string
	username string
	password string
	err      error
}
