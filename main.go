package main // Declaring package main for specifying that the file is executable

import "fmt"

type Person struct { // Declaring "Person" structure to store attributes such as name and age of the person
	Name string
	Age  int
}

func (p Person) introduction() { // This method gives the introduction of the person
	fmt.Printf("My name is %v and my age is %v\n", p.Name, p.Age)
}

func (p Person) update(NewAge int) { // This method gives the name and updates the age of the person
	p.Age = NewAge
	fmt.Printf("My name is %v and my updated age is %v\n", p.Name, p.Age)
}

func (p Person) check(NewAge int) { // Checks the condition for the age
	if NewAge > 18 {
		fmt.Println("Eligible")
	} else {
		fmt.Println("Not Eligible")
	}
}

func main() { // Declaring function main
	p1 := Person{Name: "Pranjali", Age: 16}
	p1.introduction()
	p1.update(21)
	p1.check(21)

} // End of main function
