package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	models "urlshorten/src/backend/db"
	"urlshorten/src/backend/encode"

	"github.com/gorilla/mux"
)

func main() {

	// Initialize the DB.
	db, err := models.InitDB()
	if err != nil {
		log.Println("DB connection error")
		panic(err)
	}

	dbclient := &DBClient{db: db}
	if err != nil {
		log.Println("DB client error")
		panic(err)
	}
	defer db.Close()

	// HTTP Handlers
	r := mux.NewRouter()
	r.HandleFunc("/", dbclient.GenerateShortURL).Methods("POST")
	r.HandleFunc("/{encoded_string:[a-zA-Z0-9]*}", dbclient.GetOriginalURL).Methods("GET")
	//fileServer := http.FileServer(http.Dir("./ui/client/public/"))
	//r.Handle("/", http.StripPrefix("/", fileServer))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	})
	http.ListenAndServe(":8080", r)

}

// Check if the URL is valid or not.
func isValidURL(toTest string) error {

	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return err
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return err
	}

	return nil
}

// DBClient is used for the database.
type DBClient struct {
	db *sql.DB
}

//GetOriginalURL searches the encoded URL and redirects to original URL.
func (driver *DBClient) GetOriginalURL(w http.ResponseWriter, r *http.Request) {

	var res string
	vars := mux.Vars(r)
	err := driver.db.QueryRow("SELECT url FROM web_url WHERE id = $1", vars["encoded_string"]).Scan(&res)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		http.Redirect(w, r, res, 301)
	}
	log.Printf("Getting URL: %s %s %s\n", r.RemoteAddr, r.Method, r.URL)
}

//GenerateShortURL generates the encoded URL to use.
func (driver *DBClient) GenerateShortURL(w http.ResponseWriter, r *http.Request) {

	var res string
	postBody, _ := ioutil.ReadAll(r.Body)
	url := string(postBody)
	err := isValidURL(url)

	if err != nil {
		log.Println("Invalid URL")
		w.Write([]byte("Invalid URL"))
		return
	}

	hashID := encode.HashLink(url)
	responseMap := map[string]interface{}{"encoded_string": "http://localhost:8080/" + hashID}
	err = driver.db.QueryRow("INSERT INTO web_url(id,url) VALUES($1,$2) RETURNING id", hashID, url).Scan(&res)

	if err != nil {
		log.Println(err)
		log.Println("Couldn't insert into the database.")
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(responseMap)
		w.Write(response)

	}

	log.Printf("Generating URL: %s %s %s\n", r.RemoteAddr, r.Method, r.URL)

}
