package utils

import (
	"fmt"
	"os"

	"github.com/jlaffaye/ftp"
	"github.com/joho/godotenv"
)

type FTPStruct struct {
	connection *ftp.ServerConn
}

func NewFTPStruct() (FTPStruct, error) {

	// Load ENVs
	err := godotenv.Load("../.env")
	if err != nil {
		return FTPStruct{}, err
	}

	host := os.Getenv("FTP_HOST")
	username := os.Getenv("FTP_USER")
	password := os.Getenv("FTP_PASSWORD")

	// Instantiate FTP connection
	c, err := ftp.Connect(host)
	if err != nil {
		return FTPStruct{}, err
	}

	// Start FTP connection
	err = c.Login(username, password)
	if err != nil {
		return FTPStruct{}, err
	}
	fmt.Println("Logged in")

	r, err := c.Retr("/home/igivnqlrr5nm/public_html/assets/img/icon-2.png")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	fmt.Println(r)

	return FTPStruct{connection: c}, err
}
