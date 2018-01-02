package auth

type contextKey string

var (
	privateKey = []byte("secret")
	// contextKeyClaim is the key in context for retreive the information of claim
	contextKeyName = contextKey("context-claims")
	roleKeyName    = "Role"
	jwtKeyName     = "jwt-token"
)

// SetPrivateKey set the privateKey for authentication
// the default value of privateKey is secret
func SetPrivateKey(key string) {
	privateKey = []byte(key)
}

// SetContextKey set the name of ContextKeyClaim
// the default value of contextKey is "context-claims"
func SetContextKey(name string) {
	contextKeyName = contextKey(name)
}

// SetRoleKeyName set the name of Role in jwt payload
// the default value of roleKey is "context-claims"
func SetRoleKeyName(name string) {
	roleKeyName = name
}

func SetJWTKeyName(name string) {
	jwtKeyName = name
}
