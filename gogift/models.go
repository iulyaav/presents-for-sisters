package main

type EventType string

const (
	Christmas EventType = "Christmas"
	Birthday EventType = "Birthday"
 )

type Present struct {
	Name string `json:"name"`
	// Description string
	// Link string
	// Event EventType
}

type Persons struct {
    Persons []Person `json:"persons"`
}

type Person struct {
	Name string `json:"name"`
	Presents []Present `json:"present"`
}

func (p *Person) addPresent(present Present) {
	p.Presents = append(p.Presents, present);
}

func newPerson(name string) Person {
	return Person{Name: name, Presents: []Present{}}
}