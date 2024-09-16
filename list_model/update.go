package listmodel

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Don't match any of the keys below if we're actively filtering.
		if m.list.FilterState() == list.Filtering {
			break
		}

		switch {
		case key.Matches(msg, m.keys.add):
			m.list.NewStatusMessage(statusMessageStyle("New things"))
			return m, nil

		case key.Matches(msg, m.keys.edit):
			cmd := m.list.ToggleSpinner()
			return m, cmd
		case key.Matches(msg, m.keys.tab):

			switch m.activeTab {
			case Personal:
				m.activeTab = Groups
			case Groups:
				m.activeTab = Shared
			case Shared:
				m.activeTab = Personal
			}

			m.list.Title = ""

			switch m.activeTab {
			case Personal:
				m.list.Title = activeTitleStyle.Render("Personal") +
					" | " + inactiveTitleStyle.Render("Groups") +
					" | " + inactiveTitleStyle.Render("Shared")
			case Groups:
				m.list.Title = inactiveTitleStyle.Render("Personal") +
					" | " + activeTitleStyle.Render("Groups") +
					" | " + inactiveTitleStyle.Render("Shared")
			case Shared:
				m.list.Title = inactiveTitleStyle.Render("Personal") +
					" | " + inactiveTitleStyle.Render("Groups") +
					" | " + activeTitleStyle.Render("Shared")
			}
			switch m.activeTab {
			case Personal:
				cmd := m.list.SetItems(mapSecretToListItem(m.personalSecrets))
				return m, cmd
			case Shared:
				cmd := m.list.SetItems(mapSecretToListItem(m.SharedSecrets))
				return m, cmd
			case Groups:
				cmd := m.list.SetItems(mapGroupToListItem(m.groups))
				return m, cmd
			}
		}

	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		// Set width to 50% of the terminal width
		width := (msg.Width - h) / 2
		m.list.SetSize(width, msg.Height-v)
		return m, nil

	case errMsg:
		// There was an error. Note it in the model. And tell the runtime
		// we're done and want to quit.
		m.err = msg
		return m, tea.Quit

	case personalSecretsFetchMsg:
		m.personalSecrets = msg
		return m, nil

	case groupsFetchMsg:
		m.groups = msg
		return m, nil

	case sharedSecretsFetchMsg:
		m.SharedSecrets = msg
		return m, nil
		// Assign the secrets list to the correct slice of list.Items
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}
