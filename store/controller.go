package store

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	products := c.Repository.GetProducts() // list of all products
	// log.Println(products)
	data, _ := json.Marshal(products)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// AddProduct POST /
func (c *Controller) AddProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request

	log.Println(body)

	if err != nil {
		log.Fatalln("Error AddProduct", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddProduct", err)
	}

	if err := json.Unmarshal(body, &product); err != nil { // unmarshall body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddProduct unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	log.Println(product)
	success := c.Repository.AddProduct(product) // adds the product to the DB
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}
