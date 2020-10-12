package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/Harry-027/pdf-printer/db"
	"github.com/Harry-027/pdf-printer/models"
	"github.com/Harry-027/pdf-printer/utils"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/spf13/cobra"
)

// Get the header keys based on dataType ...
func getHeaderKeys(dataType string) []string {
	var headerKeys []string
	if dataType == utils.STUDENT_TYPE {
		headerKeys = utils.STUDENT_COLUMNS
	} else if dataType == utils.EMPLOYEE_TYPE {
		headerKeys = utils.EMPLOYEE_COLUMN
	} else {
		headerKeys = utils.CUSTOMER_COLUMN
	}
	return headerKeys
}

// Render the UI ...
func setUI(headerKeys []string, rows [][]string) {
	if err := ui.Init(); err != nil {
		utils.FatalErr(err)
	}
	defer ui.Close()

	table := widgets.NewTable()
	table.Rows = [][]string{}

	table.Rows = [][]string{}
	table.Rows = append(table.Rows, headerKeys)
	for _, value := range rows {
		table.Rows = append(table.Rows, value)
	}

	table.TextStyle = ui.NewStyle(ui.ColorGreen)
	table.SetRect(0, 0, 150, 50)
	ui.Render(table)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

// Fetch the data from BoltDB & define the data structure for UI
func renderUI(records []*models.Record, dataType string) {
	var parentMap []map[string]interface{}

	for _, record := range records {
		m := map[string]interface{}{}
		err := json.Unmarshal([]byte(record.Value), &m)
		utils.FatalErr(err)
		parentMap = append(parentMap, m)
	}

	var rows [][]string
	headerKeys := getHeaderKeys(dataType)

	for _, childMap := range parentMap {
		var row []string
		for _, v := range headerKeys {
			rowValue := fmt.Sprintf("%v", childMap[v])
			row = append(row, rowValue)
		}
		rows = append(rows, row)
	}

	setUI(headerKeys, rows)
}

// Cobra verify data command ...
var verifyCmd = &cobra.Command{
	Use:   utils.VERIFY_USAGE,
	Short: utils.VERIFY_DESC,
	Run: func(cmd *cobra.Command, args []string) {
		dataType, err := cmd.Flags().GetString(utils.FLAG_TYPE)
		utils.LogErr(err)
		records := db.ViewTxn(dataType)
		renderUI(records, dataType)
	},
}

func init() {
	verifyCmd.Flags().String(utils.FLAG_TYPE, utils.FLAG_TYPE_DEFAULT, utils.FLAG_TYPE_DESC)
	RootCmd.AddCommand(verifyCmd)
}
