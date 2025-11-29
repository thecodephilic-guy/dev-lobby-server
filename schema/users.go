package schema

import "time"

type User struct {
	ID              string    `json: "id"`
	Email           string    `json: "email"`
	Password        string    `json: "password"`
	IsEmailVerified bool      `json: "isEmailVerified`
	CreatedAt       time.Time `json: "createdAt"`
}

const UsersTableDDL = `
CREATE TABLE IF NOT EXISTS users (
	id					UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	email				TEXT NOT NULL UNIQUE,
	password 			TEXT NOT NULL,
	is_email_verified	 BOOLEAN NOT NULL DEFAULT FALSE,
	created_at 			TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
`
