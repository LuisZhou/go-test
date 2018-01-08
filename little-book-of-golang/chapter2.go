package main

import (
	"fmt"
)

type Person struct {
	Name string
}

type Saiyan struct {
	Name  string
	Power int
	// add other type of field: arrays, maps, interfaces and functions
	Father *Saiyan
}

type SaiyanComposePersion struct {
	// Name  string // no useful, for introduce use Person struct.
	Power int
	*Person
	// Because we didn't give it an explicit field name, we can implicitly access
	// the fields and functions of the composed type. However, the Go compiler did
	// give it a field name
}

func main() {
	// __Section: Declarations and Initializations
	goku := Saiyan{
		Name:  "Goku",
		Power: 9000, // tailing
	}

	// empty init or only init some of the field.
	// goku := Saiyan{}

	// reply on the order of field.
	// goku := Saiyan{"Goku", 9000}

	Super(goku)
	fmt.Println(goku.Power) // 9000

	// using address of operation.
	// goku := &Saiyan{"Goku", 9000}
	SuperUsingPointer(&goku)
	fmt.Println(goku.Power) // 19000

	// __Section: Functions on Structures
	goku.Super()
	fmt.Println(goku.Power) // 29000

	// __Section: Constructors (factory pattern)
	// Structures don't have constructors. Instead, you create a function that returns an instance of the desired type (like a factory)
	goku_using_factory := NewSaiyan("using_factory", 1000)
	fmt.Println(goku_using_factory.Power) // 1000

	// __Section: New
	// Despite the lack of constructors, Go does have a built-in new function which is used to allocate the memory required by a type. The result of new(X) is the same as &X{}
	goku_using_new := new(Saiyan)
	goku_using_new.Name = "using_new"
	goku_using_new.Power = 9001
	// same as
	// goku := &Saiyan{}
	fmt.Println(goku_using_new.Power) // 9001

	// __Section: Fields of a Structure
	// gohan_with_more_field := &Saiyan{
	// 	Name:  "Gohan",
	// 	Power: 1000,
	// 	Father: &Saiyan{
	// 		Name:   "Goku",
	// 		Power:  9001,
	// 		Father: nil,
	// 	},
	// }

	// __Section: Composition
	// Go supports composition, which is the act of including one structure into another. In some languages, this is called a trait or a mixin.
	goku_with_compose := &SaiyanComposePersion{
		Person: &Person{"Goku"},
		Power:  9001,
		//Name:   "Nick name",
	}
	goku_with_compose.Introduce()

	// the following output is the same, if goku_with_compose has no field of Name.
	fmt.Println(goku_with_compose.Name)
	fmt.Println(goku_with_compose.Person.Name)

	goku_with_compose.Person.Introduce()

	// inheritance vs compose
	// Is composition better than inheritance? Many people think that it's a more robust way to share code.
	// When using inheritance, your class is tightly coupled to your superclass and you end up focusing on
	// hierarchy rather than behavior.

	// __Section: Overloading
	// example: in go: load, loadById, loadByName. in other language with overloading: load(), load(int id), load(std:string& name)
	// However, because implicit composition is really just a compiler trick, we can "overwrite" the functions of a composed type.
	goku_with_compose.Person.Introduce()

	// __Section: Pointers versus Values
	// a copy of large structures
	// thousands or possibly tens of thousands of copy.
}

func Super(s Saiyan) {
	s.Power += 10000
}

func SuperUsingPointer(s *Saiyan) {
	s.Power += 10000 // use . too.
	// s = &Saiyan{"Gohan", 1000} // if you do this, everything change to s is nothing to the value of before address which parater s point to.
}

// *Saiyan is the receiver of the Super method
func (s *Saiyan) Super() {
	s.Power += 10000
}

func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Name:  name,
		Power: power,
	}
}

// func NewSaiyan(name string, power int) Saiyan {
// 	return Saiyan{
// 		Name:  name,
// 		Power: power,
// 	}
// }

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

func (s *SaiyanComposePersion) Introduce() {
	fmt.Printf("Hi, I'm %s. Ya!\n", s.Name)
}
