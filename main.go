package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

type Event struct {
	title     string
	startDate string
	endDate   string
	time      string
	location  string
	tag       string
}

func main() {

}

func displayMenu() {
	fmt.Println(Blue + " Bienvenue dans le Système de gestion de plannings")
	fmt.Println("--------------------------------------------------" + Reset)
	fmt.Println(Cyan + " 1." + Reset + "  Créer un nouvel événement")
	fmt.Println(Cyan + " 2." + Reset + "  Visualiser les événements")
	fmt.Println(Cyan + " 3." + Reset + "  Modifier un événement")
	fmt.Println(Cyan + " 4." + Reset + "  Supprimer un événement")
	fmt.Println(Cyan + " 5." + Reset + "  Rechercher un événement")
	fmt.Println(Cyan + " 6." + Reset + "  Quitter")
	fmt.Println()
}

func switchMenu(choice int) {

	switch choice {
	case 1:
		fmt.Println("Visualiser les événements")
		break
	case 2:
		fmt.Println("Modifier un événement")
		break
	case 3:
		fmt.Println("Supprimer un événement")
		break
	case 4:
		fmt.Println("Rechercher un événement")
		break
	case 5:
		fmt.Println("Visualiser les événements")
		break
	case 6:
		fmt.Println("Quitter")
		break
	default:

	}
}

func inputInt() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strconv.Atoi(scanner.Text())
}

func inputString() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
