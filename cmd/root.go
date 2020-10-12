package cmd

import (
	"github.com/spf13/cobra"
)

// Cobra root command ...
var RootCmd = &cobra.Command{
	Use: "pdf-printer",
	Short: "pdf-printer is a CLI client to generate pdf reports in bulk.Feed the csv data & generate the pdf reports.Following features available ::\n\n" +
		"1) Students report card \n" +
		"2) Employees Payslip \n" +
		"3) Customers Invoice \n",
}
