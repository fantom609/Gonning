package database

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
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

func ConnectionDatabase() *sql.DB {

	config, err := ReadConfig()
	if err != nil {
		log.Fatalf("Une erreur est survenue lors de la lecture du fichier : %v", err)
	}

	db, err := sql.Open("postgres", config)

	if err != nil {
		log.Fatal("Une erreur est survenue lors de la connection a la base de données :", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("erreur lors de la vérification de la connexion à la base de données : %v", err)
	}

	return db
}

func InitDb(db *sql.DB) {
	req := "CREATE TABLE EVENT (" +
		"id serial PRIMARY KEY," +
		"title TEXT," +
		"startDate TEXT," +
		"endDate TEXT," +
		"location TEXT," +
		"tag TEXT" +
		");"
	_, err := db.Exec(req)

	if err != nil {
		log.Fatal(" ", err)
	}

}

func CreateEvent() {
	var id int
	db := ConnectionDatabase()
	req := "INSERT INTO Event (title,startDate,endDate,location,tag) VALUES ($1,$2,$3,$4,$5) RETURNING id"
	err := db.QueryRow(req, "Alice", "2023-03-11", "2023-03-17", "la", "ici").Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d", id)
}

func PatchEvent() {}

func DeleteEvent() {}

func GetEvent() {}

func GetEvents() {}
