package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/Harry-027/pdf-printer/db"
	"github.com/Harry-027/pdf-printer/gen"
	"github.com/Harry-027/pdf-printer/models"
	"github.com/Harry-027/pdf-printer/utils"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"math"
	"os"
	"path/filepath"
	"sync"
)

// Generate PDFs for different feature types - employee, customer, student ...
func genPdf(dataType string) {
	var wg sync.WaitGroup
	home, err := homedir.Dir() // Fetch the current user home dir.
	utils.PanicErr(err)        // Panic in case user dir not available
	pdfPath := filepath.Join(home, utils.PDFPATH)
	records := db.ViewTxn(dataType)
	totalRecords := len(records)
	limit := 30
	slots := (totalRecords / limit) + 1
	slot := 1

	switch dataType {
	case utils.STUDENT_TYPE:
		pdfPath = filepath.Join(pdfPath, utils.STUDENT_TYPE)
		err = os.MkdirAll(pdfPath, 0755)
		utils.FatalErr(err)
		for slot <= slots {
			skip := (slot - 1) * limit
			remainingRecords := totalRecords - skip
			pendingLimit := math.Min(float64(remainingRecords), float64(limit))
			upperBoundary := int(pendingLimit) + skip
			batchRecords := records[skip:upperBoundary]
			slot++
			for _, record := range batchRecords {
				studentDetails := *record
				var rec models.Student
				err := json.Unmarshal([]byte(studentDetails.Value), &rec)
				if err != nil {
					utils.LogErr(err)
					continue
				}
				wg.Add(1)
				go gen.GenerateStudentPdf(&wg, rec, pdfPath)
			}
		}
	case utils.CUSTOMER_TYPE:
		pdfPath = filepath.Join(pdfPath, utils.CUSTOMER_TYPE)
		err = os.MkdirAll(pdfPath, 0755)
		utils.FatalErr(err)
		for slot <= slots {
			skip := (slot - 1) * limit
			remainingRecords := totalRecords - skip
			pendingLimit := math.Min(float64(remainingRecords), float64(limit))
			upperBoundary := int(pendingLimit) + skip
			batchRecords := records[skip:upperBoundary]
			slot++
			for _, record := range batchRecords {
				customerDetails := *record
				var rec models.Customer
				err := json.Unmarshal([]byte(customerDetails.Value), &rec)
				if err != nil {
					utils.LogErr(err)
					continue
				}
				wg.Add(1)
				go gen.GenerateInvoicePdf(&wg, rec, pdfPath)
			}
		}
	case utils.EMPLOYEE_TYPE:
		pdfPath = filepath.Join(pdfPath, utils.EMPLOYEE_TYPE)
		err = os.MkdirAll(pdfPath, 0755)
		utils.FatalErr(err)
		for slot <= slots {
			skip := (slot - 1) * limit
			remainingRecords := totalRecords - skip
			pendingLimit := math.Min(float64(remainingRecords), float64(limit))
			upperBoundary := int(pendingLimit) + skip
			batchRecords := records[skip:upperBoundary]
			slot++
			for _, record := range batchRecords {
				employeeDetails := *record
				var rec models.Employee
				err := json.Unmarshal([]byte(employeeDetails.Value), &rec)
				if err != nil {
					utils.LogErr(err)
					continue
				}
				wg.Add(1)
				go gen.GeneratePayslipPdf(&wg, rec, pdfPath)
			}
		}
	default:
		fmt.Println(utils.INVALID_FEATURE)
	}
	wg.Wait()
}

// Cobra generate data command ...
var genCmd = &cobra.Command{
	Use:   utils.GENERATE_USAGE,
	Short: utils.GENERATE_DESC,
	Run: func(cmd *cobra.Command, args []string) {
		dataType, err := cmd.Flags().GetString(utils.FLAG_TYPE)
		utils.LogErr(err)
		genPdf(dataType)
	},
}

func init() {
	genCmd.Flags().String(utils.FLAG_TYPE, utils.FLAG_TYPE_DEFAULT, utils.FLAG_TYPE_DESC)
	RootCmd.AddCommand(genCmd)
}
