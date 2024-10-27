package utils

import "time"

type SecretRecord struct {
	ID            int64     `json:"id"`             // Bigserial primary key
	Name          string    `json:"name"`           // Name of the secret
	EncryptedData []byte    `json:"encrypted_data"` // Encrypted credentials (bytea)
	IV            []byte    `json:"iv"`             // Initialization Vector (bytea)
	OwnerID       int64     `json:"owner_id"`       // Foreign key referencing users(id)
	CreatedAt     time.Time `db:"created_at"`     // Timestamp with time zone
}
