package main

import (
	"encoding/json"
	"fmt"
	"math"
)

// Define and Instantiate a Struct

// Define a struct called Book with fields Title, Author, and Pages.
// Create an instance of Book with some sample data.
// Print the details of the book.

type Book struct {
	Title string
	Author string
	Pages uint
}


func getBook() {
	book1 := Book {
		Title: "LOTR",
		Author: "JRRT",
		Pages: 1000,
	}
	fmt.Printf("The title of the book is %v \n", book1.Title)
	fmt.Printf("The author of the book is %v\n", book1.Author)
	fmt.Printf("There are %v pages in the book\n", book1.Pages)
}
// Create a struct Car with fields Brand, Model, and Year.
// Initialize a Car instance and modify its Year field to a newer value.
// Print the updated struct.

type Car struct {
	Brand string
	Model string
	Year uint
}

func getCar() {
	car := Car {
		Brand: "BMW",
		Model: "Mod",
		Year: 1995,
	}
	car.Year = 1996
	fmt.Printf("Car year is %v\n", car.Year)
	fmt.Printf("%v", car)
}

// Create a struct Person with fields Name and Age.
// Instantiate Person and use a pointer to modify its Age.
// Print the original struct to verify the changes.

type Person struct {
	Name string
	Age uint
}
func changeAge(p *Person) {
	p.Age = 30
}

func getPerson() {
	person := Person {
		Name: "Tom",
		Age: 32,
	}
	changeAge(&person)
	fmt.Printf("%v's age is %v\n", person.Name, person.Age)
}

// Define a struct Address with fields Street, City, and ZipCode.
// Define another struct Employee with fields Name, Position, and an Address field (type Address).
// Create an Employee instance and initialize its Address.

type Address struct {
	Street string
	City string
	ZipCode string
}

type Employee struct {
	Name string
	Position string
	Address Address	
}

func getEmployee() {
	emp := Employee {
		Name: "Tom",
		Position: "Manager",
		Address: Address{
			Street: "Baker St",
			City: "London",
			ZipCode: "65DXQ",
		},
	}
	fmt.Printf("Employee is %v with address %v \n", emp, emp.Address)
}
// Define a struct Rectangle with fields Width and Height.
// Add a method Area() to calculate and return the area of the rectangle.
// Instantiate the struct and call the Area() method.

type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func getArea() {
	rect := Rectangle {
		Width: 10.0,
		Height: 5,
	}
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())


}

// Create and use an anonymous struct with fields Name and Age.
// Initialize and print the struct in a single statement.

func createStruct() {
	fmt.Println(struct {
		Name string
		Age int
	} {
		Name: "Tom",
		Age: 30,
	})
}


// Define a struct Product with fields Name, Price, and InStock.
// Convert an instance of Product into JSON format.
// Parse the JSON back into a Product struct and print it.

type Product struct {
	Name string `json:"name"`
	Price uint `json:"price"`
	InStock bool `json:"in_stock"`
}

func parseProduct() {
	product := Product {
		Name: "thing",
		Price: 10,
		InStock: true,
	}
	jsonData, err := json.Marshal(product)
	if err != nil {
		fmt.Println("error encoding json", err)
		return
	}
	fmt.Println(string(jsonData))
	var productFromJson Product
	err2 := json.Unmarshal([]byte(jsonData), &productFromJson)
	if err2 != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println(productFromJson)


}

// Define a struct Student with fields Name and Grade.
// Create a slice of Student structs.
// Loop through the slice and print each student's details.
type Student struct {
	Name string
	Grade float32
}

func getStudents() {
	students := []Student{
		{Name: "John", Grade: 10.5},
		{Name: "Tom", Grade: 90.3},
	}

	for _, student := range students {
		fmt.Println(student.Name)
	}
}

// Define a struct User with fields Name, Email, and Age.
// Add JSON tags to the fields (e.g., json:"name") and serialize the struct to JSON.
// Print the resulting JSON to see the effect of the tags.

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Age uint `json:"age"`
}

