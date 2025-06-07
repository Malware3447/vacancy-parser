package models

type ItemsModel struct {
	Id         int32
	Title      string
	Salary     int32
	Company    string
	Location   string
	Experience string
	SourceID   int32
	Urls       []string
}
