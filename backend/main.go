package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// ********** SETUP ***********

const (
	// DB_USER     = "faraway"
	DB_USER     = "postgres"
	DB_PASSWORD = "green100world"
	DB_NAME     = "faraway"
	DB_PORT     = 5454
	TABLE_NAME  = "users"

	CORS_HOST = "http://localhost:3000"
)

// ********** SETUP ***********

func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s port=%d password=%s dbname=%s sslmode=disable", DB_USER, DB_PORT, DB_PASSWORD, DB_NAME)

	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

type Users struct {
	Uid     int    `json:"uid"`
	UName   string `json:"name"`
	UEmail  string `json:"email"`
	UPhone  string `json:"phone"`
	UGender string `json:"gender"`
}

type JsonResponse struct {
	Type    string  `json:"type"`
	Data    []Users `json:"data"`
	Message string  `json:"message"`
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", CORS_HOST)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	// w.Header().Set("Access-Control-Max-Age", "600")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")

	db := setupDB()

	// printMessage("Getting users...")

	rows, err := db.Query("SELECT * FROM " + TABLE_NAME)

	checkErr(err)

	// var response []JsonResponse
	var users []Users

	for rows.Next() {
		var uid int
		var name string
		var email string
		var phone string
		var gender string

		err = rows.Scan(&uid, &name, &email, &phone, &gender)

		checkErr(err)

		users = append(users, Users{Uid: uid, UName: name, UEmail: email, UPhone: phone, UGender: gender})
	}

	var response = JsonResponse{Type: "success", Data: users}

	// json.NewEncoder(w).Encode(response)
	json.NewEncoder(w).Encode(response.Data)

	defer rows.Close()

}

func RemoveUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", CORS_HOST)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	userID := params["userId"]

	var response = JsonResponse{}

	if userID == "" {
		response = JsonResponse{Type: "error", Message: "You are missing userID parameter."}
	} else {
		db := setupDB()

		printMessage("Deleting user from DB")

		_, err := db.Exec("DELETE FROM users where uid=$1", userID)

		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The user has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)

}

func EditUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", CORS_HOST)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	fmt.Println("=============", "PUT-request is in process", "=============")

	params := mux.Vars(r)

	bodyBytes, _ := ioutil.ReadAll(r.Body)

	userID := params["userId"]

	var response = JsonResponse{}

	db := setupDB()

	type updatedUser struct {
		Uid    int    `json:"uid"`
		Name   string `json:"name"`
		Email  string `json:"email"`
		Phone  string `json:"phone"`
		Gender string `json:"gender"`
	}

	// type Request struct {
	// 	Operation string `json:"operation"`
	// 	Key       string `json:"key"`
	// 	Value     string `json:"value"`
	// }

	// s := string(`{"action":"get", "key":"example"}`)
	// data := Request{}
	// json.Unmarshal([]byte(s), &data)
	// fmt.Printf("string S: %s", s)

	body := updatedUser{}
	json.Unmarshal([]byte(bodyBytes), &body)

	sqlRequest := `UPDATE ` + TABLE_NAME +
		` SET name = '` + body.Name + `', email = '` + body.Email + `', phone = '` + body.Phone + `', gender = '` + body.Gender + `' WHERE uid = ` + userID

	fmt.Println("body: ", body)
	// fmt.Println(sqlRequest)
	// fmt.Println(err)

	db.QueryRow(sqlRequest)

	response = JsonResponse{Type: "success", Message: "The user has been inserted successfully!"}

	json.NewEncoder(w).Encode(response)
}

func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")

	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr

}

func logger(arg string) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Printf("Request was: " + arg)
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
	logger("root-request to '/'")
}

func doNothing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "dirty hack to avoid requesting favicon and doubling each request at working from browser")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeLink)

	router.HandleFunc("/getallusers", GetAllUsers).Methods("GET", "OPTIONS")

	router.HandleFunc("/edituser/{userId}", EditUser).Methods("PUT", "OPTIONS")

	router.HandleFunc("/removeuser/{userId}", RemoveUser).Methods("DELETE", "OPTIONS")

	// dirty hack for favicon
	router.HandleFunc("/favicon.ico", doNothing)

	fmt.Println("API Server is living on localhost:8080 and ready to actions!")

	log.Fatal(http.ListenAndServe(":8080", router))

}
