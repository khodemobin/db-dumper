package storage

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"log"
)

type SftpConfig struct {
	User     string
	Password string
	Host     string
	Port     string

	UploadPath string
}

func sftpLocal(path string, fileName string, cfg *SftpConfig) error {
	config := &ssh.ClientConfig{
		User: cfg.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(cfg.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", cfg.Host+":"+cfg.Port, config)
	client, err := sftp.NewClient(conn)
	if err != nil {
		return err
	}

	// walk a directory
	w := client.Walk(cfg.UploadPath)
	for w.Step() {
		if w.Err() != nil {
			continue
		}
		log.Println(w.Path())
	}

	f, err := client.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.Write([]byte("Hello world!")); err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		return err
	}

	fi, err := client.Lstat("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fi)

	if err := client.Close(); err != nil {
		return err
	}

	return nil
}
