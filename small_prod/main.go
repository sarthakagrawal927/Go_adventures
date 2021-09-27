package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	sarthak := person{
		firstName: "sarthak",
		contactInfo: contactInfo{
			email: "1@gmail",
			zip:   4,
		},
	}
	sarthak.lastName = "agrawal"
	// sarthak.print()
	sarthak.updateName("Noop")
	// sarthak.print()

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#232233",
	}
	colors["white"] = "ffffff"
	// delete(colors, "green")
	printMap(colors)
	printMap(colors)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func printMap(m map[string]string) {
	for key, val := range m {
		fmt.Println("Hex for", key, "is", val)
	}
	m["white"] = "f"
}
