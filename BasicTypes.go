package main

import "fmt"


//Type declaration
type Celsious float64

func main() {
	var i int
	var j float64
	fmt.Print(i, j)

	//create a pointer
	p := &i
	*p = 10
	println(i,*p,p)

	//new function
	pNew := new(int) // creates an unnamed variable of type T and init it to 0 and returns its address (*T)
	println(pNew, *pNew)

	//lifetime of a variable
	// interval of time during which it exist as program executes

	//t exist only in this block
	for t:=0; t < 10; t++ {
		println(t)
	}

	dishes := [] string { "paneer", "lentils","cauliflower"}

	for t2:=0; t2 < len(dishes); t2++ {
		print(dishes[t2] + " ")
	}

	var temp1 Celsious
	println(temp1)

}
