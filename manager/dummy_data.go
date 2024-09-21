package manager

import (
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/list"
)

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

// Assuming SecretType is defined as a string
type SecretType string

const (
	Personal SecretType = "Personal"
	Work     SecretType = "Work"
	Misc     SecretType = "Misc"
)

type Secret struct {
	SecretID      int        `json:"secret_id" db:"secret_id"`
	UserID        int        `json:"user_id" db:"user_id"`
	SecretType    SecretType `json:"secret_type" db:"secret_type"`
	EncryptedData string     `json:"encrypted_data" db:"encrypted_data"`
	Desc          string     `json:"description" db:"description"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
}

func (t Secret) Title() string {
	return strconv.Itoa(t.UserID)
}
func (t Secret) Description() string {
	return t.Desc
}
func (t Secret) FilterValue() string {
	return t.Desc
}

// Dummy data updated to fit the Secret structure
var dummy = map[string][]Secret{
	"Personal": {
		{
			SecretID:      1,
			UserID:        100,
			SecretType:    Personal,
			EncryptedData: "encrypted-email-password", // Placeholder for encrypted data
			Desc:          "The password for my personal email account",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			SecretID:      2,
			UserID:        100,
			SecretType:    Personal,
			EncryptedData: "encrypted-bank-pin", // Placeholder for encrypted data
			Desc:          "The PIN for my bank account",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			SecretID:      3,
			UserID:        100,
			SecretType:    Personal,
			EncryptedData: "encrypted-laptop-password", // Placeholder for encrypted data
			Desc:          "The password to unlock my laptop",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
	},
	"Work": {
		{
			SecretID:      4,
			UserID:        101,
			SecretType:    Work,
			EncryptedData: "encrypted-github-token", // Placeholder for encrypted data
			Desc:          "Token for accessing the GitHub API",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			SecretID:      5,
			UserID:        101,
			SecretType:    Work,
			EncryptedData: "encrypted-vpn-password", // Placeholder for encrypted data
			Desc:          "Password for the company's VPN",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			SecretID:      6,
			UserID:        101,
			SecretType:    Work,
			EncryptedData: "encrypted-admin-dashboard", // Placeholder for encrypted data
			Desc:          "Credentials for accessing the admin dashboard",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
	},
	"Misc": {
		{
			SecretID:      7,
			UserID:        102,
			SecretType:    Misc,
			EncryptedData: "encrypted-wifi-password", // Placeholder for encrypted data
			Desc:          "Password for my home WiFi network",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			SecretID:      8,
			UserID:        102,
			SecretType:    Misc,
			EncryptedData: "encrypted-subscription-service", // Placeholder for encrypted data
			Desc:          "Credentials for a subscription service",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			SecretID:      9,
			UserID:        102,
			SecretType:    Misc,
			EncryptedData: "encrypted-gaming-account", // Placeholder for encrypted data
			Desc:          "Password for my gaming account",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
	},
}
