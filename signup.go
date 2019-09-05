package authServer

import(
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"errors"
	"unicode"
	"golang.org/x/crypto/bcrypt"
	
)



type SignUpRequest struct{
	Username string `json:username`
	Password string `json:password`
}

func SignUpHandler(w http.ResponseWriter, r *http.Request){
	var signup SignUpRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Fatal(err)
		return
	}
	
	json.Unmarshal(body, &signup)
	check := SignUpCheck(signup.Username, signup.Password)
	if check != nil{
		http.Error(w, check.Error(), 400)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(signup.Password), bcrypt.MinCost)
	if err != nil{
		http.Error(w, err.Error(), 400)
		return
	}
	stmtString := fmt.Sprintf("INSERT INTO %s(username,password) VALUES (?,?)", TbName)
	stmt, err := Db.Prepare(stmtString)
	if err != nil{
		http.Error(w, err.Error(), 400)
		return
	}
	_, err = stmt.Exec(signup.Username, hash)
	if err != nil{
		http.Error(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Signup successful"))
}

func SignUpCheck(user string, pass string) error{
	if len(user) == 0{
		return errors.New("Username cannot be blank")
	}
	stmtString := fmt.Sprintf("SELECT * FROM %s WHERE username = ?", TbName)
	stmt, err := Db.Prepare(stmtString)
	defer stmt.Close()
	if err != nil{
		log.Fatal(err)
	}
	row := stmt.QueryRow(user)
	
	var s string
	if row.Scan(s) != sql.ErrNoRows{
		return errors.New("Username already in use, please select a different one")
	}
	if len(pass) < 8{
		return errors.New("Password length is too short, minimum of 8 characters")
	}
	uppercheck := false
	lowercheck := false
	numbercheck := false
	for _, char := range pass{
		if unicode.IsNumber(char){
			numbercheck = true
		}
		if unicode.IsLower(char){
			lowercheck = true
		}
		if unicode.IsUpper(char){
			uppercheck = true
		}
	}
	if numbercheck != true{
		return errors.New("Password requires at least one number")
	}
	if lowercheck != true{
		return errors.New("Password requires at least one lowercase character")
	}
	if uppercheck != true{
		return errors.New("Password requires at least one uppercase character")
	}
	return nil
}
