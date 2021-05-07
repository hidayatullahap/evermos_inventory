package auth

const (
	ContextTokenValue = "token-value"
)

type TokenPayload struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}
