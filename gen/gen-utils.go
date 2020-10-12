package gen

import (
	"fmt"
	"github.com/Harry-027/pdf-printer/models"
	"github.com/johnfercher/maroto/pkg/color"
)

func getGrayColor() color.Color {
	return color.Color{
		Red:   200,
		Green: 200,
		Blue:  200,
	}
}

func getRedColor() color.Color {
	return color.Color{
		Red:   200,
		Green: 0,
		Blue:  0,
	}
}

func getHeaderColor() color.Color {
	return color.Color{
		Red:   41,
		Green: 98,
		Blue:  255,
	}
}

func getResultStatusColor(status string) color.Color {
	var colour color.Color
	if status == "PASS" {
		colour = color.Color{
			Red:   0,
			Green: 255,
			Blue:  0,
		}
	} else {
		colour = color.Color{
			Red:   255,
			Green: 0,
			Blue:  0,
		}
	}
	return colour
}

func generateStudentContent() (title, logoPath, studentName, academicYear, principal string, tableHeader, subjects []string) {
	title = "School Report Card"
	logoPath = "./assets/images/school_logo.png"
	studentName = "Student Name :"
	academicYear = "Academic Year :"
	principal = "Principal"
	tableHeader = []string{"Subject", "Marks"}
	subjects = []string{"Science", "Maths", "Hindi", "English", "Computer Science"}
	return
}

func generateCustomerContent() (title, picPath, summaryDesc, note string, tableHeader []string) {
	title = "INVOICE"
	picPath = "./assets/images/shopping.jpg"
	summaryDesc = "Here's the Summary for your order Id "
	note = "never contacts its customer to share personal/financial information for contest or to make payments." +
		"Ignore such fraudulent calls."
	tableHeader = []string{"Order Date", "Product Name", "Quantity", "Price", "Discount", "Total"}
	return
}

func generatePayslipContent(employee models.Employee) (title, employerName, employeeName, payPeriod, payDate, designation, workDays,
	ctc, netSalary, absence, signature string, tableHeader []string, contents [][]string) {
	title = "PAYSLIP"
	employerName = fmt.Sprintf("Employer's Name: %s", employee.Company_Name)
	employeeName = fmt.Sprintf("Employee's Name: %s", employee.Employee_Name)
	payPeriod = fmt.Sprintf("Pay period: %s", employee.Pay_Period)
	payDate = fmt.Sprintf("Date of payment: %s", employee.Date_Of_Payment)
	designation = fmt.Sprintf("Designation: %s", employee.Designation)
	workDays = fmt.Sprintf("Work Days: %s", employee.Work_Days)
	ctc = fmt.Sprintf("CTC: %s", employee.CTC)
	netSalary = fmt.Sprintf("Net Salary: %s", employee.Net_Salary)
	absence = fmt.Sprintf("Absence: %s", employee.Absence)
	signature = "Employer's Signature"
	tableHeader = []string{"Earnings", "Rate (Rs/-)"}
	contents = [][]string{
		{"Basic", employee.Basic},
		{"HRA", employee.Hra},
		{"Skill Allowance", employee.Skill_Allowance},
		{"Medical Allowance", employee.Medical_Allowance},
		{"PF Contribution", employee.PF_Contribution},
		{"Income Tax", employee.Income_Tax},
		{"Exemption", employee.Exemption},
		{"Net Payment", employee.Net_Payment},
	}
	return
}
