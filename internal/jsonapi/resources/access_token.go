package resources

type AccessToken struct {
	ID       string `jsonapi:"primary,access_tokens"`
	Email    string `jsonapi:"attr,email"`
	Password string `jsonapi:"attr,password"`
}
