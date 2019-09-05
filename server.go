package authServer

import(
	"net/http"
  _ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/gorilla/mux"

)
var SecretKey []byte  
var TbName string  
var Db *sql.DB




func StartServer(SqlConnection *sql.DB, tableName string, secretKeyParam []byte, signup string, login string, base string){
	SecretKey = secretKeyParam
	Db = SqlConnection
	TbName = tableName

	defer Db.Close()
	r := mux.NewRouter()

	r.HandleFunc(signup,SignUpHandler)
	r.HandleFunc(login, LoginHandler)
	http.ListenAndServe(base, r)
}

