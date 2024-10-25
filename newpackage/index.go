package newpackage

import "fmt"

type Person struct{
	id int
	height float32
	name string
}

func (p *Person) SetID(id int){
	p.id = id
}

func (p *Person) SetHeight(height float32){
	p.height = height
}

func (p *Person) SetName(name string){
	p.name = name
}

func (p *Person) GetID() int {
	return p.id
}

func (p *Person) GetHeight() float32{
	return p.height
}

func (p *Person) GetName() string{
	return p.name
}

type Car struct{
	id int
}

func (c *Car) SetId(id int){
	c.id = id
}

func (c *Car) GetId() int {
	return c.id
}


func main(){

	var n int
	fmt.Println("Please provide the value for the variable n")
	fmt.Scan(&n)

	var v = make([]Person, n)
	for i:=0;i<n;i++{
		var idVal int
		var heightVal float32
		var nameVal string
		fmt.Println("Please provide the value of the person ", i+1)
		fmt.Scan(&idVal, &heightVal, &nameVal)

		personVal := Person{idVal, heightVal, nameVal} 

		v[i] = personVal
	}

	for i:=0;i<n;i++{
		fmt.Println("The value of the ", i+1, " index person is ", v[i].id, v[i].height, v[i].name)
	}

}