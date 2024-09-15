package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func mapSecretToListItem(secrets []Secret) []list.Item {
	list := make([]list.Item, len(secrets))
	for i, secret := range secrets {
		list[i] = item{title: secret.SecretName, desc: secret.Description}
	}
	return list
}

func mapGroupToListItem(groups []Group) []list.Item {
	list := make([]list.Item, len(groups))
	for i, group := range groups {
		list[i] = item{title: group.GroupName, desc: group.Role}
	}
	return list
}


func fetchSecrets(url string, target interface{}) tea.Msg {
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := c.Get(url)
	if err != nil {
		return errMsg{err}
	}
	defer res.Body.Close() // nolint:errcheck

	err = json.NewDecoder(res.Body).Decode(target)
	if err != nil {
		return errMsg{err}
	}

	return nil // Return nil to indicate success; the caller should handle it
}
