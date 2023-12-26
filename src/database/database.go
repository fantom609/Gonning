package database

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"main/src/Event"
	"os"
	"strings"
)

func ReadConfig() (string, error) {
	configFile := "config.txt"

	// Ouvrez le fichier en lecture seulement
	file, err := os.Open(configFile)
	if err != nil {
		return "", errors.New("le fichier n'a pu etre ouvert")
	}
	defer file.Close()

	config := make(map[string]string)
	scanner := bufio.NewScanner(file)

	// Parcourez chaque ligne du fichier
	for scanner.Scan() {
		// Divisez la ligne en clé et valeur en fonction du signe égal (=)
		parts := strings.Split(scanner.Text(), "=")
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			switch key {
			case "host", "user", "password", "dbName":
				config[key] = value
			}
		}
	}

	for _, value := range config {
		if value == "" {
			return "", errors.New("la configuration n'est pas correcte")
		}
	}

	configString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config["host"], config["user"], config["password"], config["dbName"])

	return configString, nil
}

func ConnectionDatabase() (*sql.DB, error) {

	config, err := ReadConfig()
	if err != nil {
		return nil, fmt.Errorf("une erreur est survenue lors de la lecture du fichier : %v", err)
	}

	db, err := sql.Open("postgres", config)

	if err != nil {
		return nil, fmt.Errorf("une erreur est survenue lors de la connection a la base de données : %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la vérification de la connexion à la base de données : %v", err)
	}

	return db, nil
}

func InitDb(db *sql.DB) {
	req := "CREATE TABLE UTILISATEUR (" +
		"id serial PRIMARY KEY," +
		"username TEXT," +
		"password TEXT" +
		");"

	req += "CREATE TABLE EVENT (" +
		"id serial PRIMARY KEY," +
		"title TEXT," +
		"startDate TIMESTAMP," +
		"endDate TIMESTAMP," +
		"location TEXT," +
		"tag TEXT," +
		"description TEXT," +
		"id_utilisateur INTEGER REFERENCES UTILISATEUR(id) ON DELETE CASCADE" +
		");"

	req += "INSERT INTO UTILISATEUR (username, password) VALUES " +
		"('Noe', 'Noe')," +
		"('Arthur', 'Arthur')," +
		"('Nicolas', 'Nicolas')," +
		"('Lena', 'Lena');"
	_, err := db.Exec(req)

	if err != nil {
		log.Fatal(" ", err)
	}

}

func CreateEvent(event *Event.Event, db *sql.DB) (int, error) {
	var id int
	req := "INSERT INTO Event (title,startdate,enddate,location,tag,description) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id"
	err := db.QueryRow(req, event.Title, event.StartDate, event.EndDate, event.Location, event.Tag, event.Description).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("erreur lors de l'insertion des données' : %v", err)
	}
	return id, nil
}

func PatchEvent() {}

func DeleteEvent(db *sql.DB, id int) error {

	result, err := db.Exec("DELETE FROM EVENT WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("erreur lors de la requete %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erreur lors de la récupération du nombre de lignes affectées %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("aucun evenement n'a ete supprimé pour l'ID %d", id)
	}

	fmt.Println("L'evenement a ete supprime")

	return nil
}

func GetEvent() {}

func GetEvents(db *sql.DB) ([]Event.Event, error) {

	var events []Event.Event

	req := "SELECT id,title,startdate,enddate,location,tag,description FROM event"
	rows, err := db.Query(req)
	if err != nil {
		return []Event.Event{}, fmt.Errorf("erreur lors de la requéte %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var event Event.Event
		err := rows.Scan(&event.Id, &event.Title, &event.StartDate, &event.EndDate, &event.Location, &event.Tag, &event.Description)
		if err != nil {
			return []Event.Event{}, err
		}
		events = append(events, event)
	}

	return events, nil
}
