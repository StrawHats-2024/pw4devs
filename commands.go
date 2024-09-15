package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func getPersonalSecrets() tea.Msg {
	const url = "http://localhost:8080/secrets"
	var secrets []Secret
	if err := fetchSecrets(url, &secrets); err != nil {
		return err
	}
	return personalSecretsFetchMsg(secrets)
}

func getSharedSecrets() tea.Msg {
	const url = "http://localhost:8080/groups"
	var groups []Group
	if err := fetchSecrets(url, &groups); err != nil {
		return err
	}
	return groupsFetchMsg(groups)
}

func getGroups() tea.Msg {
	const url = "http://localhost:8080/shared"
	var secrets []Secret
	if err := fetchSecrets(url, &secrets); err != nil {
		return err
	}
	return sharedSecretsFetchMsg(secrets)
}

type personalSecretsFetchMsg []Secret
type sharedSecretsFetchMsg []Secret
type groupsFetchMsg []Group

type errMsg struct{ error }

func (e errMsg) Error() string { return e.error.Error() }
