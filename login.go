package authserver

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type LoginRequest struct {
	Username string `json:username`
	Password string `json:password`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var login LoginRequest
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &login)
	check := loginCheck(login.Username, login.Password)
	if check != nil {
		http.Error(w, check.Error(), 400)
		return
	}
	curTime := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"aud": "fillwaud",
		"exp": curTime.Add(time.Hour * 3),
	})
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	var g map[string]string
	g = make(map[string]string)
	g["token"] = tokenString
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(g)

	return
}

func loginCheck(user string, pass string) error {
	stmtString := fmt.Sprintf("SELECT password FROM %s WHERE username = ?", TbName)
	stmt, err := Db.Prepare(stmtString)
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(user)
	var s []byte
	noRowCheck := row.Scan(&s)
	if noRowCheck == sql.ErrNoRows {
		return errors.New("Invalid username/password")
	}
	err = bcrypt.CompareHashAndPassword(s, []byte(pass))
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Invalid username/password1")
	}
	return nil

}
