package resources

type TimeLog struct {
	ID          string   `jsonapi:"primary,time-logs"`
	Description string   `jsonapi:"attr,description"`
	StartsAt    string   `jsonapi:"attr,starts-at"`
	EndsAt      string   `jsonapi:"attr,ends-at"`
	CreatedAt   string   `jsonapi:"attr,created-at"`
	UpdatedAt   string   `jsonapi:"attr,updated-at"`
	User        *User    `jsonapi:"relation,user"`
	Project     *Project `jsonapi:"relation,project"`
}
