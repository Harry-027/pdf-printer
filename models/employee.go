package models

import "github.com/Harry-027/pdf-printer/utils"

//Employee struct ...
type Employee struct {
	Company_Name      string
	Employee_Name     string
	Designation       string
	Pay_Period        string
	Date_Of_Payment   string
	CTC               string
	Net_Salary        string
	Work_Days         string
	Absence           string
	Basic             string
	Hra               string
	Skill_Allowance   string
	Medical_Allowance string
	PF_Contribution   string
	Income_Tax        string
	Exemption         string
	Net_Payment       string
}

var EmployeeBucket = []byte(utils.EMPLOYEE_TYPE)
