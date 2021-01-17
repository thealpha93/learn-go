package main

import (
	"fmt"
)

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empID    int
	basicPay int
	pf       int
}

type Contract struct {
	empID    int
	basicPay int
}

type Freelancer struct {
	empId       int
	ratePerHour int
	totalHours  int
}

// Salary of permanent employee is the sum of basic pay and pf
func (p Permanent) CalculateSalary() int {
	return p.basicPay + p.pf
}

// Salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {
	return c.basicPay
}

//salary of freelancer
func (f Freelancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours
}

// Total expense is calculated by iterating through the salaryCalculator slice
// and summing the salaries of individual employees
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}

	fmt.Printf("Total expense per month $%d", expense)
}

func main() {
	pemp1 := Permanent{
		empID:    1,
		basicPay: 5000,
		pf:       20,
	}
	pemp2 := Permanent{
		empID:    2,
		basicPay: 6000,
		pf:       30,
	}
	cemp1 := Contract{
		empID:    3,
		basicPay: 3000,
	}
	freelancer1 := Freelancer{
		empId:       4,
		ratePerHour: 70,
		totalHours:  120,
	}
	freelancer2 := Freelancer{
		empId:       5,
		ratePerHour: 100,
		totalHours:  100,
	}

	employees := []SalaryCalculator{pemp1, pemp2, cemp1, freelancer1, freelancer2}
	totalExpense(employees)
}
