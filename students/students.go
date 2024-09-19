package students

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s Student) Printdata() {
	fmt.Printf("Name: %s, Age: %d", s.Name, s.Age)
}
