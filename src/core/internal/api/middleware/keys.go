// middleware/keys.go
package middleware

type ContextKey string

const (
	UserIDKey ContextKey = "user_id"
	// UsernameKey ContextKey = "username"
)
