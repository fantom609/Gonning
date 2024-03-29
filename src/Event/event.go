package Event

import (
	"fmt"
	"main/src/input"
	"regexp"
	"time"
)

type Event struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	StartDate   time.Time `json:"start-date"`
	EndDate     time.Time `json:"end-date"`
	Location    string    `json:"localisation"`
	Tag         string    `json:"tag"`
	Description string    `json:"description"`
}

func CreateEvent() (*Event, error) {

	event := new(Event)
	var err error

	fmt.Print("Entrez le titre de l'événement: ")
	event.Title = input.InputString()
	for {
		for {
			fmt.Print("Entrez la date de début (YYYY-MM-DD hh:mm): ")
			startDateString := input.InputString()

			event.StartDate, err = valideDate(startDateString)
			if err == nil {
				break
			}
			fmt.Println("Le format de la date n'est pas valide")
			if event.StartDate.Before(time.Now()) {
				fmt.Println("La date début ne peut pas être dans le passé")
				continue
			}
		}
		for {
			fmt.Print("Entrez la date de fin (YYYY-MM-DD hh:mm): ")
			EndDateString := input.InputString()

			event.EndDate, err = valideDate(EndDateString)
			if err == nil {
				break
			}
			fmt.Println("Le format de la date n'est pas valide")
		}
		if event.StartDate.After(event.EndDate) {
			fmt.Println("La date début ne peut pas être après la date de fin")
			continue
		}
		break
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

func UpdateChoices(Event *Event, choices int) {

	var err error

	for choices%10 != 0 {

		choice := choices % 10
		switch choice {
		case 1:
			fmt.Print("Nouveau titre :")
			Event.Title = input.InputString()
			break

		case 2:

			for {
				fmt.Print("Nouvelle date de debut : ")
				StartDateString := input.InputString()
				Event.StartDate, err = valideDate(StartDateString)
				if err == nil {
					break
				}
				fmt.Println("Le format de la date n'est pas valide")
			}
			break

		case 3:
			for {
				fmt.Print("Nouvelle date de fin :")
				EndDateString := input.InputString()
				Event.EndDate, err = valideDate(EndDateString)
				if err == nil {
					break
				}
				fmt.Println("Le format de la date n'est pas valide")
			}
			break

		case 4:
			fmt.Print("Nouvelle localisation :")
			Event.Location = input.InputString()
			break

		case 5:
			fmt.Print("Nouveau Tag :")
			Event.Tag = input.InputString()
			break

		case 6:
			fmt.Print("Nouvelle description :")
			Event.Description = input.InputString()
			break

		default:
			fmt.Errorf("erreur de saisie, le champ n %d n'existe pas", choice)
			break
		}

		choices /= 10
	}
}
