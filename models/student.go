package models

import "github.com/Harry-027/pdf-printer/utils"

//Student struct ...
type Student struct {
	School_Name      string
	Principal_Name   string
	Student_Name     string
	Academic_Year    string
	Class            string
	English          string
	Hindi            string
	Maths            string
	Science          string
	Computer_Science string
	Result           string
}

var StudentBucket = []byte(utils.STUDENT_TYPE)
