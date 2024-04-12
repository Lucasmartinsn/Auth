package db

import (
	"database/sql"
	"fmt"
	"log"

	config "authentication-system/config"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func ConectionDB() (*sql.DB, error) {
	conf := config.GetDB()

	//pega os dados pala conectar ao banco vindos de config
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	// cria duas variaves e atribue a conecção com o banco. passando junto o nome e as conf de acesso
	conn, err := sql.Open("postgres", sc)
	if err != nil {
		log.Fatalln(err)
	}

	// se nao der erro aqui vai ser confirmada e aberta a conecção
	err = conn.Ping()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("falha ao pingar o banco de dados: %v", err)
	}

	return conn, err
}

// Aqui é o exemplificação de como você pode ativar a segunda conexao com o database do Mysql
// OBS: Uma vez que descomentar essa função você deve alterar o codigo nos demais aquivos para que seja utilizado essa conexao

// func ConectionDB() (*sql.DB, error) {
// 	_, conf := config.GetDB()

// 	//pega os dados pala conectar ao banco vindos de config
// 	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
// 		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

// 	// cria duas variaves e atribue a conecção com o banco. passando junto o nome e as conf de acesso
// 	conn, err := sql.Open("mysql", sc)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	// se nao der erro aqui vai ser confirmada e aberta a conecção
// 	err = conn.Ping()
// 	if err != nil {
// 		conn.Close()
// 		return nil, fmt.Errorf("falha ao pingar o banco de dados: %v", err)
// 	}

// 	return conn, err
// }