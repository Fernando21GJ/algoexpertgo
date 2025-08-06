package main

import "fmt"

type Name interface {
	GetName() string
	//GetLastName() string
}

type Person1 struct {
	firstName string
	lastName  string
}

func (p Person1) GetName() string {
	return p.firstName + " " + p.lastName
}
func (p Person1) SayHi() {
	fmt.Println("Hi")
}

type Employee struct {
	name string
}

func (e Employee) GetName() string {
	return e.name
}

func PrintName(obj Name) {
	fmt.Println(obj.GetName())
}

func main() {
	var name Name = Person1{"Alice", "Alice"}
	names := []Name{Employee{"Alice"}, Person1{"Fer", "Bob"}}
	for _, name := range names {
		fmt.Println(name.GetName())
	}
	fmt.Println(name)
	fmt.Println(names)
}
