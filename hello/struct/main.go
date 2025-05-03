package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Address struct {
	Street  string
	City    string
	Country string
}

type Contact struct {
	Email string
	Phone string
}

type Employee struct {
	Person   // Embedded struct
	Address  //Embedded struct
	Contact  //Embedded struct
	Position string
}

func main() {
	var p Person
	p.FirstName = "kartik"
	p.Age = 22
	fmt.Println("Name:", p.FirstName)
	fmt.Println("Age:", p.Age)

	p2 := Person{
		FirstName: "dhiraj",
		Age:       23,
	}
	fmt.Println("Name:", p2.FirstName)
	fmt.Println("Age:", p2.Age)

	p3 := new(Person)
	p3.FirstName = "amol"
	p3.Age = 25
	fmt.Println("Name:", p3.FirstName)
	fmt.Println("Age:", p3.Age)

	// Creating an instance of the Employee struct
	employee := Employee{
		Person: Person{
			FirstName: "Frank",
			LastName:  "Miller",
			Age:       45,
		},
		Address: Address{
			Street:  "223 Main st",
			City:    "Anytown",
			Country: "Usa",
		},
		Contact: Contact{
			Email: "frank@example.com",
			Phone: "55-44-33-22",
		},
		Position: "manager",
	}

	fmt.Println("Employee Name:", employee.FirstName)
	fmt.Println("Employee Age:", employee.Age)
	fmt.Println("Employee Address:", employee.Address.Street, employee.Address.City, employee.Address.Country)
	fmt.Println("Employee Contact:", employee.Contact.Email, employee.Contact.Phone)
}
