package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while loading config file %s", err)
	}

	appPort := viper.Get("APP_PORT")
	log.Println("Server started on port " + appPort.(string))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		person := Person{Name: "John Doe", Age: 30}
		json.NewEncoder(w).Encode(person)
	})

	http.ListenAndServe(":8080", nil)
}
