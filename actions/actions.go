package actions

import (
	"net/http"
	"encoding/json"
	"../models"
	"fmt"
	"log"
	//"github.com/gorilla/mux"
	//"strconv"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	"strconv"
	//"image/color"
	"gopkg.in/mgo.v2/bson"
)


func getSession() *mgo.Session{
	session,error := mgo.Dial("mongodb://localhost")
	if(error != nil){
		panic(error)
	}
	return session
}

var collection = getSession().DB("DbPersonas").C("Personas")


func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"hola servidor")
}

func PersonaList(w http.ResponseWriter, r *http.Request){

	var results []models.Persona
	err := collection.Find(nil).Sort("+id").All(&results)
	if(err != nil){
		log.Fatal(err)
	}else{
		fmt.Println("Resultados : ", results)
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)

}


func GetPersonaById(w http.ResponseWriter, r *http.Request){

	params := mux.Vars(r)
	personaId := params["id"]

	if !bson.IsObjectIdHex(personaId){
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid := bson.ObjectIdHex(personaId)


	var results models.Persona
	err := collection.FindId(oid).One(&results)
	if(err != nil){
		log.Fatal(err)
	}else{
		fmt.Println("Resultados : ", results)
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

func GetPersonaId2(w http.ResponseWriter, r *http.Request){

	params := mux.Vars(r)
	personaId2 := params["id2"]

	fmt.Println((personaId2))

	var results models.Persona
	idConv,_ := strconv.Atoi(personaId2);

	err := collection.Find(bson.M{"id":idConv}).One(&results)
	if(err != nil){
		fmt.Println((err))
		log.Fatal(err)
	}else{
		fmt.Println("Resultados : ", results)
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}




func AddPersona(w http.ResponseWriter, r *http.Request)  {
	decoder := json.NewDecoder(r.Body)

	var persona_data models.Persona
	err := decoder.Decode(&persona_data)
	if(err != nil){
		panic(err)
	}
	defer r.Body.Close()

	collection.Insert(persona_data)

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(persona_data)

}


