package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	json.NewDecoder(r.Body).Decode(&emp)
	Database.Create(&emp)
	json.NewEncoder(w).Encode(emp)
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees []Employee
	Database.Find(&employees)
	json.NewEncoder(w).Encode(employees)

}

func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees Employee
	Database.First(&employees, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode(employees)

}

func UpdateEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees Employee
	Database.First(&employees, mux.Vars(r)["eid"])
	json.NewDecoder(r.Body).Decode(&employees)
	Database.Save(&employees)
	json.NewEncoder(w).Encode(employees)

}

func DeleteEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees Employee
	Database.Delete(&employees, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode("Employee are deleted")

}