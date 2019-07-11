package defs

type UserCredential struct {
	UserName string `json:"user_name"`
	Pwd string `json:"pwd"`
}

type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}
