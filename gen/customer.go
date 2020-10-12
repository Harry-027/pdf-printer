package gen

import (
	"fmt"
	"github.com/Harry-027/pdf-printer/models"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"sync"
)

// Generate invoice for given customer ...
func GenerateInvoicePdf(wg *sync.WaitGroup, customer models.Customer, pdfPath string) {
	defer wg.Done()
	grayColor := getGrayColor()
	redColor := getRedColor()
	title, picPath, summaryDesc, note, tableHeader := generateCustomerContent()
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 5, 10)
	m.RegisterHeader(func() {
		m.SetBackgroundColor(grayColor)
		m.Row(40, func() {
			m.Col(4, func() {
				m.Text(title, props.Text{
					Top:         10,
					Size:        35,
					Style:       consts.Bold,
					Align:       consts.Center,
					Extrapolate: false,
					Family:      consts.Courier,
				})
			})
			m.ColSpace(4)
			m.Col(4, func() {
				m.Text(customer.Company_Name, props.Text{
					Top:         15,
					Size:        10,
					Align:       consts.Right,
					Extrapolate: false,
				})
				m.Text(customer.Company_Contact, props.Text{
					Top:         25,
					Size:        8,
					Align:       consts.Right,
					Extrapolate: false,
				})
			})
		})
	})
	m.Row(20, func() {})
	m.Row(60, func() {
		m.Col(12, func() {
			_ = m.FileImage(picPath, props.Rect{
				Top:     60,
				Percent: 200,
				Center:  true,
			})
		})
	})

	m.Row(20, func() {})
	m.Row(30, func() {
		m.Col(8, func() {
			customerName := fmt.Sprintf("Dear %s,", customer.Customer_Name)
			m.Text(customerName, props.Text{
				Top:         10,
				Family:      consts.Helvetica,
				Style:       consts.Italic,
				Size:        16,
				Align:       consts.Left,
				Extrapolate: false,
			})
			summary := fmt.Sprintf("%s-%s:", summaryDesc, customer.Order_ID)
			m.Text(summary, props.Text{
				Top:         22,
				Family:      consts.Helvetica,
				Size:        12,
				Align:       consts.Left,
				Extrapolate: false,
			})
		})
	})
	header := tableHeader
	contents := [][]string{
		{customer.Order_Date, customer.Product_Name, customer.Quantity, customer.Price, customer.Discount, customer.Total},
	}
	m.Line(10)
	m.TableList(header, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      12,
			GridSizes: []uint{3, 3, 1, 2, 1, 2},
		},
		ContentProp: props.TableListContent{
			Size:      10,
			GridSizes: []uint{3, 3, 1, 2, 1, 2},
		},
		Align:              consts.Center,
		HeaderContentSpace: 5,
		Line:               true,
	})

	m.Row(20, func() {
		m.Col(12, func() {
			m.Text(customer.Company_Name, props.Text{
				Top:         40,
				Family:      consts.Helvetica,
				Style:       consts.Bold,
				Size:        12,
				Align:       consts.Right,
				Extrapolate: false,
			})
		})
	})
	m.Row(10, func() {
		m.SetBackgroundColor(redColor)
	})
	m.Row(20, func() {
		m.Col(12, func() {
			ps := fmt.Sprintf("%s %s", customer.Company_Name, note)
			m.Text(ps, props.Text{
				Top:         5,
				Family:      consts.Arial,
				Style:       consts.Italic,
				Size:        12,
				Align:       consts.Left,
				Extrapolate: false,
			})
		})
	})
	pdfPath = pdfPath + fmt.Sprintf("/%s-%s.pdf", customer.Order_ID, customer.Customer_Name)
	err := m.OutputFileAndClose(pdfPath)
	if err != nil {
		errDesc := fmt.Sprintf("Could not save invoice for: %s", customer.Customer_Name)
		fmt.Println(errDesc, err)
		return
	}
	successDesc := fmt.Sprintf("Invoice generated for customer: %s", customer.Customer_Name)
	fmt.Println(successDesc)
}
