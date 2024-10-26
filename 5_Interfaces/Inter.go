package main
import "fmt"
import "math"

// Interfaces in Go
// Interfaces are implemented implicitly.
type shape interface {
	area() float64
	perimeter() float64
}

type rectan struct {
	width, height float64
}

func (r rectan) area() float64 {
	return r.width * r.height
}
func (r rectan) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Multiple Interface
type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}

// ---------MI-----------------
func (e email) cost() float64 {
	if !e.isSubscribed {
		return 0.05 * float64(len(e.body))
	}
	return 0.01 * float64(len(e.body))
}

func (e email) format() string {
	return fmt.Sprintf(e.body)
}

//---------------------------------------------------------------------

//Type Assertions in GO

/* -> we want to check if s is a circle in order to cast it into its underlying concrete type
   -> we know s is an instance of the shape interface, but we do not know if it's also a circle
   -> c is a new circle struct cast from s
   -> ok is true if s is indeed a circle, or false if s is NOT a circle*/

// type shape interface {
// 	area() float64
// }

// type circle struct {
// 	radius float64
// }

// c, ok := s.(circle)
// if !ok {
// 	// log an error if s isn't a circle
// 	log.Fatal("s is not a circle")
// }

// radius := c.radius


//Type Switch
func printNumericValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
}

func main(){
	//Interfaces
	// r := rectan{
	// 	width:  5,
	// 	height: 10,
	// }
	// fmt.Println(r.area(), r.perimeter())

	// c := circle{
	// 	radius: 20,
	// }
	// fmt.Println(c.area(), c.perimeter())


	//Type Switch
	printNumericValue(1)
	// prints "int"

	printNumericValue("1")
	// prints "string"

	printNumericValue(struct{}{})
	// prints "struct {}"
}