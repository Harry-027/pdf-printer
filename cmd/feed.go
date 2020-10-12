package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/Harry-027/pdf-printer/db"
	"github.com/Harry-027/pdf-printer/models"
	"github.com/Harry-027/pdf-printer/utils"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// Feed the csv file data to Bolt DB respective buckets ...
func feedFile(absPath string, dataType string) {
	csvfile, err := os.Open(absPath)
	utils.FatalErr(err)
	r := csv.NewReader(csvfile) // Parse the file
	records, err := r.ReadAll()
	utils.FatalErr(err)

	switch dataType {
	case utils.STUDENT_TYPE:
		var students [][]byte
		for i := 1; i < len(records); i++ {
			student := models.Student{
				School_Name:      strings.Trim(records[i][0], " "),
				Principal_Name:   strings.Trim(records[i][1], " "),
				Student_Name:     strings.Trim(records[i][2], " "),
				Academic_Year:    strings.Trim(records[i][3], " "),
				Class:            strings.Trim(records[i][4], " "),
				English:          strings.Trim(records[i][5], " "),
				Hindi:            strings.Trim(records[i][6], " "),
				Maths:            strings.Trim(records[i][7], " "),
				Science:          strings.Trim(records[i][8], " "),
				Computer_Science: strings.Trim(records[i][9], " "),
				Result:           strings.Trim(records[i][10], " "),
			}
			studentDetails, err := json.Marshal(student)
			utils.FatalErr(err)
			students = append(students, studentDetails)
		}
		db.CreateTxn(students, utils.STUDENT_TYPE)
		fmt.Println("Data fed successfully !! Run the command 'pdf-printer verify --type=feature_type' to verify ..")
	case utils.EMPLOYEE_TYPE:
		var employees [][]byte
		for i := 1; i < len(records); i++ {
			employee := models.Employee{
				Company_Name:      strings.Trim(records[i][0], " "),
				Employee_Name:     strings.Trim(records[i][1], " "),
				Designation:       strings.Trim(records[i][2], " "),
				Pay_Period:        strings.Trim(records[i][3], " "),
				Date_Of_Payment:   strings.Trim(records[i][4], " "),
				CTC:               strings.Trim(records[i][5], " "),
				Net_Salary:        strings.Trim(records[i][6], " "),
				Work_Days:         strings.Trim(records[i][7], " "),
				Absence:           strings.Trim(records[i][8], " "),
				Basic:             strings.Trim(records[i][9], " "),
				Hra:               strings.Trim(records[i][10], " "),
				Skill_Allowance:   strings.Trim(records[i][11], " "),
				Medical_Allowance: strings.Trim(records[i][12], " "),
				PF_Contribution:   strings.Trim(records[i][13], " "),
				Income_Tax:        strings.Trim(records[i][14], " "),
				Exemption:         strings.Trim(records[i][15], " "),
				Net_Payment:       strings.Trim(records[i][16], " "),
			}
			employeeDetails, err := json.Marshal(employee)
			utils.FatalErr(err)
			employees = append(employees, employeeDetails)
		}
		db.CreateTxn(employees, utils.EMPLOYEE_TYPE)
		fmt.Println("Data fed successfully !! Run the command 'pdf-printer verify --type=feature_type' to verify ..")
	case utils.CUSTOMER_TYPE:
		var customers [][]byte
		for i := 1; i < len(records); i++ {
			customer := models.Customer{
				Company_Name:    strings.Trim(records[i][0], " "),
				Company_Contact: strings.Trim(records[i][1], " "),
				Order_ID:        strings.Trim(records[i][2], " "),
				Order_Date:      strings.Trim(records[i][3], " "),
				Customer_Name:   strings.Trim(records[i][4], " "),
				Product_Name:    strings.Trim(records[i][5], " "),
				Quantity:        strings.Trim(records[i][6], " "),
				Price:           strings.Trim(records[i][7], " "),
				Discount:        strings.Trim(records[i][8], " "),
				Total:           strings.Trim(records[i][9], " "),
			}
			customerDetails, err := json.Marshal(customer)
			utils.FatalErr(err)
			customers = append(customers, customerDetails)
		}
		db.CreateTxn(customers, utils.CUSTOMER_TYPE)
		fmt.Println("Data fed successfully !! Run the command 'pdf-printer verify --type=feature_type' to verify ..")
	default:
		fmt.Println(utils.INVALID_FEATURE)
	}
}

// Cobra feed data command ...
var feedCmd = &cobra.Command{
	Use:   utils.FEED_USAGE,
	Short: utils.FEED_DESC,
	Run: func(cmd *cobra.Command, args []string) {
		absPath, err := cmd.Flags().GetString(utils.FLAG_PATH)
		utils.LogErr(err)
		dataType, err := cmd.Flags().GetString(utils.FLAG_TYPE)
		utils.LogErr(err)
		feedFile(absPath, dataType)
	},
}

func init() {
	feedCmd.Flags().String(utils.FLAG_PATH, utils.FLAG_PATH_DEFAULT, utils.FLAG_PATH_DESC)
	feedCmd.Flags().String(utils.FLAG_TYPE, utils.FLAG_TYPE_DEFAULT, utils.FLAG_TYPE_DESC)
	RootCmd.AddCommand(feedCmd)
}
