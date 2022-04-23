package models

//UNCHANGED
type PrimaryKey struct {
	Id uint `form:"id"  json:"id"  binding:"required,gt=0"`
}

type Page struct {
	PageNum  int `form:"pageNum"  json:"pageNum"  binding:"required,gt=0"`
	PageSize int `form:"pageSize" json:"pageSize" binding:"required,gt=0"`
}
