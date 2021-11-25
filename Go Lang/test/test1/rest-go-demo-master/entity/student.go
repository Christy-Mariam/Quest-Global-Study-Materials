package entity

//Person object for REST(CRUD)
type Person struct {
	ID        int    `json:"id"`
	Name string `json:"name"`
	Class  int `json:"class"`
	Marks      int    `json:"marks"`
}
