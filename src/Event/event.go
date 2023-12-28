package Event

import (
	"database/sql"
	"fmt"
	"main/src/color"
	"main/src/input"
	"regexp"
	"time"
)

type Event struct {
	Id          int
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

func confirmEvent(event *Event) bool {

	fmt.Printf("\nTitre de l'événement:"+color.Blue+" %s"+color.Reset, event.Title)
	fmt.Printf("\nDate de début       :"+color.Blue+" %s"+color.Reset, event.StartDate.Format("2006-01-02 15:04"))
	fmt.Printf("\nDate de fin         :"+color.Blue+" %s"+color.Reset, event.EndDate.Format("2006-01-02 15:04"))
	fmt.Printf("\nLieu                :"+color.Blue+" %s"+color.Reset, event.Location)
	fmt.Printf("\nCatégorie           :"+color.Blue+" %s"+color.Reset, event.Tag)
	fmt.Printf("\nDescription         :"+color.Blue+" %s\n"+color.Reset, event.Description)

	for {
		fmt.Println("\ninformation correct ? (yes/no)")
		res := input.InputString()
		if res == "yes" {
			return true
		}
		if res == "no" {
			return false
		}
		fmt.Println("valeur incorrecte")
	}
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

func GetEvents(db *sql.DB) []Event {
	return []Event{}
}
