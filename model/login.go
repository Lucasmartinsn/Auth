package model

import (
	tipo "authentication-system/data_type"
	db "authentication-system/db"
	services "authentication-system/services"
)

func Verificalogin(user string, senha string) (login tipo.Login, err error) {
	conn, err := db.ConectionDB()
	if err != nil {
		return
	}
	defer conn.Close()

	if senha == "root" && user == "root" {
		row := conn.QueryRow(`SELECT id FROM login WHERE username=$1 and password=$2`, user, senha)
		err = row.Scan(&login.Id)

		return

	} else {
		senha = services.Sha256Encoder(senha)
		row := conn.QueryRow(`SELECT id FROM login WHERE username=$1 and password=$2`, user, senha)
		err = row.Scan(&login.Id)

		return
	}
}
