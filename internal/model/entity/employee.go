package entity

type Employee struct {
	Id    int    `orm:"id,primary" json:"id"`
	Name  string `orm:"name" json:"name"`
	Age   int    `orm:"age" json:"age"`
	Email string `orm:"email" json:"email"`
	Phone string `orm:"phone" json:"phone"`

	// Relations
}
