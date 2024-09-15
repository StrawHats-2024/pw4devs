package main

import (
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

type model struct {
	lists     list.Model
	keys      *listKeyMap
	activeTab Tab
	secrests  map[Tab][]list.Item
}

func (m model) Init() tea.Cmd {
	return nil
}

func newModel() model {
	items := [][]list.Item{
		{

			item{title: "Raspberry Pi’s", desc: "I have ’em all over my house"},
			item{title: "Nutella", desc: "It's good on toast"},
			item{title: "Bitter melon", desc: "It cools you down"},
		},
		{

			item{title: "Nice socks", desc: "And by that I mean socks without holes"},
			item{title: "Eight hours of sleep", desc: "I had this once"},
			item{title: "Cats", desc: "Usually"},
			item{title: "Plantasia, the album", desc: "My plants love it too"},
			item{title: "Pour over coffee", desc: "It takes forever to make though"},
		},
		{
			item{title: "VR", desc: "Virtual reality...what is there to say?"},
			item{title: "VR", desc: "Virtual reality...what is there to say?"},
			item{title: "VR", desc: "Virtual reality...what is there to say?"},
			item{title: "VR", desc: "Virtual reality...what is there to say?"},
		},
	}
	listKeys := newKeyMap()
	secretsList := list.New(items[0], list.NewDefaultDelegate(), 0, 0)
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

	return model{
		lists:     secretsList,
		keys:      listKeys,
		activeTab: Personal,
		secrests: map[Tab][]list.Item{
			Personal: items[0],
			Groups:   items[1],
			Shared:   items[2],
		},
	}
}
