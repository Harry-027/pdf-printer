package gen

import (
	"fmt"
	"github.com/Harry-027/pdf-printer/models"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"sync"
)

// Generate report card for given student ...
func GenerateStudentPdf(wg *sync.WaitGroup, student models.Student, pdfPath string) {
	defer wg.Done()
	headerColor := getHeaderColor()
	title, logoPath, studentName, academicYear, principal, tableHeader, subjects := generateStudentContent()
	m := pdf.NewMaroto(consts.Landscape, consts.A4)
	m.SetPageMargins(10, 5, 10)
	m.RegisterHeader(func() {
		m.SetBackgroundColor(headerColor)
		m.Row(18, func() {
			m.Col(12, func() {
				m.Text(title, props.Text{
					Top:         0,
					Size:        25,
					Style:       consts.BoldItalic,
					Align:       consts.Center,
					Extrapolate: true,
					Family:      consts.Helvetica,
				})
			})
		})
		m.Row(25, func() {
			m.Col(12, func() {
				_ = m.FileImage(logoPath, props.Rect{
					Center:  true,
					Percent: 80,
					Top:     0,
				})
			})
		})
		m.Row(18, func() {
			m.Col(12, func() {
				m.Text(student.School_Name, props.Text{
					Top:   0,
					Size:  25,
					Style: consts.BoldItalic,
					Align: consts.Center,
				})
			})
		})
	})

	m.Row(10, func() {
		m.Col(2, func() {
			m.Text(studentName, props.Text{
				Top:   40,
				Style: consts.Bold,
				Size:  15,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text(student.Student_Name, props.Text{
				Top:   20,
				Style: consts.Normal,
				Size:  15,
				Align: consts.Left,
			})
		})
	})
	m.Line(15)
	m.Row(10, func() {
		m.Col(2, func() {
			m.Text(academicYear, props.Text{
				Top:   2,
				Style: consts.Bold,
				Size:  15,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text(student.Academic_Year, props.Text{
				Top:   2,
				Style: consts.Normal,
				Size:  15,
				Align: consts.Left,
			})
		})
	})
	m.Line(1)

	header := tableHeader
	contents := [][]string{
		{subjects[0], student.Science},
		{subjects[1], student.Maths},
		{subjects[2], student.Hindi},
		{subjects[3], student.English},
		{subjects[4], student.Computer_Science},
	}

	m.Row(10, func() {})

	m.TableList(header, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      20,
			GridSizes: []uint{6, 6},
		},
		ContentProp: props.TableListContent{
			Size:      16,
			GridSizes: []uint{6, 6},
		},
		Align:                consts.Center,
		AlternatedBackground: &headerColor,
		HeaderContentSpace:   5,
		Line:                 false,
	})
	m.Row(5, func() {})
	m.Row(8, func() {
		m.SetBackgroundColor(getResultStatusColor(student.Result))
		m.ColSpace(4)
		m.Col(4, func() {
			status := fmt.Sprintf("RESULT : %s", student.Result)
			m.Text(status, props.Text{
				Top:   2,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Center,
			})
		})
		m.ColSpace(4)
	})
	m.Row(2, func() {
		m.SetBackgroundColor(color.Color{
			Red:   255,
			Green: 255,
			Blue:  255,
		})
	})
	m.Row(10, func() {
		m.ColSpace(8)
		m.Col(4, func() {
			m.Text(student.Principal_Name, props.Text{
				Top:    4,
				Family: consts.Courier,
				Style:  consts.Italic,
				Size:   8,
				Align:  consts.Center,
			})
			m.Signature(principal, props.Font{
				Family: consts.Helvetica,
				Style:  consts.BoldItalic,
				Size:   8,
			})
		})
	})
	pdfPath = pdfPath + fmt.Sprintf("/%s-%s.pdf", student.Student_Name, student.Academic_Year)
	err := m.OutputFileAndClose(pdfPath)
	if err != nil {
		errDesc := fmt.Sprintf("Could not save report card for: %s", student.Student_Name)
		fmt.Println(errDesc, err)
		return
	}
	successDesc := fmt.Sprintf("Report card generated for student: %s", student.Student_Name)
	fmt.Println(successDesc)
}
