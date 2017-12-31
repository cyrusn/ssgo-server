package auth

type privateKey []byte
type contextKey string

var (
	key = []byte("secret")
	// ContextKeyClaim is the key in context for retreive the information of claim
	contextKeyName = contextKey("claims")
	roleKeyName    = "jwt-token"
)

// SetPrivateKey set the privateKey for authentication
func SetPrivateKey(newKey string) {
	key = []byte(newKey)
}

// SetContextKey set the name of ContextKeyClaim
func SetContextKey(name string) {
	contextKeyName = contextKey(name)
}

// SetRoleKeyName set the name of Role in jwt payload
func SetRoleKeyName(name string) {
	roleKeyName = name
}
