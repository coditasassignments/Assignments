package Questions

import "fmt"

type Employee struct {
	Name   string
	Age    int
	Salary float64
}

type Department struct {
	Department_name string
	Employees       []Employee
}

func (d Department) calculate_avg() float64 {
	var sum_of_salary float64
	var average float64
	for _, i := range d.Employees {
		sum_of_salary += i.Salary
	}
	average = sum_of_salary / float64(len(d.Employees))
	return average

}
func Main2() {
	e1 := Employee{Name: "Pranjali", Salary: 2000.0}
	e2 := Employee{Name: "Bajpai", Salary: 4000.0}
	e3 := Employee{Name: "Priya", Salary: 6000.0}

	dept := Department{
		Department_name: "Pharmacy",
		Employees:       []Employee{e1, e2, e3},
	}

	averageSalary := dept.calculate_avg()

	// Printing results
	fmt.Println("Department:", dept.Department_name)
	fmt.Println("Average Salary of Employees:", averageSalary)

}

//Hello
