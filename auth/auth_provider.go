package auth

// New Auth Client
func NewAuthClient() AuthClient {
	return &authClient{}
}

// Auth Keys
type AuthKeys struct {
	Url       string
	ApiKey    string
	SecretKey string
}

// Auth Client
type AuthClient interface {
	New(*string, *string, *string) *AuthKeys
}

// Auth struct
type authClient struct{}

// Provider
func (a *authClient) New(apiKey, secretKey, url *string) *AuthKeys {
	return &AuthKeys{
		Url:       *url,
		ApiKey:    *apiKey,
		SecretKey: *secretKey,
	}
}
