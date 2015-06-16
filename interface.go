
package main
import "fmt"

type Named interface {
	Name() string
}

func greet(someone Named) { fmt.Println("Greetings, " + someone.Name()) }

type Human struct {
	firstName string
}

func (h Human) Name() string {
	return h.firstName
}

func main() {
	greet(Human{firstName: "Jack Bauer"})
}