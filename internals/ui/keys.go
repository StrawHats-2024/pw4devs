package ui

import "github.com/charmbracelet/bubbles/key"


type keyMap struct {
	copyPossword key.Binding
	copyUsername key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.copyPossword, k.copyUsername}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.copyPossword, k.copyUsername}, // second column
	}
}

func newListKeyMap() *keyMap {
	return &keyMap{
		copyPossword: key.NewBinding(
			key.WithKeys("p"),
			key.WithHelp("p", "copy password"),
		),
		copyUsername: key.NewBinding(
			key.WithKeys("u"),
			key.WithHelp("u", "copy username"),
		),
	}
}
