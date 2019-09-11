package main

import(
	"github.com/vivevincere/auth-server"
  _ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
)


func main(){
	SecretKeyParam := []byte("okb7Ey_DALa6JRGmnB1oSeGIrnfqLSR9nYHFk10ZlMpDVNdJ2ESdwpCSAAWNxEGhn0ZagVFoiWReiw_jmaUy5PkViS9k3QnKwOjcVevXSDG9CHF1JGoinoAuQRY-9NLtqbCgD_OTBAYRv_q7bDy2snw7Ak-ije85VuzzkFpLBOuy4rinWlNSjvfWOYca9w70axqt2TXAzrzPNeZjDOKSknWouwFbItJBtL3brZoXCg0VGDCk70lTqa91RaGZcJTBWjQqfRnyaojEzZtEBJ5MU5hG1PgyPzQMAFosvh1kt93AiUF0xnOa_5sU9e9wNhmgTo9m5SCtRTkeHO3DWLrSXg")
	TableName := "login_details"
	Database, err := sql.Open("mysql", "root:password1@tcp(localhost:3306)/AuthServer")
	if err != nil{
		log.Fatal(err)
	}
	authServer.StartServer(Database, TableName, SecretKeyParam, "/signup", "/login", "localhost:8080")
}