package main

import (
	config "authentication-system/config"
	server "authentication-system/server"
	"log"
)

func main() {
	// Inicializando a conexão Database
	err := config.Load()
	if err != nil {
		log.Fatalln(err) //panic(err)
	}

	// Inicia um Servidor do Tipo GIN
	// Server := server.NewServer()

	// Chama a função RUN e Inicia o serviço do Server
	// Server.Run()
	server.Server()
}
