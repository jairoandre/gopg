package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"fmt"
	_ "github.com/lib/pq"
)

type Item struct {
	ID    int
	Attrs Attrs
}

type Attrs struct {
	Name    string                 `json:"name,omitempty"`
	Content map[string]interface{} `json:"content,omitempty"`
}


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Ehdpoint hit: HomePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
// Make the Attrs struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (a *Attrs) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// Make the Attrs struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (a Attrs) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func writeOnDb() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/keep?sslmode=disable")

	attrs := new(Attrs)
	attrs.Name = "Teste 2"
	attrs.Content = map[string]interface{}{"type": "document", "value": 3}

	_, err = db.Exec("insert into items (attrs) values($1)", attrs)

	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	handleRequests()
}
