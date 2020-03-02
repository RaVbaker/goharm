package resources

type Project struct {
	ID       string `jsonapi:"primary,projects"`
	Name     string `jsonapi:"attr,name"`
	Color    string `jsonapi:"attr,color"`
	ClientId int    `jsonapi:"attr,client-id"`
}
