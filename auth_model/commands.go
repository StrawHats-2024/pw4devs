package authmodel

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func performLogin(email, password string) tea.Msg {
	time.Sleep(2 * time.Second)
	err := Login(email, password)
	if err != nil {
		return errMsg{err}
	}
	return Success
}
