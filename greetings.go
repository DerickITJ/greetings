package greetings

import (
	"math/rand"
	"errors"
	"fmt"
)

func Hello(name string) (string,error) {
	if name == ""{
		return "",errors.New("Nombre Vacio")
	}

	message := fmt.Sprintf(ramdomFormat(),name)
	return message,nil
}

func Hellos(names []string)(map[string]string,error){
	messages := make(map[string]string)
	for _,name := range names{
		message,err := Hello(name)
		if err != nil{
			return nil,err
		}
		messages[name]=message
	}
	return messages,nil
}

func ramdomFormat()string{
	formats:=[]string{
		"¡Hola! %v ¡Bienvenido!",
		"Que bueno verte, %v",
		"¡Saludo, %v! ¡Encantado de Conocerte!",
	}
	return formats[rand.Intn(len(formats))]
}

