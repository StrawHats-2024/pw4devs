package listmodel

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

var ErrFileNotFound = errors.New("file not found")

func ReadFileContent(filePath string) (string, error) {
	// Read the entire content of the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		// Check if the error is because the file doesn't exist
		if os.IsNotExist(err) {
			return "", ErrFileNotFound
		}
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	// Convert the content to a string and return
	return string(content), nil
}
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
