package ui

import (
	"fmt"

	"github.com/atotto/clipboard"
	tea "github.com/charmbracelet/bubbletea"
	"strawhats.pm4dev/internals/encryption"
	"strawhats.pm4dev/internals/utils"
)

func decryptedCredentails(value item, m *model, copy string) (tea.Model, tea.Cmd) {
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
			return m, m.list.NewStatusMessage(fmt.Sprintf("Error while copying password: %s", err.Error()))
		}
		return m, m.list.NewStatusMessage(fmt.Sprintf("password copied"))
	} else {
		err = clipboard.WriteAll(credentials.Username)
		if err != nil {
			return m, m.list.NewStatusMessage(fmt.Sprintf("Error while copying Username: %s", err.Error()))
		}
		return m, m.list.NewStatusMessage(fmt.Sprintf("Username copied: %s", credentials.Username))
	}
}
