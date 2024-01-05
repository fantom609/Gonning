# Gonning

## login & mot de passe
- ('Noe', 'Noe')
- ('Arthur', 'Arthur')
- ('Nicolas', 'Nicolas')
- ('Lena', 'Lena')

## Base de données
 Toutes les infos pour se connecter à la db sont dans config.txt

### Fonctionnalités du programme

1. [Créer un nouvel événement](#creer-evenement)
2. [Visualiser les événements](#visualiser-evenements)
3. [Visualiser un événement par l'id](#visualiser-evenement-par-id)
4. [Enregistrer les événements](#enregistrer-evenements)
5. [Rechercher un événement](#rechercher-evenement)
6. [Quitter](#quitter)

Egalement chaque information envoyée à la db est verifiée et **si elles ne sont pas correctes alors le programme throw une erreur.**

Après s'être connecté, vous avez directement la possibilité de voir les événements que vous avez dans les prochaines 24H, vous avez la possibilité de les voir en détail. Ce qui veut donc dire que **dés la connexion établie les événements sont get et stockés dans un map**.

### Créer un événement

Utilise la fonction CreateEvent du package Event puis la fonction CreateEvent du package database.  
Elle permet d'ajouter un titre, une date de début/fin, lieu, catégorie, description.

### Visualiser les événements un événement

Les fonctions utilisées pour get un ou des events sont soit `getEvents` ou alors `func getEvent(id int)` du main Event puis `GetEvents(db *sql.DB, events map[int]Event.Event, userId int)` du package database.  
Cela permet de visualiser les événements dans leur intégralité, vous pouvez les trier par Titre, date, tag (catégorie) ou lieu. Vous pouvez get un événement en entrant sont id ou revenir au menu principal.
Sinon vous pouvez voir un événement de manière plus détaillé vous ouvrant donc un nouveau menu : 

- Modifier l'événement : 1  
- Supprimer l'événement : 2  
- Revenir au menu : 3  
- Quitter : 4  

#### Modifier un événement
En effet pour modifier un événement vous devez forcément le get avant de pouvoir le faire. Vous pouvez tout modifier, pour cela on passe par `UpdateChoices(Event *Event, choices int)` qui permet de récupérer tout ce que l'utilisateur veut modifier: 
1.  pour modifier le titre
2.  pour modifier la date de début
3.  pour modifier la date de fin
4.  pour modifier la localisation
5.  pour modifier le tag
6.  pour modifier la description

entrer votre choix : 13 (me permet de modifier le titre ainsi que la date de fin)

Une fois que j'ai récupéré en input ce qu'on veut modifier cela affiche en conséquences ce qu'on veut de nouveau écrire. Une fois que les informations sont récupérées c'est avec la fonction `PatchEvent(event *Event.Event, db *sql.DB)` de database qu'on peut update un event.

#### Supprimer un événement

Vu qu'on a déjà get l'événement avant, on récupère son id, on a une étape de vérification qui nous permet d'être sûr si oui ou non on veut supprimer l'événement, cela appelle la fonction `DeleteEvent(db *sql.DB, id int)` de database.

### Visualiser un événement par l'id

Permet de voir un événement avec tous les détails à l'aide de son id, étant donné que les événements sont stockés dans un map, nous utilisons juste la fonction `getEvent(id int)` du main et qui permet de nouveau d'ouvrir le menu précedement présenté.

### Enregistrer les événements

Enregistrer les événements en format json via la fonction `jsonEvents(events *map[int]Event.Event)` du main et qui utilise également `type EventsWrapper struct`.

### Rechercher un événement

Dans le main la fonction `searchEvent(query string) map[int]Event.Event` est appelée qui appelle a son tour `containsEvent(event Event.Event, query string)`. Cela nous permet de chercher si le str envoyé par l'utilisateur est présent dans le titre, la catégorie ou alors le lieu. De plus la casse est gérée via la fonction `containsEvent(event Event.Event, query string)`



