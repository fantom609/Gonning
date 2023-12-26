package main

import (
	"database/sql"
	"fmt"
	"log"
	"main/src/Event"
	"main/src/color"
	"main/src/database"
	"main/src/input"
	"os"
	"os/exec"
	"runtime"
)

var (
	db *sql.DB
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
	var choice int
	for {
		for {
			displayMenu()
			choice, err = input.InputInt()
			if err != nil || choice != 1 && choice != 2 && choice != 3 && choice != 4 && choice != 5 && choice != 6 {
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
	fmt.Println(color.Cyan + " 6." + color.Reset + "  Quitter")
	fmt.Println()
	fmt.Println("entrer votre choix :")
}

func switchMenu(choice int) {

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
		fmt.Println("Visualiser les événements")
		events, err := database.GetEvents(db)
		if err != nil {
			log.Printf("%v", err)
		}

		for i := 0; i < len(events); i++ {
			fmt.Printf("ID: %d\n", events[i].Id)
			fmt.Printf("Title: %s\n", events[i].Title)
			fmt.Printf("StartDate: %s\n", events[i].StartDate)
			fmt.Printf("EndDate: %s\n", events[i].EndDate)
			fmt.Printf("Location: %s\n", events[i].Location)
			fmt.Printf("Tag: %s\n", events[i].Tag)
			fmt.Printf("Description: %s\n", events[i].Description)

		}

		input.InputString()
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
		fmt.Println("Aurevoir !")
		os.Exit(1)
	default:

	}
}

//func addEvent() Event.Event {}

//func writeEvent() Event {}

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
