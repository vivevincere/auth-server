package authserver

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

var SecretKey []byte
var TbName string
var Db *sql.DB

func StartServer(SqlConnection *sql.DB, tableName string, secretKeyParam []byte, signup string, login string, base string) {
	SecretKey = secretKeyParam
	Db = SqlConnection
	TbName = tableName

	defer Db.Close()
	r := mux.NewRouter()

	r.HandleFunc(signup, SignUpHandler)
	r.HandleFunc(login, LoginHandler)
	http.ListenAndServe(base, r)
}
