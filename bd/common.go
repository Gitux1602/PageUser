package bd

import (
	"database/sql"
	"fmt"
	"os"

	"Gambit/Users/PageUser/models"
	"Gambit/Users/PageUser/secretm"

	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}
func DbConect() error {

	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {

		println(err.Error())
		return err
	}
	err = Db.Ping()
	if err != nil {
		println(err.Error())
	}

	println("Conexion exitosa a la DB")
	return nil
}

func ConnStr(claves models.SecretRDSJson) string {

	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	dbName = "gambit"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
}
