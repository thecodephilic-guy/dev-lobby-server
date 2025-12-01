package models

import "time"

// Verification model
type Verification struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	Code      string    `json:"-"` // Don't expose code in JSON
	Attempts  int       `json:"attempts"`
	ExpiresAt time.Time `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
}

// DDL for verifications table
const VerificationsTableDDL = `
CREATE TABLE IF NOT EXISTS verifications (
	id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	code        VARCHAR(6) NOT NULL,
	attempts    INTEGER NOT NULL DEFAULT 0,
	expires_at  TIMESTAMPTZ NOT NULL,
	created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	
	-- Ensure only one active verification per user (optional but recommended)
	CONSTRAINT unique_user_verification UNIQUE (user_id)
);

-- Create index for cleanup queries
CREATE INDEX IF NOT EXISTS idx_verifications_expires_at ON verifications(expires_at);
`
