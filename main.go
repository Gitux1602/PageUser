package main

import (
	"Gambit/Users/PageUser/awsgo"
	"context"
	"errors"
	"fmt"
	"os"

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
}

func ValidoParametros() bool {

	var TraeParametro bool
	_, TraeParametro = os.LookupEnv("SecretName")
	return TraeParametro
}
