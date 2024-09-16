package listmodel

import (
	"log"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type item struct {
	title string
	desc  string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type Tab int

const (
	Personal Tab = iota
	Groups
	Shared
)

type Model struct {
	list            list.Model
	keys            *listKeyMap
	activeTab       Tab
	personalSecrets []Secret
	SharedSecrets   []Secret
	groups          []Group
	loading         bool
	err             errMsg
}

func (m Model) Init() tea.Cmd {
	return func() tea.Msg {
		secrets, err := getSecrets()
		if err != nil {
			log.Fatal("Failed to fetch: ", err.Error())
		}
		return personalSecretsFetchMsg(secrets)
	}
}

func NewModel() Model {

	listKeys := newKeyMap()
	secretsList := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	secretsList.Title = activeTitleStyle.Render("Personal") +
		" | " + inactiveTitleStyle.Render("Groups") +
		" | " + inactiveTitleStyle.Render("Shared")
	secretsList.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.add,
			listKeys.edit,
			listKeys.tab,
		}
	}

	return Model{
		list:      secretsList,
		keys:      listKeys,
		activeTab: Personal,
		loading:   true,
	}
}
