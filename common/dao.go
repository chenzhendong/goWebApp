package common

import (
	"fmt"
	"os"
	"database/sql"
	_ "github.com/lib/pq"
)

type ExpenseReport struct {
	Id     int
	Title  string
	AuthorId int
	SubjectId int
}

func GetExpenseReportById(db *sql.DB, id int) (*ExpenseReport, error) {
	const query = `SELECT * FROM books`
	var retval ExpenseReport
	err := db.QueryRow(query).Scan(&retval.Id, &retval.Title, &retval.AuthorId, &retval.SubjectId)
	retval.Id = id
	return &retval, err
}

func main() {
	db, err := sql.Open("postgres", "user=postgres password='dyslmt' dbname=booktown sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	for i := 1; i != 3; i++ {
		expenseReport, err := GetExpenseReportById(db, i)
		if err != nil {
			fmt.Fprintf(os.Stdout, "Error:%s\n", err)
		} else {
			fmt.Fprintf(os.Stdout, "Expense Report:%v\n", expenseReport)
		}
	}
}
