package authserver

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

var SecretKey []byte
var TbName string
var Db *sql.DB

func HomeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Reached Home")
	w.Write([]byte("Home reached Successfully"))

}


func StartServer(SqlConnection *sql.DB, tableName string, secretKeyParam []byte, signup string, login string, base string) {
	SecretKey = secretKeyParam
	Db = SqlConnection
	TbName = tableName

	defer Db.Close()
	r := mux.NewRouter()

	r.HandleFunc(signup, SignUpHandler)
	r.HandleFunc(login, LoginHandler)
	r.HandleFunc("/", HomeHandler)
	http.ListenAndServe(base, r)
}
