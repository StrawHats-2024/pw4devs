package manager

import "github.com/charmbracelet/bubbles/list"

// Tabs represents a map where the key is the tab name and the value is a slice of secrets
type Secret struct {
	title string
	desc  string
}

func (s Secret) Title() string {
	return s.title
}
func (s Secret) Description() string {
	return s.desc
}
func (s Secret) FilterValue() string {
	return s.title
}

type Tab struct {
	name    string
	secrets []list.Item
}

func (t Tab) Title() string {
	return t.name
}
func (t Tab) Description() string {
	return t.name
}
func (t Tab) FilterValue() string {
	return t.name
}

func getDummyData() []Tab {
	var tabs = []Tab{}
	for key, data := range dummy {
		tmp := Tab{name: key, secrets: mapToListItems(data)}
		tabs = append(tabs, tmp)
	}
	return tabs
}

func mapToListItems(sec []Secret) []list.Item {
	tmp := []list.Item{}
	for _, s := range sec {
		tmp = append(tmp, list.Item(s))
	}
	return tmp
}

var dummy = map[string][]Secret{
	"Personal": {
		{title: "Email Password", desc: "The password for my personal email account"},
		{title: "Bank PIN", desc: "The PIN for my bank account"},
		{title: "Laptop Password", desc: "The password to unlock my laptop"},
	},
	"Work": {
		{title: "GitHub Token", desc: "Token for accessing the GitHub API"},
		{title: "VPN Password", desc: "Password for the company's VPN"},
		{title: "Admin Dashboard", desc: "Credentials for accessing the admin dashboard"},
	},
	"Misc": {
		{title: "WiFi Password", desc: "Password for my home WiFi network"},
		{title: "Subscription Service", desc: "Credentials for a subscription service"},
		{title: "Gaming Account", desc: "Password for my gaming account"},
	},
}
