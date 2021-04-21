package main

import (
	"flag"
	"github.com/hoanbentley/URL-shortener/internal/controller"
	"net/http"
)

var db string

func main() {
	flag.StringVar(&db, "db", "", "db run")
	flag.Parse()
	http.ListenAndServe(":8080", controller.NewToDoService(db))

	/*database, _ := sql.Open("sqlite3", "./rabbit.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS urls (id INTEGER PRIMARY KEY, short_code TEXT, full_url TEXT, expiry TEXT, number_of_hits TEXT)")
	statement.Exec()*/
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
