package ui

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/atotto/clipboard"
	tea "github.com/charmbracelet/bubbletea"
	"strawhats.pm4dev/internals/encryption"
	"strawhats.pm4dev/internals/utils"
)

func decryptedCredentailsWithCopy(value item, m *model, copy string) (tea.Model, tea.Cmd) {
	decryptedData, err := encryption.DecryptAESGCM(value.encryptedData, value.iv,
		utils.GetEncryptionKey())

	if err != nil {
		return m, m.list.NewStatusMessage(fmt.Sprintf("Error while decrypting: %s", err.Error()))
	}
	credentials, err := utils.ParseJSONToCredentials(decryptedData)
	if err != nil {
		return m, m.list.NewStatusMessage(fmt.Sprintf("Error while parsing credentials: %s", err.Error()))
	}

	if copy == "password" {
		err := clipboard.WriteAll(credentials.Password)
		if err != nil {
			return m, m.list.NewStatusMessage(
				errorStatusStyle.Render(fmt.Sprintf("Error while copying password: %s", err.Error())))
		}
		return m, m.list.NewStatusMessage(
			notificationStatusStyle.Render(fmt.Sprintf("password copied")))
	} else {
		err = clipboard.WriteAll(credentials.Username)
		if err != nil {
			return m, m.list.NewStatusMessage(
				errorStatusStyle.Render(fmt.Sprintf("Error while copying username: %s", err.Error())))
		}
		return m, m.list.NewStatusMessage(
			notificationStatusStyle.Render(fmt.Sprintf("Username copied: %s", credentials.Username)))
	}
}

func decryptedData(value item) (DecryptionResultMsg, error) {

	decryptedData, err := encryption.DecryptAESGCM(value.encryptedData, value.iv,
		utils.GetEncryptionKey())

	if err != nil {
		return DecryptionResultMsg{}, err
	}
	credentials, err := utils.ParseJSONToCredentials(decryptedData)
	if err != nil {
		return DecryptionResultMsg{}, err
	}

	return DecryptionResultMsg{title: value.title,
		username: credentials.Username,
		password: credentials.Password,
	}, nil
}

func fetchSecrets() ([]utils.SecretRecord, error) {

	type resBody struct {
		Data    []utils.SecretRecord `json:"data"`
		Message string               `json:"message"`
	}
	// Here we would normally call the logic to fetch and list secrets.
	// Currently, just validating inputs and placeholder message.
	res, err := utils.MakeRequest[resBody]("/v1/secrets/user", http.MethodGet, nil, utils.GetAuthtoken())
	if err != nil {
		return nil, err
	}
	data := res.ResBody.Data

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Fetch request failed with status code: %d", res.StatusCode)
	}
	return data, nil
}

func maskPassword(password string) string {
	return strings.Repeat("•", len(password))
}
