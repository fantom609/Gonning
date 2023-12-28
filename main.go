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
	"sort"
	"strings"
	"time"
)

var (
	db            *sql.DB
	eventsMap     = make(map[int]Event.Event)
	userId        int
	exitRequested bool
)

func main() {
	exitRequested := false
	var err error
	loopConnexionDb() // connection a la base de données
	for {
		userId = userConnection()
		if userId != 0 {
			break
		}
	} // Connexion de l'utilisateur

	loopGetEvents() // récupération des événements de l'utilisateur connecté

	if len(upcomingEvents()) > 0 {
		fmt.Printf("Vous avez"+color.Orange+" %d "+color.Reset+"évenement aujoursd'hui voulez vous les voirs ? (yes/no)", len(upcomingEvents()))
		exitRequested = loopUpComingEvents()
	} // Affichage des événements jounaliés s'il y en a
	var choice int
	if exitRequested == true {
		choice = 6
	}
	for {
		for exitRequested == false {
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

// fonction liée au menu à et l'affichage
/*
Cette fonction permet d'afficher le menu
*/
func displayMenu() {
	clearScreen()
	fmt.Println()
	fmt.Println(color.Blue + " Bienvenue dans le Système de gestion de plannings")
	fmt.Println("--------------------------------------------------" + color.Reset)
	fmt.Println(color.Cyan + " 1." + color.Reset + "  Créer un nouvel événement")
	fmt.Println(color.Cyan + " 2." + color.Reset + "  Visualiser les événements")
	fmt.Println(color.Cyan + " 3." + color.Reset + "  Rechercher un événement")
	fmt.Println(color.Cyan + " 4." + color.Reset + "  Quitter")
	fmt.Println()
	fmt.Println("entrer votre choix :")
}

/*
Cette fonction permet d'effectuer les traitements souhaités par l'utilisateur
*/
func switchMenu(choice int) {
	switch choice {
	case 1:
		createEvent()
		break
	case 2:
		clearScreen()
		getEvents()
		break
	case 3:
		clearScreen()
		fmt.Println("Votre recherche ?")
		tag := input.InputString()
		events := searchEvent(tag)
		displayEvents(events)
		fmt.Println("Appuier sur une touche pour retourner au menu")
		input.InputString()
		break
	case 4:
		exitRequested = true
	}
	if exitRequested {
		db.Close()
		fmt.Println("Aurevoir !")
		os.Exit(1)
	}
}

/*
Cette fonction permet de clear le terminal
*/
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

/*
Cette fonction permet d'afficher une liste d'événements trier par date
*/
func displayEvents(Events map[int]Event.Event) {
	var sortBy string
	type kv struct {
		Key   int
		Value Event.Event
	}

	var ss []kv
	for k, v := range Events {
		ss = append(ss, kv{k, v})
	}

	for {
		fmt.Println("Comment voulez vous trier les événements ? (Titre,Date,Tag,Lieux)")
		sortBy = input.InputString()
		if sortBy != "Titre" && sortBy != "Date" && sortBy != "Tag" && sortBy != "Lieux" {
			fmt.Println("La valeur saisi est incorrect")
			continue
		}
		break
	}

	// Fonction de tri conditionnelle basée sur le critère choisi
	sort.Slice(ss, func(i, j int) bool {
		switch sortBy {
		case "Date":
			return ss[j].Value.StartDate.After(ss[i].Value.StartDate)
		case "Titre":
			return strings.ToLower(ss[i].Value.Title) < strings.ToLower(ss[j].Value.Title)
		case "Tag":
			return strings.ToLower(ss[i].Value.Tag) < strings.ToLower(ss[j].Value.Tag)
		case "Lieux":
			return strings.ToLower(ss[i].Value.Location) < strings.ToLower(ss[j].Value.Location)
		default:
			// Par défaut, tri par date
			return ss[j].Value.StartDate.After(ss[i].Value.StartDate)
		}
	})

	for _, event := range ss {
		fmt.Printf(color.Cyan+" %d."+color.Reset+" %s - %s - %s\n", event.Key, event.Value.Title, event.Value.StartDate.Format("2006-01-02 15:04"), event.Value.Tag)
	}
}

/*
Cette fonction permet d'afficher les détails d'un événement
*/
func displayEvent(event Event.Event) {

	clearScreen()
	fmt.Printf(color.Blue+"%s\n"+color.Reset, event.Title)
	fmt.Printf("Débute a    : %s\n", event.StartDate.Format("2006-01-02 15:04"))
	fmt.Printf("Termine a   : %s\n", event.EndDate.Format("2006-01-02 15:04"))
	fmt.Printf("durée       : %s\n", event.EndDate.Sub(event.StartDate))
	fmt.Printf("Tag         : %s\n", event.Tag)
	fmt.Printf("description : %s\n", event.Description)
	return
}

func displayModification() {
	fmt.Println()
	fmt.Println(color.White + "Que souhaitez vous modifiez : ")
	fmt.Println("--------------------------------------------------" + color.Reset)
	fmt.Println(color.Red + " 1." + color.Reset + "  pour modifier le titre")
	fmt.Println(color.Orange + " 2." + color.Reset + "  pour modifier la date de début")
	fmt.Println(color.Yellow + " 3." + color.Reset + "  pour modifier la date de fin")
	fmt.Println(color.Green + " 4." + color.Reset + "  pour modifier la localisation")
	fmt.Println(color.Blue + " 5." + color.Reset + "  pour modifier le tag")
	fmt.Println(color.Purple + " 6." + color.Reset + "  pour modifier la description")
	fmt.Println()
	fmt.Println("entrer votre choix :")
}

/*
Cette fonction permet de récupérer l'id de l'utilisateur avec son username et son password
*/
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

/*
Cette fonction permet de récupérer les événements qui vont commencer dans les prochaines 24h et de les renvoyer
*/
func upcomingEvents() map[int]Event.Event {
	events := make(map[int]Event.Event)
	for id, event := range eventsMap {
		if event.StartDate.After(time.Now()) && event.StartDate.Before(time.Now().Add(24*time.Hour)) {
			events[id] = event
		}
	}
	return events
}

/*
Cette fonction permet d'établir la connexion avec la base de données
et de réitérer ma requête si une erreur s'est produite
*/
func loopConnexionDb() {
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
	return
}

/*
Cette fonction permet la récupération de tous les événements liée à l'utilisateur connecter
elle permet de réitérer la requête  si une erreur s'est produite
*/
func loopGetEvents() {
	loop := "yes"
	for loop == "yes" {
		err := database.GetEvents(db, eventsMap, userId)
		if err != nil {
			log.Printf(color.Red+"%v"+color.Reset, err)
			fmt.Printf("Voulez-vous tenter de re récupérer les données ? : (yes/no)\n")
			loop = input.InputString()
		} else {
			break
		}
	}
	return
}

/*
Cette fonction permet la récupération des événements journaliers
elle gère les erreurs de saisie de l'utilisateur et renvoie exitRequested à true
si l'utilisateur souhaite quitter le programme
*/
func loopUpComingEvents() bool {
	exitRequested := false
	for {
		choice := input.InputString()
		if choice == "yes" {
			events := upcomingEvents()
			displayEvents(events)
			fmt.Println("Acces au menu : 1")
			fmt.Println("quitter : 2")
			for {
				res, err := input.InputInt()
				if err != nil || res != 1 && res != 2 {
					log.Printf(color.Red + "La valeur saisie est incorrecte" + color.Reset)
					continue
				}
				if res == 2 {
					exitRequested = true
				}
				break
			}
			break
		}
		if choice == "no" {
			break
		}
		log.Printf(color.Red + "La valeur saisie est incorrecte" + color.Reset)
		break
	}
	return exitRequested
}

func deleteEvent(id int) {
	err := database.DeleteEvent(db, id)
	if err != nil {
		log.Printf(color.Red+"%v"+color.Reset, err)
		return
	}
	delete(eventsMap, id)
	return
}

func createEvent() {
	for {
		clearScreen()
		event, err := Event.CreateEvent()
		if err != nil {
			log.Printf("%v", err)
		}

		var id int
		id, err = database.CreateEvent(event, db, userId)
		if err != nil {
			log.Printf("%v", err)
		}
		event.Id = id
		eventsMap[event.Id] = *event

		displayEvent(*event)

		for {
			fmt.Println("\ninformation correct ? (yes/no)")
			res := input.InputString()
			if res == "yes" {
				break
			}
			if res == "no" {
				updateEvent(id)
				continue
			}
			fmt.Println("valeur incorrecte")
		}

		var res string
		for res != "yes" && res != "no" {
			fmt.Println("Voulez vous saisir un autre evennement ? (yes/no)")
			res = input.InputString()
			if res == "yes" {
				break
			}
			if res == "no" {
				return
			}
			fmt.Println("valeur incorrecte")
		}
		continue
	}
}

func getEvents() {

	var id int
	var err error

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
		_, existe := eventsMap[id]
		if !existe {
			fmt.Println("L'id saisie n'existe pas")
			continue
		}
		displayEvent(eventsMap[id])
		break
	}

	fmt.Println("\nmodifier l'évenement : 1")
	fmt.Println("supprimer l'évenement : 2")
	fmt.Println("Revenir au menu : 3")
	fmt.Println("quitter : 4")

	var res int
	for {
		res, err = input.InputInt()
		if err != nil || res != 1 && res != 2 && res != 3 && res != 4 {
			fmt.Println("Saisi invalide")
			continue
		}
		break
	}
	switch res {
	case 1:
		updateEvent(id)
		break
	case 2:
		deleteEvent(id)
		break
	case 3:
		break
	case 4:
		exitRequested = true
	}
	return
}

func updateEvent(id int) {
	fmt.Println("Modifier un événement")

	var err error
	var choice int

	//Ceci affiche le menu de ce qui est modifiable
	//Pour modifier un/des champs il suffit d'entrer les chiffres associés aux choses modifiables
	// ex : 1 ou 145 ou 23
	displayModification()

	choice, err = input.InputInt()
	if err != nil {
		log.Printf(color.Red + "La valeur saisie est incorrecte" + color.Reset)
	}

	event := eventsMap[id]

	//Switch qui permet de récuperer les champs que l'utilisateur veut modifier
	//Il prend en parametres la saisie du dessus et l'adresse de la struct Event get auparavant
	Event.UpdateChoices(&event, choice)

	//Requete SQL pour effectuer la modification
	err = database.PatchEvent(&event, db)

	if err != nil {
		log.Printf("%v", err)
	}
	eventsMap[id] = event
	return
}

func searchEvent(query string) map[int]Event.Event {
	matchingEvents := make(map[int]Event.Event)

	for _, event := range eventsMap {
		// Si la requête est trouvée dans n'importe quel attribut de l'événement, ajouter à la liste
		if containsEvent(event, query) {
			matchingEvents[event.Id] = event
		}
	}

	return matchingEvents
}

func containsEvent(event Event.Event, query string) bool {
	// Rechercher la requête dans tous les attributs de l'événement
	return subString(event.Title, query) ||
		subString(event.Location, query) ||
		subString(event.Tag, query)
}

func subString(str, substr string) bool {
	// Vérifier si substr est une sous-chaîne de str (insensible à la casse)
	str, substr = strings.ToLower(str), strings.ToLower(substr)
	return strings.Contains(str, substr)
}
