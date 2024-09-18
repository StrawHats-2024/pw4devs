package manager

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Pane int
type CurrentTab string

// FilterValue is the searchable value for the folder.
func (f CurrentTab) FilterValue() string {
	return string(f)
}

const (
	TabsPane Pane = iota
	SecretsPane
	PreviewPane
)

type Model struct {
	ActivePane Pane
	Loading    bool

	Tabs    list.Model
	Secrets map[CurrentTab]*list.Model
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

func newList(items []list.Item, height int) *list.Model {
	snippetList := list.New(items, list.NewDefaultDelegate(), 25, height)
	snippetList.SetShowHelp(false)
	snippetList.SetShowFilter(false)
	snippetList.SetShowTitle(false)
	snippetList.Styles.StatusBar = lipgloss.NewStyle().Margin(1, 2).Foreground(lipgloss.Color("240")).MaxWidth(35 - 2)
	snippetList.Styles.NoItems = lipgloss.NewStyle().Margin(0, 2).Foreground(lipgloss.Color("8")).MaxWidth(35 - 2)
	snippetList.FilterInput.Prompt = "Find: "
	snippetList.SetStatusBarItemName("snippet", "snippets")
	snippetList.DisableQuitKeybindings()

	return &snippetList
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if m.Loading {
			m.initModel(msg.Width, msg.Height)
			m.Loading = false
		}
	}
	var cmd tea.Cmd
	m.Tabs, cmd = m.Tabs.Update(msg)
	return m, cmd
}
func (m Model) View() string {
	if m.Loading {
		return "Loading..."
	}
	return lipgloss.JoinHorizontal(lipgloss.Left, m.Tabs.View(), m.List().View())
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
	m.Secrets = tabMap
}
