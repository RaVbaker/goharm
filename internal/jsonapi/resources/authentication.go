package resources

type Authentication struct {
	ID          string `jsonapi:"primary,authentications"`
	AccessToken string `jsonapi:"attr,access-token"`
	UserRole    string `jsonapi:"attr,user-role"`
	UserId      string `jsonapi:"attr,user-id"`
}
