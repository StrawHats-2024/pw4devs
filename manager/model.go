package manager

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type Pane int
type CurrentTab string

// FilterValue is the searchable value for the folder.
func (f CurrentTab) FilterValue() string {
	return string(f)
}

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	inactiveTitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#A8A8A8")). // A softer, muted gray color for text
				Background(lipgloss.Color("#1F1F1F")). // A darker, neutral background
				Padding(0, 1)

	activeTitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFDF5")).
				Background(lipgloss.Color("#25A065")).
				Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
)

const (
	TabsPane Pane = iota
	SecretsPane
)

type Model struct {
	ActivePane Pane
	Loading    bool

	Tabs    list.Model
	Secrets map[CurrentTab]*list.Model
	Preview string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) selectedTab() Tab {
	item := m.Tabs.SelectedItem()
	if item == nil {
		return Tab{}
	}
	return item.(Tab)
}

// List returns the active list.
func (m *Model) List() *list.Model {
	return m.Secrets[CurrentTab(m.selectedTab().name)]
}

func (m *Model) getPreviewString() string {
	// TODO: give preview of the secret
	activeList := m.List()
	secret := activeList.SelectedItem().(Secret)
	return secret.Desc
}
func (m *Model) getTestPreview() string {
	// Example secret data (you can replace this with dynamic content)
	secretTitle := "Email Password"
	username := "user@example.com"
	password := "password123"

	// Markdown content with secret title, login ID, password, and help keys
	in := fmt.Sprintf(`# %s

- **Login ID**: %s 
- **Password**: %s  

### Actions:
- Press **'c'** to copy the Login ID (username)
- Press **'C'** to copy the Password
- Press **'e'** to edit the secret

`, secretTitle, username, password)

	// Render the markdown using Glamour
	out, err := glamour.Render(in, "dark")
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func newList(items []list.Item, height int) *list.Model {
	secretsList := list.New(items, list.NewDefaultDelegate(), 25, height)
	secretsList.SetShowHelp(false)
	// secretsList.SetShowFilter(false)
	secretsList.SetShowTitle(true)
	secretsList.Styles.StatusBar = lipgloss.NewStyle().Margin(1, 2).Foreground(lipgloss.Color("240")).MaxWidth(35 - 2)
	secretsList.Styles.NoItems = lipgloss.NewStyle().Margin(0, 2).Foreground(lipgloss.Color("8")).MaxWidth(35 - 2)
	secretsList.FilterInput.Prompt = "Find: "
	secretsList.Title = "Secrets"
	secretsList.Styles.Title = inactiveTitleStyle
	secretsList.SetStatusBarItemName("snippet", "snippets")
	// secretsList.DisableQuitKeybindings()

	return &secretsList
}

func (m Model) View() string {
	if m.Loading {
		return "Loading..."
	}
	return lipgloss.JoinHorizontal(lipgloss.Left, m.Tabs.View(), m.List().View(), "\n", m.getTestPreview())
}

func (m *Model) initModel(width, height int) {
	data := getDummyData()
	m.ActivePane = TabsPane
	tmp := []list.Item{}
	tabMap := map[CurrentTab]*list.Model{}
	for _, tab := range data {
		tmp = append(tmp, list.Item(tab))
		tabMap[CurrentTab(tab.name)] = newList(tab.secrets, 20)
	}
	m.Tabs = list.New(tmp, list.NewDefaultDelegate(), width, height)
	m.Tabs.Title = "Tabs"
	m.Tabs.Styles.Title = activeTitleStyle
	m.Tabs.SetShowStatusBar(false)
	m.Tabs.SetFilteringEnabled(false)
	// m.Tabs.SetShowTitle(false)
	m.Secrets = tabMap
}
