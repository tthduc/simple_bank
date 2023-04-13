package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// CeateToken creates a new token for a specific username and duration
	CeateToken(username string, duration time.Duration) (string, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
