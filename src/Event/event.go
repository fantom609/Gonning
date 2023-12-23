package Event

import (
	"fmt"
	"main/src/input"
)

type Event struct {
	title       string
	startDate   string
	time        string
	location    string
	tag         string
	description string
}

func AddEvent() *Event {

	event := new(Event)

	fmt.Print("Entrez le titre de l'événement: ")
	event.title = input.InputString()
	fmt.Print("Entrez la date (YYYY-MM-DD): ")
	event.startDate = input.InputString()
	fmt.Print("Entrez l'heure (HH:MM): ")
	event.time = input.InputString()
	fmt.Print("Entrez le lieu: ")
	event.location = input.InputString()
	fmt.Print("Choisissez une catégorie (Professionnel, Personnel, Loisir): ")
	event.tag = input.InputString()
	fmt.Print("Entrez une brève description: ")
	event.description = input.InputString()

	return event
}

//func writeEvent() Event {}
