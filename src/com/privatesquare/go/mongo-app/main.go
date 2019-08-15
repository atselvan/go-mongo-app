package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Employee struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	EmpID     int           `bson:"empId" json:"empId"`
	FirstName string        `bson:"firstName" json:"firstName"`
	LastName  string        `bson:"lastName" json:"lastName"`
	Age       int           `bson:"age" json:"age"`
	Address   struct {
		HouseNumber string `bson:"houseNumber" json:"houseNumber"`
		Street      string `bson:"street" json:"street"`
		City        string `bson:"city" json:"city"`
		State       string `bson:"state" json:"state"`
		PostalCode  string `bson:"postalCode" json:"postalCode"`
		Country     string `bson:"country" json:"country"`
	} `bson:"address" json:"address"`
	Phone []string `bson:"phone" json:"phone"`
}

var db *mgo.Database

const (
	host       = "mongo-db"
	database   = "privatesquare"
	collection = "employees"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employees", GetEmployeesHandler).Methods("GET")
	router.HandleFunc("/employees/{empId}", GetEmployeeHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func GetEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(GetEmployeesService())
}

func GetEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	empId, err := strconv.Atoi(params["empId"])
	HandleError(err, "Error converting string to int")
	employee := GetEmployeeService(empId)
	if employee.EmpID != 0 {
		json.NewEncoder(w).Encode(employee)
	} else {
		http.Error(w, fmt.Sprintf("Employee with id %s does not exist.", params["empId"]), http.StatusNotFound)
	}
}

func GetEmployeesService() []Employee {
	dbConnect()
	var employees []Employee
	err := db.C(collection).Find(bson.M{}).All(&employees)
	HandleError(err, "Error getting the data")
	return employees
}

func GetEmployeeService(empId int) Employee {
	dbConnect()
	var employee Employee
	err := db.C(collection).Find(bson.M{"empId": empId}).One(&employee)
	HandleError(err, "Error getting the data")
	return employee
}

func dbConnect() {
	session, err := mgo.Dial(host)
	HandleError(err, "Error dialing to the database")
	db = session.DB(database)
}

func logger(message string) {
	logger := log.New(os.Stdout, "[INFO]:", log.LstdFlags)
	logger.Println(message)
}

func HandleError(err error, errorMessage string) {
	if err != nil {
		logger(errorMessage)
		logger(err.Error())
	}
}
