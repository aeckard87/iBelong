package main

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

// Account struct
type Account struct {
	FirstName string
	LastName  string
}

// Purchase struct
type Purchase struct {
	Date          time.Time
	Description   string
	AmountInCents int
}

// Statement Struct
type Statement struct {
	FromDate  time.Time
	ToDate    time.Time
	Account   Account
	Purchases []Purchase
}

func main() {
	fmap := template.FuncMap{
		"formatAsDollars": formatAsDollars,
		"formatAsDate":    formatAsDate,
		"urgentNote":      urgentNote,
	}
	file := "email.tmpl"
	t := template.Must(template.New(file).Funcs(fmap).ParseFiles(file))
	err := t.Execute(os.Stdout, createMockStatement())
	if err != nil {
		panic(err)
	}
}

func formatAsDollars(valueInCents int) (string, error) {
	dollars := valueInCents / 100
	cents := valueInCents % 100
	return fmt.Sprintf("$%d.%2d", dollars, cents), nil
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%d/%d", day, month, year)
}

func urgentNote(acc Account) string {
	fmt.Println("Account:", acc.LastName)
	return fmt.Sprintf("You have earned 100 VIP points that can be used for purchases")
}

func createMockStatement() Statement {
	return Statement{
		FromDate: time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
		ToDate:   time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC),
		Account: Account{
			FirstName: "John",
			LastName:  "Dow",
		},
		Purchases: []Purchase{
			{
				Date:          time.Date(2016, 1, 3, 0, 0, 0, 0, time.UTC),
				Description:   "Shovel",
				AmountInCents: 2326,
			},
			{
				Date:          time.Date(2016, 1, 8, 0, 0, 0, 0, time.UTC),
				Description:   "Staple remover",
				AmountInCents: 5432,
			},
		},
	}
}
package main

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

// Account struct
type Account struct {
	FirstName string
	LastName  string
}

// Purchase struct
type Purchase struct {
	Date          time.Time
	Description   string
	AmountInCents int
}

// Statement Struct
type Statement struct {
	FromDate  time.Time
	ToDate    time.Time
	Account   Account
	Purchases []Purchase
}

func main() {
	fmap := template.FuncMap{
		"formatAsDollars": formatAsDollars,
		"formatAsDate":    formatAsDate,
		"urgentNote":      urgentNote,
	}
	file := "email.tmpl"
	t := template.Must(template.New(file).Funcs(fmap).ParseFiles(file))
	err := t.Execute(os.Stdout, createMockStatement())
	if err != nil {
		panic(err)
	}
}

func formatAsDollars(valueInCents int) (string, error) {
	dollars := valueInCents / 100
	cents := valueInCents % 100
	return fmt.Sprintf("$%d.%2d", dollars, cents), nil
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%d/%d", day, month, year)
}

func urgentNote(acc Account) string {
	fmt.Println("Account:", acc.LastName)
	return fmt.Sprintf("You have earned 100 VIP points that can be used for purchases")
}

func createMockStatement() Statement {
	return Statement{
		FromDate: time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
		ToDate:   time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC),
		Account: Account{
			FirstName: "John",
			LastName:  "Dow",
		},
		Purchases: []Purchase{
			{
				Date:          time.Date(2016, 1, 3, 0, 0, 0, 0, time.UTC),
				Description:   "Shovel",
				AmountInCents: 2326,
			},
			{
				Date:          time.Date(2016, 1, 8, 0, 0, 0, 0, time.UTC),
				Description:   "Staple remover",
				AmountInCents: 5432,
			},
		},
	}
}
