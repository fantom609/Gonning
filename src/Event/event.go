package Event

import (
	"fmt"
	"main/src/input"
	"regexp"
	"time"
)

type Event struct {
	Title       string
	StartDate   time.Time
	EndDate     time.Time
	Location    string
	Tag         string
	Description string
}

func CreateEvent() (*Event, error) {

	event := new(Event)
	var err error

	fmt.Print("Entrez le titre de l'événement: ")
	event.Title = input.InputString()

	for {
		fmt.Print("Entrez la date de début (YYYY-MM-DD hh:mm): ")
		startDateString := input.InputString()

		event.StartDate, err = valideDate(startDateString)
		if err == nil {
			break
		}
		fmt.Println("Le format de la date n'est pas valide")
	}
	for {
		fmt.Print("Entrez la date de fin (YYYY-MM-DD): ")
		EndDateString := input.InputString()

		event.EndDate, err = valideDate(EndDateString)
		if err == nil {
			break
		}
		fmt.Println("Le format de la date n'est pas valide")
	}

	fmt.Print("Entrez le lieu: ")
	event.Location = input.InputString()

	fmt.Print("Choisissez une catégorie (Professionnel, Personnel, Loisir): ")
	event.Tag = input.InputString()

	fmt.Print("Entrez une brève description: ")
	event.Description = input.InputString()

	return event, nil
}

func valideDate(dateString string) (time.Time, error) {
	regex := regexp.MustCompile(`\b\d{4}-\d{2}-\d{2} \d{2}:\d{2}\b`)
	if !regex.MatchString(dateString) {
		return time.Time{}, fmt.Errorf("le format n'est pas valide")
	}
	date, err := time.Parse("2006-01-02 15:04", dateString)
	if err != nil {
		return time.Time{}, fmt.Errorf("une erreur est survenue lors de la conversion de la date : %v", err)
	}
	return date, nil
}

//func writeEvent() Event {}
