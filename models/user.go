package models

type Users struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string
	RoleId    uint
	Role      Role
	Bookmarks []Bookmark
	Filters   []UserFilter
}
