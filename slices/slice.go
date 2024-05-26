package main
import "fmt"

func main() {
	x := [3]string{"Лайка", "Белка", "Стрелка"}
	s := x[:] // a slice referencing the storage of x
	t := x[1:] //including left element
	a := x[:2] //excluding right element
	var ex []byte
	ex = make([]byte, 5)
	//capacity if omitted is defaulted to length i.e. 5 here
	fmt.Println(s)
	fmt.Println(t)
	fmt.Println(a)
	d := []byte{'r','o','a','d'}
//	e := d[2:]
	fmt.Printf("%c\n", d)
	fmt.Println(ex)
}
