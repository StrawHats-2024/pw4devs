package main

import (
	authmodel "github.com/StrawHats-2024/pw4devs/auth_model"
	listmodel "github.com/StrawHats-2024/pw4devs/list_model"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	login bool
	list       listmodel.Model
	auth       authmodel.Model
}

func (m model) Init() tea.Cmd {
	if m.login {
		return m.list.Init()
	}
	return m.auth.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.login {
		return m.list.Update(msg)
	}
	return m.auth.Update(msg)
}

func (m model) View() string {
	if m.login {
		return m.list.View()
	}
	return m.auth.View()
}
