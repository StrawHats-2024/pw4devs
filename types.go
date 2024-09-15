package main

import (
	"time"
)

type SecretType string

const (
	SSHKey   SecretType = "ssh_key"
	Password SecretType = "password"
	APIKey   SecretType = "api_key"
)

type Secret struct {
	SecretID       int        `json:"secret_id" db:"secret_id"`
	UserID         int        `json:"user_id" db:"user_id"`
	SecretType     SecretType `json:"secret_type" db:"secret_type"`
	SecretName     string     `json:"secret_name"`
	SecretUsername string     `json:"secret_username"`
	EncryptedData  string     `json:"encrypted_data" db:"encrypted_data"`
	Description    string     `json:"description" db:"description"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
}

type Group struct {
	GroupID   int    `json:"group_id"`
	GroupName string `json:"group_name"`
	Role      string `json:"role"`
}
