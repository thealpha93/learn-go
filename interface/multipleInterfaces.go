package main

import "fmt"

type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	calculateLeavesLeft() int
}

// Embed interfaces
type EmployeeOperations interface {
	SalaryCalculator
	LeaveCalculator
}

type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) calculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func main() {
	e := Employee{
		firstName:   "Aakash",
		lastName:    "TM",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}

	// Zero value of an interface is nil

	var empOp EmployeeOperations = e
	var s SalaryCalculator = e
	var l LeaveCalculator = e

	empOp.DisplaySalary()
	s.DisplaySalary()
	fmt.Println("\nLeaves left = ", l.calculateLeavesLeft())
}
