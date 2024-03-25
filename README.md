# Saludos en GO 


## Instalación
Ejecuta el siguiente comando para instalar el paquete:

```bash
go get -u  github.com/felipmora49/greetings
```
## USO 
Aqui tienes un ejemplo de cómo utlizar el paquete en tu código:

```go

package main

import (
	
	"fmt"
	"log"
	"github.com/felipmora49/greetings"
)

func main(){

	log.SetPrefix("greetings: ")
	log.SetFlags(0) 

	names := []string{"Alex", "Julian", "Sara", "Diego"}
	messages, err := greetings.Hellos(names)

// 	message , err := greetings.Hello()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)

}

```
Este es un ejemplo ultizando el paquete "github.com/felipmora49/greetings"