package port

import "time"

type TokenManager interface {
    CreateToken(userID uint, duration time.Duration) (string, error)
    ParseToken(token string) (uint, error)
}
