package token

import "time"

// Interfave for managing tokens
type Maker interface {
	// CreateToken Creates a new token for a specific username and duration
	CreateToken(
		username string,
		role string,
		duration time.Duration,
	) (string, *Payload, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
