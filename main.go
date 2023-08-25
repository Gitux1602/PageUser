package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"Gambit/Users/PageUser/awsgo"
	"Gambit/Users/PageUser/bd"
	"Gambit/Users/PageUser/models"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutoLambda)
}

// Devolvemos estos 2 objetos el mismo y un error
func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) { //(event) cuando el usuario confimró su email recibe este evento
	awsgo.InicializoAWS()
	if !ValidoParametros() { //Si no ValidoParametros que haya un false
		fmt.Println("Error en los parametros, debe enviar 'SecretName'")
		err := errors.New("error en los parametros, debe debe enviar Secret Name")
		return event, err
	}

	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {

		switch row {

		case "email":
			datos.UserEmail = att //En att viene el Email del usuario
			fmt.Println("Email: " + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("Sub: " + datos.UserUUID)
		}
	}

	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("Error al leer el Secret" + err.Error())
		return event, err
	}

	err = bd.SignUp(datos)
	return event, err
}

func ValidoParametros() bool {

	var TraeParametro bool
	_, TraeParametro = os.LookupEnv("SecretName")
	return TraeParametro
}
