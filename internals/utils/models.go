package utils

import "time"

type SecretRecord struct {
	ID            int64     `json:"id"`             // Bigserial primary key
	Name          string    `json:"name"`           // Name of the secret
	EncryptedData []byte    `json:"encrypted_data"` // Encrypted credentials (bytea)
	IV            []byte    `json:"iv"`             // Initialization Vector (bytea)
	OwnerID       int64     `json:"owner_id"`       // Foreign key referencing users(id)
	CreatedAt     time.Time `db:"created_at"`       // Timestamp with time zone
}

type GroupRecordWithUsers struct {
	ID        int64          `db:"id" json:"id"`                 // Primary key
	Name      string         `db:"name" json:"name"`             // Group name (unique, not null)
	CreatorID int64          `db:"creator_id" json:"creator_id"` // Foreign key referencing users (creator)
	CreatedAt time.Time      `db:"created_at" json:"created_at"` // Timestamp when the group was created
	Users     []UserRecord   `json:"users"`
	Secrets   []SecretRecord `json:"secrets"`
}
type UserRecord struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	Activated bool      `json:"activated"`
	CreatedAt time.Time `json:"created_at"`
	Version   int       `json:"-"`
}

type GroupRecord struct {
	ID        int64     `db:"id" json:"id"`                 // Primary key
	Name      string    `db:"name" json:"name"`             // Group name (unique, not null)
	CreatorID int64     `db:"creator_id" json:"creator_id"` // Foreign key referencing users (creator)
	CreatedAt time.Time `db:"created_at" json:"created_at"` // Timestamp when the group was created
}
