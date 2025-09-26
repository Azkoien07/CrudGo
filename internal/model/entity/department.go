package entity

type Department struct {
	Id             int    `orm:"id,primary" json:"id"`
	Name           string `orm:"name" json:"name"`
	CodeDepartment string `orm:"code_department" json:"code_department"`
	City           string `orm:"city" json:"city"`

	// Relations
}
