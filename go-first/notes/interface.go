package notes

import "fmt"

type Stringify interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("[%s] %s", c.Number, c.Model)
}

func InterfaceTest() {
	person := &Person{Name: "Taro", Age: 21}
	fmt.Println(person.ToString())
	car := &Car{Number: "08-10", Model: "AX-32"}
	fmt.Println(car.ToString())
}
