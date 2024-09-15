package listmodel

import "github.com/charmbracelet/lipgloss"

// Style for the active tab (highlighted)

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

  titleStyle = lipgloss.NewStyle().
    Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
	activeTitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFDF5")).
				Background(lipgloss.Color("#25A065")).
				Padding(0, 1).
				Bold(true)

	// Style for the inactive tabs (muted)
	inactiveTitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#A9A9A9")).
				Background(lipgloss.Color("#2D2D2D")).
				Padding(0, 1)
)

func (m Model) View() string {
	return appStyle.Render(m.list.View())
}
