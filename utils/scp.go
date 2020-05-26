package utils

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

type SCPClient struct {
	config ssh.ClientConfig
	client *ssh.Client
	sftpClient *sftp.Client
}

func NewSCPClient() (SCPClient, error) {
	// Load ENVs
	err := godotenv.Load("../.env")
	if err != nil {
		err = godotenv.Load(".env")
		if err != nil {
			return SCPClient{}, err
		}
	}

	host := os.Getenv("SSH_HOST")
	user := os.Getenv("SSH_USER")
	password := os.Getenv("SSH_PASSWORD")

	// Create connection configuration
	config := ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			sshAgent(),
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Instantiate Client
	client, err := ssh.Dial("tcp", host+":22", &config)
	if err != nil {
		return SCPClient{}, err
	}

	// Instantiate SFTP Client
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		return SCPClient{}, err
	}

	// Instantiate Config
	return SCPClient{
		config: config,
		client: client,
		sftpClient: sftpClient,
	}, err
}

// UploadFile uses scp to upload a file
// session is the session to use to upload
// srcImage is the source image
// destImagePath is the path without the name for the desired location of the image
func (s *SCPClient) UploadFile(srcImage []byte, destImagePath string) error {

	// Create file
	f, err := s.sftpClient.Create(destImagePath)
	if err != nil {
		log.Fatal(err)
	}

	// Write to file
	if _, err := f.Write(srcImage); err != nil {
		log.Fatal(err)
	}

	// Check for file
	if _, err := s.sftpClient.Lstat(destImagePath); os.IsNotExist(err) {
		return fmt.Errorf("failed to copy, no such file or directory: %s", destImagePath)
	}

	return err
}

// DeleteFile deletes a file at a given path
func (s *SCPClient) DeleteFile(filePath string) error {
	err := s.sftpClient.Remove(filePath)
	if err != nil {
		return err
	}

	return err
}

// CloseClients closes the sftp + client connections
func (s *SCPClient) CloseClients() error {
	err := s.sftpClient.Close()
	if err != nil {
		return err
	}
	err = s.client.Close()
	return err
}


func sshAgent() ssh.AuthMethod {
	if sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		return ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers)
	}
	return nil
}