package ui

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strawhats.pm4dev/internals/utils"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc       string
	id                int
	encryptedData, iv string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "d": // delete item
			currItem := m.list.SelectedItem()
			value, ok := currItem.(item)
			if !ok {
				return m, m.list.ToggleSpinner()
			}
			return m, m.list.NewStatusMessage(fmt.Sprintf("Selected %d: ", value.id))
		case "u": // copy username
			currItem := m.list.SelectedItem()
			value, ok := currItem.(item)
			if !ok {
				return m, m.list.NewStatusMessage("Error occured")
			}
			return decryptedCredentails(value, &m, "username")
		case "p": // copy password
			currItem := m.list.SelectedItem()
			value, ok := currItem.(item)
			if !ok {
				return m, m.list.NewStatusMessage("Error occured")
			}
			return decryptedCredentails(value, &m, "password")
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func Run() error {
	items := []list.Item{}
	data, err := fetchSecrets()
	if err != nil {
		return err
	}
	for _, secret := range data {
		fish := item{title: secret.Name, id: int(secret.ID),
			encryptedData: string(secret.EncryptedData), iv: string(secret.IV),
			desc: fmt.Sprintf("id: %d \t Created on %s", secret.ID,
				secret.CreatedAt.Format("Jan 02, 2005 03:04 PM")),
		}
		items = append(items, fish)
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Secrets"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		return err
	}
	return nil
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
