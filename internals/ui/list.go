package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	docStyle         = lipgloss.NewStyle().Margin(1, 2).Width(65)
	errorStatusStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#FF5555")). // Red color for errors
				Background(lipgloss.Color("#330000")). // Darker background for contrast
				Padding(0, 1)

	notificationStatusStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#00FF7F")). // Green color for notifications
				Background(lipgloss.Color("#002F1F")). // Dark green background
				Padding(0, 1)
)

type item struct {
	title, desc       string
	id                int
	encryptedData, iv string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list         list.Model
	help         help.Model
	helpKyes     keyMap
	inputs       []textinput.Model
	focusedInput int
	err          error
}

func (m model) Init() tea.Cmd {
	if len(m.list.VisibleItems()) > 0 {
		first := m.list.SelectedItem()
		value, ok := first.(item)
		if !ok {
			return m.list.NewStatusMessage(errorStatusStyle.Render("Error occured"))
		}
		return func() tea.Msg {
			data, err := decryptedData(value)
			if err != nil {
				return fmt.Errorf("Error while decryptedData")
			}
			return data
		}
	}
	return nil
}

func (m model) View() string {
	return lipgloss.JoinHorizontal(0, docStyle.Render(m.list.View()),
		lipgloss.JoinVertical(0, m.viewPreviewer(), m.help.View(m.helpKyes)))
}

func Run() error {
	items := []list.Item{}
	data, err := fetchSecrets()
	if err != nil {
		return err
	}
	for _, secret := range data {
		fish := item{title: secret.Name, id: int(secret.ID),
			encryptedData: string(secret.EncryptedData), iv: string(secret.IV),
			desc: fmt.Sprintf("id: %d  • last update: %s", secret.ID,
				secret.CreatedAt.Format("3:04PM 1/2/2006")),
		}
		items = append(items, fish)
	}
	helpModel := help.New()
	helpkeys := newListKeyMap()
	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0),
		help:     helpModel,
		helpKyes: *helpkeys,
	}
	m.InitalPreviewModel()
	m.list.Title = "Secrets"
	m.list.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			key.NewBinding(
				key.WithKeys("enter"),
				key.WithHelp("enter", "select"),
			),
		}
	}
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		return err
	}
	return nil
}
