package utils

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
	"github.com/tmc/scp"
	"golang.org/x/crypto/ssh"
)

type SCPClient struct {
	config ssh.ClientConfig
	client *ssh.Client
}

func NewSCPClient() (SCPClient, error) {
	// Load ENVs
	err := godotenv.Load("../.env")
	if err != nil {
		return SCPClient{}, err
	}

	host := os.Getenv("SSH_HOST")
	user := os.Getenv("SSH_USER")
	privateKeyPath := os.Getenv("SSH_PRIVATE_KEY_PATH")

	// Read the private key file
	key, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return SCPClient{}, fmt.Errorf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return SCPClient{}, fmt.Errorf("unable to parse private key: %v", err)
	}

	// Create connection configuration
	config := ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Instantiate Client
	client, err := ssh.Dial("tcp", host+":22", &config)
	if err != nil {
		return SCPClient{}, err
	}

	// Instantiate Config
	return SCPClient{
		config: ssh.ClientConfig{},
		client: client,
	}, err
}

func (s SCPClient) closeClient() error {
	fmt.Println("b4 client close")
	err := s.client.Close()
	fmt.Println("after client close")
	return err
}

// uploadFile uses scp to upload a file
// session is the session to use to upload
// srcImagePath is the path including the name desired for the source image
// destImagePath is the path without the name for the desired location of the image
func (s SCPClient) uploadFile(session *ssh.Session, srcImagePath string, destImagePath string) error {

	// // Load the image to upload
	// file, _ := os.Open(srcImagePath)
	// defer file.Close()
	//
	// // Instantiate file stat struct
	// stat, err := file.Stat()
	// if err != nil {
	// 	return err
	// }

	// hostIn, err := session.StdinPipe()
	// if err != nil {
	// 	return err
	// }
	// defer hostIn.Close()

	// Prepares writer
	err := scp.CopyPath(srcImagePath, destImagePath, session)
	if err != nil {
		panic("Failed to Copy: " + err.Error())
	}
	if _, err := os.Stat(destImagePath); os.IsNotExist(err) {
		return fmt.Errorf("failed to copy, no such file or directory: %s", destImagePath)
	}

	return err
}
