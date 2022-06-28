package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func connectDB() (connection *sql.DB) {
	driver := "mysql"
	username := "root"
	password := ""
	name := "sistema"
	connection, err := sql.Open(driver, username+":"+password+"@tcp(127.0.0.1)/"+name)
	if err != nil {
		panic(err.Error())
	}
	return connection
}

var templates = template.Must(template.ParseGlob("templates/*"))

func main() {
	http.HandleFunc("/", Start)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/update", Update)
	log.Println("Starting server")
	http.ListenAndServe(":8080", nil)
}

type Employee struct {
	Id    int
	Name  string
	Email string
}

func Start(w http.ResponseWriter, r *http.Request) {

	connectionEstablished := connectDB()
	register, err := connectionEstablished.Query("SELECT 	* FROM empleados")

	if err != nil {
		panic(err.Error())
	}
	employee := Employee{}
	arrayEmployees := []Employee{}

	for register.Next() {
		var id int
		var name, email string
		err = register.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}
		employee.Id = id
		employee.Name = name
		employee.Email = email

		arrayEmployees = append(arrayEmployees, employee)
	}
	// fmt.Println(arrayEmployees)

	// fmt.Fprintf(w, "Welcome")
	templates.ExecuteTemplate(w, "start", arrayEmployees)
}
func Create(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create", nil)
}
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")
		t := 301
		connectionEstablished := connectDB()
		insertRegister, err := connectionEstablished.Prepare("INSERT INTO empleados (nombre,correo) VALUES (?,?) ")

		if err != nil {
			panic(err.Error())
		}
		insertRegister.Exec(name, email)
		http.Redirect(w, r, "/", t)
	}
}
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		email := r.FormValue("email")
		t := 301
		connectionEstablished := connectDB()
		updateRegister, err := connectionEstablished.Prepare("UPDATE empleados SET nombre=?, correo=? WHERE id=?")

		if err != nil {
			panic(err.Error())
		}
		updateRegister.Exec(name, email, id)
		http.Redirect(w, r, "/", t)
	}
}
func Delete(w http.ResponseWriter, r *http.Request) {
	idEmployee := r.URL.Query().Get("id")
	// fmt.Println(idEmployee)
	t := 301
	connectionEstablished := connectDB()
	deleteRegister, err := connectionEstablished.Prepare("DELETE FROM empleados WHERE id=? ")

	if err != nil {
		panic(err.Error())
	}
	deleteRegister.Exec(idEmployee)
	http.Redirect(w, r, "/", t)
}
func Edit(w http.ResponseWriter, r *http.Request) {
	idEmployee := r.URL.Query().Get("id")
	fmt.Println(idEmployee)

	connectionEstablished := connectDB()
	registers, err := connectionEstablished.Query("SELECT * FROM empleados WHERE id=?", idEmployee)
	employee := Employee{}

	for registers.Next() {
		var id int
		var name, email string
		err = registers.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}
		employee.Id = id
		employee.Name = name
		employee.Email = email
	}
	fmt.Println(employee)
	templates.ExecuteTemplate(w, "edit", employee)
}
