package entity

import "gorm.io/gorm"

type Employees struct {
	gorm.Model
	Name	string `vaild:"requried,stringlength(2|80)"`
	Salary  float64 `vaild:"requried~Salary must be between 15000 and 200000,range(15000|200000)"`
	EmployeesCode string `vaild:"required~EmployeeCode must be 2 uppercase English letters (A-Z) followed by - and 4 digits (0-9),uppercase({2}),matches(^[0-9]{4}$)"`
}