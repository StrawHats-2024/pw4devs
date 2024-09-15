package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
)

type item struct {
	title string
	desc  string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list list.Model
	keys *listKeyMap
}

func (m model) Init() tea.Cmd {
	return nil
}

type listKeyMap struct {
	add  key.Binding
	edit key.Binding
	tab  key.Binding
}

func newKeyMap() *listKeyMap {
	return &listKeyMap{
		add: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "add"),
		),
		edit: key.NewBinding(
			key.WithKeys("e"),
			key.WithHelp("e", "edit"),
		),
		tab: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "switch tabs"),
		),
	}
}

func newModel() model {
	items := []list.Item{
		item{title: "Raspberry Pi’s", desc: "I have ’em all over my house"},
		item{title: "Nutella", desc: "It's good on toast"},
		item{title: "Bitter melon", desc: "It cools you down"},
		item{title: "Nice socks", desc: "And by that I mean socks without holes"},
		item{title: "Eight hours of sleep", desc: "I had this once"},
		item{title: "Cats", desc: "Usually"},
		item{title: "Plantasia, the album", desc: "My plants love it too"},
		item{title: "Pour over coffee", desc: "It takes forever to make though"},
		item{title: "VR", desc: "Virtual reality...what is there to say?"},
		item{title: "VR", desc: "Virtual reality...what is there to say?"},
		item{title: "VR", desc: "Virtual reality...what is there to say?"},
		item{title: "VR", desc: "Virtual reality...what is there to say?"},
	}
	listKeys := newKeyMap()

	secretsList := list.New(items, list.NewDefaultDelegate(), 0, 0)
	secretsList.Title = "Secrets"
	secretsList.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.add,
			listKeys.edit,
			listKeys.tab,
		}
	}

	return model{
		list: secretsList,
		keys: listKeys,
	}
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.add):
			m.list.NewStatusMessage(statusMessageStyle("New things"))
		}
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		// Set width to 50% of the terminal width
		width := (msg.Width - h) / 2
		m.list.SetSize(width, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return appStyle.Render(m.list.View())
}

func main() {
	if _, err := tea.NewProgram(newModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("err :", err)
		os.Exit(1)
	}
}
