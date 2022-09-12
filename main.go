package main

import (
	"student-management/dao"
	"student-management/model"
	"encoding/json"
	"log"
	"net/http"
)
//var ead = dao.EmployeeDAO{}
var studenDao = dao.StudentDAO{}

func CreateNewEmployee(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		respondWithError(w, http.StatusBadRequest, "Invalid method")
	}

	var student model.Student

	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
	}

	if err := studenDao.Insert(student); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
	} else {
		respondWithJson(w, http.StatusAccepted, map[string]string{
			"message": "Record inserted successfully",
		})
	}
}


func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}



func init() {
	//studenDao.Server = "mongodb+srv://m001-student:m001-mongodb-basics@sandbox.t7xjyrg.mongodb.net/test"
	studenDao.Server = "mongodb://localhost:27017"
	studenDao.Database = "student_management"
	studenDao.Collection = "student"

	studenDao.Connect()
}

func main() {
	//http.HandleFunc("/employee/", getEmployeesById)
	http.HandleFunc("/add-employee", CreateNewEmployee)
//	http.HandleFunc("/delete-employee", deleteEmployeeById)
//	http.HandleFunc("/update-employee", updateEmployeeById)

	log.Fatal(http.ListenAndServe(":8080", nil))
}