package storage

var (
	userToken string
)

// SetUserToken is used to set userToken
func SetUserToken(token string) {
	userToken = token
}

// GetUserToken is used to get userToken
func GetUserToken() string {
	return userToken
}
