package structs

type User struct {
	Name string `json:"name" xml:"name" form:"name" query:"name"`
	Age int `json:"age" xml:"age" form:"age" query:"age"`
}