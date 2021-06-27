package structures

// EmailVerifierRead is read all data of email verifier
type EmailVerifierRead struct {
	Email       string `query:"email"`
	AccessToken string `query:"access_token"`
}
