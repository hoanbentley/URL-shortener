package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/hoanbentley/URL-shortener/internal/controller"
	"net/http"
)

var db string

func main() {
	flag.StringVar(&db, "db", "", "db run")
	flag.Parse()
	router := mux.NewRouter()
	todo := controller.NewToDoService()
	router.HandleFunc("/login", todo.GetAuthToken).Methods(http.MethodGet)
	router.HandleFunc("/admin/list", todo.ListUrl).Methods(http.MethodGet)
	router.HandleFunc("/admin/search", todo.SearchUrl).Methods(http.MethodPost)
	router.HandleFunc("/admin/delete/{id}", todo.DeleteUrl).Methods(http.MethodDelete)
	router.HandleFunc("/create", todo.CreateUrl).Methods(http.MethodPost)
	router.HandleFunc("/{id}", todo.RedirectUrl).Methods(http.MethodGet)
	http.ListenAndServe(":8080", router)

	// statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS urls (id INTEGER PRIMARY KEY, short_code TEXT, full_url TEXT, expiry INTEGER, create_date INTEGER, number_of_hits INTEGER)")
	/*database, _ := sql.Open("sqlite3", "./rabbit.db")
	statement, _ := database.Prepare("DELETE FROM urls;")
	statement.Exec()*/

	/*database, _ := sql.Open("sqlite3", "./rabbit.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS users (user_id TEXT, password TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO users (user_id, password) VALUES (?, ?)")
	statement.Exec("admin", "admin")
	rows, _ := database.Query("SELECT user_id, password FROM users")
	var user_id string
	var password string
	for rows.Next() {
		rows.Scan(&user_id, &password)
		fmt.Println(user_id + " " + password)
	}*/

	/*statement, _ = database.Prepare("INSERT INTO urls (short_code, full_url, expiry, number_of_hits) VALUES (?, ?, ?, ?)")
	statement.Exec("Nic", "Raboy", "hoantk", "1")
	rows, _ := database.Query("SELECT id, shortcode,fullurl, expiry,numberofhits  FROM urls")
	var id int
	var shortcode string
	var fullurl string
	var expiry string
	var numberofhits string
	for rows.Next() {
		rows.Scan(&id, &shortcode, &fullurl, &expiry, &numberofhits)
		fmt.Println(strconv.Itoa(id) + ": " + shortcode + " " + fullurl + " " + expiry + " " + numberofhits)
	}*/
}
