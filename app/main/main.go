package main

import (
	connDB "db/conn"
	cfg "firesimple/config"
	eng "firesimple/engine"
	"fmt"
	"os"
	svr "server"
)

func main() {
	sockerDB := connDB.SocketConfigDB{
		USER:     os.Getenv("FIRESIMPLE_DEFAULT_USER_DB"),
		PASSWORD: os.Getenv("FIRESIMPLE_DEFAULT_PASSWORD_DB"),
		HOST:     "localhost",
		DBNAME:   "postgres",
	}
	querysDB, err := connDB.Connect(sockerDB)
	if err != nil {
		fmt.Printf("\nERRO NA CONEXAO DO BANCO DE DADOS: %s", err.Error())
		return
	}
	fmt.Printf("\nCONEXAO FOI UM SUCESSO")
	err = querysDB.InitTabelsDB()
	if err != nil {
		fmt.Printf("\nERRO NA EXECUCAO DA QUERY DE CRIACAO DE TABELAS: %s", err.Error())
		return
	}
	fmt.Printf("\nINICIANDO FSP CONFIG")
	fspConfig := cfg.Default(querysDB)
	fmt.Printf("\nINICIANDO FSP ENGINE")
	fspEngine := eng.New(fspConfig, querysDB)
	fmt.Printf("\nINICIANDO FSP SERVICES")
	fspServices := fspEngine.Start()
	fmt.Printf("\nINICIANDO SERVER")
	server := svr.New(fspServices)
	server.Run()
}
