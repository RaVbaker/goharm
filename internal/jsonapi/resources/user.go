package resources

type User struct {
	ID               string `jsonapi:"primary,users"`
	Name             string `jsonapi:"attr,name"`
	Description      string `jsonapi:"attr,description"`
	Email            string `jsonapi:"attr,email"`
	Role             string `jsonapi:"attr,role"`
	TypeOfEmployment string `jsonapi:"attr,type-of-employment"`
	Locale           string `jsonapi:"attr,locale"`
}
