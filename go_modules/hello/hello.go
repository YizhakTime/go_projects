package main
import (
	"fmt"
	"gm1/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Yizhak", "Jorge", "Thomas"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
