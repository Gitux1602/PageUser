package tools

import (
	"fmt"
	"time"
)

func FechaMySQL() string {

	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", //Muestra en pantalla pero su salida es un string
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

}
