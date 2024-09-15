package main

import "github.com/charmbracelet/bubbles/key"

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
