package manager

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if m.Loading {
			m.initModel(msg.Width, msg.Height)
			m.Loading = false
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			if m.ActivePane == SecretsPane {
				m.ActivePane = TabsPane
			} else {
				m.ActivePane = SecretsPane
			}
		}
		switch m.ActivePane {
		case TabsPane:
			m.Tabs.Styles.Title = activeTitleStyle
			m.Secrets[CurrentTab(m.selectedTab().name)].Styles.Title = inactiveTitleStyle
		case SecretsPane:
			m.Tabs.Styles.Title = inactiveTitleStyle
			m.Secrets[CurrentTab(m.selectedTab().name)].Styles.Title = activeTitleStyle
		}
	}
	var cmd tea.Cmd
	switch m.ActivePane {
	case TabsPane:
		m.Tabs, cmd = m.Tabs.Update(msg)
	case SecretsPane:
		var res list.Model
		res, cmd = m.List().Update(msg)
		m.Secrets[CurrentTab(m.selectedTab().name)] = &res
	}
	return m, cmd
}
