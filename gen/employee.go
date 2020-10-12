package gen

import (
	"fmt"
	"github.com/Harry-027/pdf-printer/models"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"sync"
)

// Generate payslip for given employee ...
func GeneratePayslipPdf(wg *sync.WaitGroup, employee models.Employee, pdfPath string) {
	defer wg.Done()
	grayColor := getGrayColor()
	title, employerName, employeeName, payPeriod, payDate, designation, workDays,
		ctc, netSalary, absence, signature, tableHeader, contents := generatePayslipContent(employee)
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 5, 10)
	m.RegisterHeader(func() {
		m.Row(30, func() {
			m.Col(6, func() {
				m.Text(title, props.Text{
					Top:         0,
					Size:        50,
					Style:       consts.Bold,
					Align:       consts.Left,
					Extrapolate: false,
					Family:      consts.Courier,
				})
			})
		})
	})
	m.Row(10, func() {
		m.Col(6, func() {
			m.Text(employerName, props.Text{
				Top:         5,
				Family:      consts.Helvetica,
				Size:        12,
				Align:       consts.Left,
				Extrapolate: false,
			})
		})
		m.Col(6, func() {
			m.Text(payPeriod, props.Text{
				Top:         5,
				Family:      consts.Helvetica,
				Size:        12,
				Align:       consts.Right,
				Extrapolate: false,
			})
		})
	})
	m.Row(10, func() {
		m.Col(6, func() {
			m.Text(employeeName, props.Text{
				Top:         2,
				Family:      consts.Helvetica,
				Size:        12,
				Align:       consts.Left,
				Extrapolate: false,
			})
		})
		m.Col(6, func() {
			m.Text(payDate, props.Text{
				Top:         2,
				Family:      consts.Helvetica,
				Size:        12,
				Align:       consts.Right,
				Extrapolate: false,
			})
		})
	})
	m.Line(10)

	m.Row(8, func() {
		m.Col(6, func() {
			m.Text(designation, props.Text{
				Size:        10,
				Align:       consts.Left,
				Extrapolate: false,
			})
		})
		m.Col(6, func() {
			m.Text(workDays, props.Text{
				Size:        10,
				Align:       consts.Right,
				Extrapolate: false,
			})
		})
	})
	m.Row(5, func() {
		m.Col(4, func() {
			m.Text(ctc, props.Text{
				Size:        10,
				Align:       consts.Left,
				Extrapolate: false,
			})
		})
		m.Col(4, func() {
			m.Text(netSalary, props.Text{
				Size:        10,
				Align:       consts.Center,
				Extrapolate: false,
			})
		})
		m.Col(4, func() {
			m.Text(absence, props.Text{
				Size:        10,
				Align:       consts.Right,
				Extrapolate: false,
			})
		})
	})
	m.Line(10)
	m.Row(20, func() {})
	m.TableList(tableHeader, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      12,
			GridSizes: []uint{6, 6},
		},
		ContentProp: props.TableListContent{
			Size:      10,
			GridSizes: []uint{6, 6},
		},
		Align:                consts.Center,
		HeaderContentSpace:   5,
		AlternatedBackground: &grayColor,
		Line:                 true,
	})
	m.Row(40, func() {

	})
	m.Row(20, func() {
		m.ColSpace(6)
		m.Col(6, func() {
			m.Text(employee.Company_Name, props.Text{
				Top:         8,
				Family:      consts.Courier,
				Size:        8,
				Align:       consts.Center,
				Extrapolate: false,
			})
			m.Signature(signature, props.Font{
				Family: consts.Helvetica,
				Style:  consts.Italic,
				Size:   12,
			})
		})
	})
	pdfPath = pdfPath + fmt.Sprintf("/%s-%s.pdf", employee.Employee_Name, employee.Pay_Period)
	err := m.OutputFileAndClose(pdfPath)
	if err != nil {
		errDesc := fmt.Sprintf("Could not save payslip for: %s", employee.Employee_Name)
		fmt.Println(errDesc, err)
		return
	}
	successDesc := fmt.Sprintf("Payslip generated for employee: %s", employee.Employee_Name)
	fmt.Println(successDesc)
}
