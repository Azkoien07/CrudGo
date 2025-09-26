package dto

type departmentDto struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	CodeDepartment string `json:"code_department"`
	City           string `json:"city"`

	// Relations
}
