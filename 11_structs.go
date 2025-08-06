package main

import "fmt"

type Person struct {
	name string
	age  uint
	//clothing Clothing
}

/*
type Clothing struct {
	shoeSize   uint
	shirtColor string
}
*/
/*
	func (p Person) GetName() string {
		return p.name
	}

	func (p Person) SetName(newName string) {
		p.name = newName
		fmt.Println(p)
	}
*/

/*
	func (c Clothing) GetShoeSize() uint {
		return c.shoeSize
	}

	func (p Person) GetShoeSize() uint {
		return p.clothing.GetShoeSize()
	}

/*
*/
func (p Person) Equal(p2 Person) bool {
	return p.age == p2.age && p.name == p2.name
}
func main() {
	slice := []Person{{"Alice", 20}, {"Alice", 20}}
	fmt.Println(slice[0].Equal(slice[1]))
	//p1 := Person{"Tim", 21}
	//var p1 Person
	//p1.name = "JTim"
	//p1.age = 18
	//fmt.Println(p1)

	/*p1 := Person{"Tim", 21, Clothing{12, "Blue"}}
	shoesize := p1.clothing.GetShoeSize()

	fmt.Println(p1)
	fmt.Println(p1.clothing.shirtColor)
	fmt.Println(shoesize)*/
	/*p1.SetName("Joeu")1
	fmt.Println(p1.GetName())*/
}
