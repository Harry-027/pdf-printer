package utils

// App constants ...
const (
	DB_FILE           = "pdf-printer-db"
	STUDENT_TYPE      = "student"
	EMPLOYEE_TYPE     = "employee"
	CUSTOMER_TYPE     = "customer"
	GENERATE_USAGE    = "gen"
	GENERATE_DESC     = "Generate the pdfs for fed data of given type. Do provide the type flag to determine the type of data"
	VERIFY_USAGE      = "verify"
	VERIFY_DESC       = "Visualise the data fed in system.Do provide the type flag to determine the type of data"
	FEED_USAGE        = "feed"
	FEED_DESC         = "Feeds data for the given csv file. Do provide the feature type & absolute path to csv file."
	FLAG_PATH         = "path"
	FLAG_TYPE         = "type"
	FLAG_PATH_DESC    = "absolute path to the csv file"
	FLAG_TYPE_DESC    = "csv data type - student, employee, customer"
	FLAG_PATH_DEFAULT = "./assets/csv/students.csv"
	FLAG_TYPE_DEFAULT = "student"
	INVALID_FEATURE   = "Invalid feature type ..."
	PDFPATH           = "pdf-printer-output"
)

// App reusable vars ...
var (
	STUDENT_COLUMNS = []string{"School_Name", "Principal_Name", "Student_Name", "Academic_Year", "Class", "English", "Hindi", "Maths", "Science", "Computer_Science", "Result"}
	EMPLOYEE_COLUMN = []string{"Company_Name", "Employee_Name", "Designation", "Pay_Period", "Date_Of_Payment", "CTC", "Net_Salary", "Work_Days", "Absence", "Basic", "Hra", "Skill_Allowance", "Medical_Allowance", "PF_Contribution", "Income_Tax", "Exemption", "Net_Payment"}
	CUSTOMER_COLUMN = []string{"Company_Name", "Company_Contact", "Order_ID", "Order_Date", "Customer_Name", "Product_Name", "Quantity", "Price", "Discount", "Total"}
)
