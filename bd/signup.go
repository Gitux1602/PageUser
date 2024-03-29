package bd

import (
	"Gambit/Users/PageUser/models"
	"Gambit/Users/PageUser/tools"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sig models.SignUp) error {

	fmt.Println("Comienza Registro")
	err := DbConect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.FechaMySQL() + "')"
	fmt.Println(sentencia)

	_, err = Db.Exec(sentencia)
	if err != nil {

		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Sign Up -- Ejecución Exitosa")
	return nil
}
