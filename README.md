# go-personal-diary
---
A CLI that generates PDFs in bulk very quickly. Thanks to Golang concurrency.
Currently it supports following use cases:-

* Employee Payslip
* Customer Invoice
* Student Report Card

## Installation & setup :-
---
* Clone the repository.
* Run the command `make download`.
* Run the command `make install`.
* CLI setup is ready. Run the command `pdf-printer --help` to explore various operations.
* Sample generated [PDFs](https://github.com/Harry-027/pdf-printer/tree/master/samplePdfs).

## Feed the data
---
* Feed the csv data as per the feature type - student, employee, customer.
* Data should be fed in pre-defined csv format.Checkout [CSV samples](https://github.com/Harry-027/pdf-printer/tree/master/assets/csv) for required fields & format.
* Command to feed data `pdf-printer feed --type=customer|student|employee --path=path_to_csv_file`.
* Default value for flag 'type' -> 'student' and 'path' -> './assets/csv/students.csv'.

## Verify the data
---
* Data has been fed correctly or not can be verified by visualising it over command line interface in table format.
* Run the command `pdf-printer verify --type=customer|employee|student` to verify the data as per the given flag type.
* Default value for flag 'type' -> 'student'.

![pdf-printer verify](https://github.com/Harry-027/pdf-printer/blob/master/screenshots/ui_table.PNG "pdf-printer verify")

## Generate the Pdfs in bulk
---
* Run the command `pdf-printer gen --type=customer|employee|student` to generate Pdfs.
* Default value for flag 'type' -> 'student'.

![pdf-printer gen](https://github.com/Harry-027/pdf-printer/blob/master/screenshots/pdfs.PNG "pdf-printer gen")

