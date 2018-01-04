package main

import "fmt"

type Humen struct {
	name string
	age int
}

type Student struct {
	Humen
	school string
}

type Employee struct {
	Humen
	comany string
}

func (h Humen) SayHello()  {
	fmt.Println(h.name+" Say Hello")
}

func (h Humen) Sing(str string) {
	fmt.Println("La La La... ",str)
}



func (stu Student) SayHello()  {
	stu.name="zcc_new"
	fmt.Println(stu.Humen.name+" Say Hello, then he is in "+stu.school)
}

type Men interface {
	SayHello()
	Sing(str string)
}

func main() {
	stu:=Student{Humen{name:"zcc",age:13},"chengzhang experimental middle school"}
	me:=Employee{Humen{name:"akey",age:42},"KH company"}

	/*var h Men
	h = stu
	h.SayHello()
	h.Sing("student")
	h = me
	h.SayHello()
	h.Sing("employee")*/
	h := make([]Men,2)
	h[0], h[1] = me, stu

	for _,v := range h {
		v.SayHello()
	}
}
