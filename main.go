package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"main/src/Event"
	"main/src/color"
	"main/src/database"
	"main/src/input"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"time"
)

var (
	db        *sql.DB
	eventsMap = make(map[int]Event.Event)
	userId    int
)

func main() {
	loop := "yes"
	var err error
	for loop == "yes" {
		db, err = database.ConnectionDatabase()

		if err != nil {
			log.Printf(color.Red+"%v"+color.Reset, err)
			fmt.Printf("Voulez-vous tenter de vous reconnecter : (yes/no)\n")
			loop = input.InputString()
		} else {
			break
		}
	}
	for {
		userId = userConnection()
		if userId != 0 {
			break
		}
	}
	loop = "yes"
	for loop == "yes" {
		err = database.GetEvents(db, eventsMap, userId)
		if err != nil {
			log.Printf(color.Red+"%v"+color.Reset, err)
			fmt.Printf("Voulez-vous tenter de re récupérer les données ? : (yes/no)\n")
			loop = input.InputString()
		} else {
			break
		}
	}
	fmt.Printf("Vous avez %d évenement aujoursd'hui voulez vous les voirs ? (yes/no)", len(upcomingEvents()))
	for {
		choice := input.InputString()
		if choice == "yes" {
			events := upcomingEvents()
			displayEvents(events)
			fmt.Printf("Acces au menu : 1")
			fmt.Printf("quitter : 2")
			break
		}
		if choice == "no" {
			break
		}
		log.Printf(color.Red + "La valeur saisie est incorrecte" + color.Reset)
	}

	var choice int
	for {
		for {
			displayMenu()
			choice, err = input.InputInt()
			if err != nil || choice != 1 && choice != 2 && choice != 3 && choice != 4 && choice != 5 && choice != 6 && choice != 7 {
				log.Printf(color.Red + "La valeur saisie est incorrecte" + color.Reset)
			} else {
				break
			}
		}
		switchMenu(choice)
	}
}

func displayMenu() {
	clearScreen()
	fmt.Println()
	fmt.Println(color.Blue + " Bienvenue dans le Système de gestion de plannings")
	fmt.Println("--------------------------------------------------" + color.Reset)
	fmt.Println(color.Cyan + " 1." + color.Reset + "  Créer un nouvel événement")
	fmt.Println(color.Cyan + " 2." + color.Reset + "  Visualiser les événements")
	fmt.Println(color.Cyan + " 3." + color.Reset + "  Modifier un événement")
	fmt.Println(color.Cyan + " 4." + color.Reset + "  Supprimer un événement")
	fmt.Println(color.Cyan + " 5." + color.Reset + "  Rechercher un événement")
	fmt.Println(color.Cyan + " 6." + color.Reset + "  Voir les rappels")
	fmt.Println(color.Cyan + " 7." + color.Reset + "  Quitter")
	fmt.Println()
	fmt.Println("entrer votre choix :")
}

func switchMenu(choice int) {
	var exitRequested bool
	switch choice {
	case 1:
		for {
			clearScreen()
			event, err := Event.CreateEvent()
			if err != nil {
				log.Printf("%v", err)
			}
			var id int
			id, err = database.CreateEvent(event, db)
			if err != nil {
				log.Printf("%v", err)
			}
			event.Id = id
			var res string
			for res != "yes" && res != "no" {
				fmt.Println("Voulez vous saisir un autre evennement ? (yes/no)")
				res = input.InputString()
				if res == "yes" {
					continue
				}
				if res == "no" {
					break
				}
			}
		}
	case 2:

		var id int
		var err error

		clearScreen()
		fmt.Println(color.Blue + "\nListe du planning :" + color.Reset)

		displayEvents(eventsMap)

		fmt.Println("Entrez le numéro de l'événement pour voir plus de détails ou 0 pour revenir :")
		for {
			for {
				id, err = input.InputInt()
				if err == nil {
					break
				}
				fmt.Println("Saisi invalide")
			}
			if id == 0 {
				return
			}
			err = displayEvent(id)
			if err == nil {
				break
			}
			fmt.Println("l'id saisi n'existe pas")
		}

		fmt.Println("\nmodifier l'évenement : 1")
		fmt.Println("Revenir au menu : 2")
		fmt.Println("quitter : 3")

		var res int
		for {
			res, err = input.InputInt()
			if err == nil && res != 1 && res != 2 && res != 3 {
				continue
			}
			break
		}
		switch res {
		case 1:
			// updateEvent
			break
		case 2:
			return
		case 3:
			exitRequested = true
		}
		break
	case 3:
		fmt.Println("Modifier un événement")

		break
	case 4:
		fmt.Println("Supprimer un événement")

		break
	case 5:
		fmt.Println("Rechercher un événement")
		break
	case 6:
		exitRequested = true
	default:

	}
	if exitRequested {
		db.Close()
		fmt.Println("Aurevoir !")
		os.Exit(1)
	}
}

func clearScreen() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux", "darwin":
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		fmt.Println("Unsupported operating system")
		return
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func displayEvents(Events map[int]Event.Event) {

	type kv struct {
		Key   int
		Value Event.Event
	}

	var ss []kv
	for k, v := range Events {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[j].Value.StartDate.After(ss[i].Value.StartDate)
	})

	for _, event := range ss {
		fmt.Printf(color.Cyan+" %d."+color.Reset+" %s - %s - %s\n", event.Key, event.Value.Title, event.Value.StartDate.Format("2006-01-02 15:04"), event.Value.Tag)
	}

}

func displayEvent(id int) error {

	event, existe := eventsMap[id]
	if !existe {
		return errors.New("clé incorrecte")
	}

	clearScreen()
	fmt.Printf(color.Blue+"%s\n"+color.Reset, event.Title)
	fmt.Printf("Débute a    : %s\n", event.StartDate.Format("2006-01-02 15:04"))
	fmt.Printf("Termine a   : %s\n", event.EndDate.Format("2006-01-02 15:04"))
	fmt.Printf("durée       : %s\n", event.EndDate.Sub(event.StartDate))
	fmt.Printf("Tag         : %s\n", event.Tag)
	fmt.Printf("description : %s\n", event.Description)

	return nil
}

func userConnection() int {

	fmt.Println("saisisez votre identifiant")
	username := input.InputString()

	fmt.Println("saisisez votre mot de passe")
	password := input.InputString()

	id, err := database.ConnectUser(db, username, password)
	if err != nil {
		log.Printf(color.Red+"%v"+color.Reset, err)
	}
	return id
}

func upcomingEvents() map[int]Event.Event {
	events := make(map[int]Event.Event)
	for id, event := range eventsMap {
		if event.StartDate.After(time.Now()) && event.StartDate.Before(time.Now().Add(24*time.Hour)) {
			events[id] = event
		}
	}
	return events
}
