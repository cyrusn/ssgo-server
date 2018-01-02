package auth

type contextKey string

var (
	privateKey = []byte("secret")
	// contextKeyClaim is the keyName in http.Context which store information
	// jwt.Claims
	contextKeyName = contextKey("context-claims")
	jwtKeyName     = "jwt-token"
	roleKeyName    = "Role"
)

// SetPrivateKey set the privateKey for authentication
// the default value of privateKey is secret
func SetPrivateKey(key string) {
	privateKey = []byte(key)
}

// SetContextKeyName set the name of ContextKeyClaim,
// the default value of contextKey is "context-claims"
func SetContextKeyName(name string) {
	contextKeyName = contextKey(name)
}

// SetJWTKeyName set the name of Role in jwt payload,
// the default value of roleKey is "jwt-token"
func SetJWTKeyName(name string) {
	jwtKeyName = name
}

// SetRoleKeyName declares rolekeyName in jwt.Claims. auth package will
// get the value with the key named by rolekeyName to valdate the scope
// of authentication, the default value of roleKeyname is "Role"
func SetRoleKeyName(name string) {
	roleKeyName = name
}
