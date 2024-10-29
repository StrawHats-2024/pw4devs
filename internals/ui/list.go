package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strawhats.pm4dev/internals/ui/preview"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2).Width(55)

type item struct {
	title, desc       string
	id                int
	encryptedData, iv string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list      list.Model
	previewer preview.Model
	help      help.Model
	helpKyes  keyMap
	secrets   []item
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	return lipgloss.JoinHorizontal(0, docStyle.Render(m.list.View()),
		lipgloss.JoinVertical(0, m.previewer.View(), m.help.View(m.helpKyes)))
}

func Run() error {
	items := []list.Item{}
	data, err := fetchSecrets()
	if err != nil {
		return err
	}
	secrets := []item{}
	for _, secret := range data {
		fish := item{title: secret.Name, id: int(secret.ID),
			encryptedData: string(secret.EncryptedData), iv: string(secret.IV),
			desc: fmt.Sprintf("id: %d  • last update: %s", secret.ID,
				secret.CreatedAt.Format("3:04PM 1/2/2006")),
		}
		secrets = append(secrets, fish)
		items = append(items, fish)
	}
	helpModel := help.New()
	helpkeys := newListKeyMap()
	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0),
		previewer: preview.InitalModel(),
		help:      helpModel,
		helpKyes:  *helpkeys,
		secrets:   secrets,
	}
	m.list.Title = "Secrets"
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		return err
	}
	return nil
}
