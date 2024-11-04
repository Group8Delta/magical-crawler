package models

type Users struct {
	Id        int64
	FirstName string
	LastName  string
	Email     string
	Role_id   int64
	Role      Roles
	BookMarks []BookMarks
}
