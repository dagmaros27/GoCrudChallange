package model

import "github.com/google/uuid"

// Person represents a person object in the system.
type Person struct {
	ID      string   `json:"id"`
	Name    string   `json:"name" binding:"required"`
	Age     int      `json:"age" binding:"required"`
	Hobbies []string `json:"hobbies" binding:"required"`
}

// In-memory storage for persons
var People = make(map[string]Person)

func CreatePerson(p Person) Person {
	p.ID = uuid.New().String()
	People[p.ID] = p
	return p
}

func GetAllPersons() []Person {
	personList := make([]Person, 0, len(People))
	for _, person := range People {
		personList = append(personList, person)
	}
	return personList
}

func GetPersonByID(id string) (Person, bool) {
	person, exists := People[id]
	return person, exists
}

func UpdatePerson(id string, p Person) (Person, bool) {
	_, exists := People[id]
	if !exists {
		return Person{}, false
	}
	p.ID = id
	People[id] = p
	return p, true
}

func DeletePerson(id string) bool {
	_, exists := People[id]
	if exists {
		delete(People, id)
	}
	return exists
}
