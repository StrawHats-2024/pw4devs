package main

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.add):
			m.lists.NewStatusMessage(statusMessageStyle("New things"))

		case key.Matches(msg, m.keys.tab):
			switch m.activeTab {
			case Personal:
				m.activeTab = Groups
			case Groups:
				m.activeTab = Shared
			case Shared:
				m.activeTab = Personal
			}

			m.lists.Title = ""

			switch m.activeTab {
			case Personal:
				m.lists.Title = activeTitleStyle.Render("Personal") +
					" | " + inactiveTitleStyle.Render("Groups") +
					" | " + inactiveTitleStyle.Render("Shared")
				m.lists.SetItems(m.secrests[m.activeTab])
			case Groups:
				m.lists.Title = inactiveTitleStyle.Render("Personal") +
					" | " + activeTitleStyle.Render("Groups") +
					" | " + inactiveTitleStyle.Render("Shared")
				m.lists.SetItems(m.secrests[m.activeTab])
			case Shared:
				m.lists.Title = inactiveTitleStyle.Render("Personal") +
					" | " + inactiveTitleStyle.Render("Groups") +
					" | " + activeTitleStyle.Render("Shared")
				m.lists.SetItems(m.secrests[m.activeTab])
			}
		}
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		// Set width to 50% of the terminal width
		width := (msg.Width - h) / 2
		m.lists.SetSize(width, msg.Height-v)
	}

	var cmd tea.Cmd
	m.lists, cmd = m.lists.Update(msg)
	return m, cmd
}