func parseUser() {
	usr := User {
		Name: "Jack",
		Email: "jack@mail.com",
		Age: 20,
	}

	jsonData, err := json.Marshal(usr)
	if err != nil {
		fmt.Println("error encoding json", err)
		return
	}
	fmt.Println(string(jsonData))
}
// Define a struct Contact with fields Phone and Email.
// Define another struct Profile that embeds Contact and has an additional Name field.
// Create an instance of Profile and access both Name and Phone fields directly.

type Contact struct {
	Phone string
	Email string
}

type Profile struct {
	Contact
	Name string
}

func getProfile() {
	profile := Profile {
		Contact: Contact {
			Phone: "901501",
			Email: "some@mail.ru",
		},
		Name: "Tom",
	}
	fmt.Println(profile)
	fmt.Print(profile.Phone)
}

// Define a struct Circle with a single field Radius.
// Add methods Area() and Circumference() to calculate the area and circumference of the circle.
// Instantiate the struct and call these methods.

type Circle struct {
	Radius float64
}

func (c Circle) getArea() float64 {
	p := 3.14
	return p * math.Pow(c.Radius, 2)
}

func (c Circle) getCircumference() float64 {
	p := 3.14
	return p*(c.Radius*2)
}

func getCircleAreaAndCirc() {
	circle := Circle {
		Radius: 4,
	}
	fmt.Printf("Area: %.2f\n", circle.getArea())
	fmt.Printf("Circ: %.2f\n", circle.getCircumference())

}

// Define a struct Animal with fields Name and Age.
// Define another struct Dog that embeds Animal and adds a Breed field.
// Instantiate a Dog and access its fields, including the embedded Animal fields.

type Animal struct {
	Name string
	Age uint
}

type Dog struct {
	Breed string
	Animal
}

func getAnimal() {
	dog := Dog {
		Breed: "shepherd dog",
		Animal: Animal{
			Name: "Reks",
			Age: 3,
		},
	}
	fmt.Printf("Dog's breed is %v and dog's name is %v\n", dog.Breed, dog.Name)
}

// Define a struct Product with fields Name, Price, and InStock.
//     Create a constructor function NewProduct(name string, price float64) Product that initializes a Product with InStock set to true by default.
//     Use this constructor to create and print a Product.

type Product2 struct {
	Name string
	Price float64
	InStock bool
}

func NewProduct(name string, price float64) Product2 {
	return Product2{
		Name: name,
		Price: price,
		InStock: true,
	}
}

func getProduct2() {
	p := NewProduct("Phone", 120.00)
	fmt.Printf("Product name is %v, price is %v and stock is %v\n", p.Name, p.Price, p.InStock)
}

// Define a struct Author with fields Name and Country.
// Define another struct Book with fields Title, Pages, and an Author field of type Author.
// Add a method Description() to Book that returns a string containing all its details.
// Instantiate a Book and call the Description() method.

type Author struct {
	Name string
	Country string
}

type Book2 struct {
	Title string
	Pages uint
	Author
}

func (b Book2) Description() string {
	return fmt.Sprintf("Book title is %v, %v pages, author's name is %v, country is %v", b.Title, b.Pages, b.Name, b.Country)
}

func getBook2() {
	book := Book2{
		Title: "LOTR",
		Pages: 1000,
		Author: Author {
			Name: "JRRT",
			Country: "UK",
		},
	}
	descr := book.Description()
	fmt.Println(descr)
}

// Define a struct User with fields Name, Email, and Age.
// Add JSON tags to the fields.
// Serialize a User instance to JSON and print the result.
// Deserialize the JSON back into a User struct and print its fields.



func runStructs() {
	fmt.Println("hey")
	getBook()
	getCar()
	getPerson()
	getEmployee()
	getArea()
	createStruct()
	parseProduct()
	getStudents()
	parseUser()
	getProfile()
	getCircleAreaAndCirc()
	getAnimal()
	getProduct2()
	getBook2()
}
