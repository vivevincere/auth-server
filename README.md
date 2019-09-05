# auth-server

Authentication server written in Go with bcrypt hashed passwords saved to a mySQL server. Returns a HS256 JWT Token when username and password is verified against the saved hash.

Call StartServer() with the relevant parameters for use.

jwt package: 	https://github.com/dgrijalva/jwt-go 

mysql driver: https://github.com/go-sql-driver/mysql 
