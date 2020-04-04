package main

import . "fmt"

// import "os"

// import "time"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person
	id   int
	addr string
}

func (tmp *Person) PrintInfo() {
	Printf("name is %s,sex is %c,age is %d", tmp.name, tmp.sex, tmp.age)
}

func (tmp *Student) PrintInfo() {
	Println("student tmp=", tmp)
}

func main() {
	s := Student{Person{"lzk", 'n', 30}, 1, "beijing"}
	s.PrintInfo()

	s.Person.PrintInfo()
}
