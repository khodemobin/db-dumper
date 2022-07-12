package storage

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
)

type SftpConfig struct {
	User     string
	Password string
	Host     string
	Port     string

	UploadPath string
}

func sftpLocal(path string, fileName string, cfg *SftpConfig) error {
	client, err := sftpConnect(cfg)
	if err != nil {
		return err
	}

	file, err := client.Create(cfg.UploadPath + "/" + fileName)
	if err != nil {
		return err
	}

	input, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if _, err := file.Write(input); err != nil {
		return err
	}

	if err := file.Close(); err != nil {
		return err
	}

	if err := client.Close(); err != nil {
		return err
	}

	return nil
}

func sftpConnect(cfg *SftpConfig) (*sftp.Client, error) {
	config := &ssh.ClientConfig{
		User: cfg.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(cfg.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", cfg.Host+":"+cfg.Port, config)
	if err != nil {
		return nil, err
	}

	return sftp.NewClient(conn)
}
