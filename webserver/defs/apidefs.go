package defs

//request
type UserCredential struct {
	Username string `json:"username"`
	Pwd string `json:"pwd"`
}

//Data model
type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}

type Comment struct {
	Id string
	VideoId string
	Author string
	Content string
}

