package storage

import (
	"fmt"
	"os"
	"time"

	"github.com/secsy/goftp"
)

type FtpConfig struct {
	Host     string
	User     string
	Port     string
	Password string

	UploadPath string
}

func ftpUpload(path, fileName string, cfg *FtpConfig) error {
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	config := goftp.Config{
		User:               cfg.User,
		Password:           cfg.Password,
		ConnectionsPerHost: 10,
		Timeout:            10 * time.Second,
		Logger:             os.Stderr,
	}

	client, err := goftp.DialConfig(config, addr)
	if err != nil {
		return err
	}

	backup, err := os.Open(path)
	if err != nil {
		return err
	}

	return client.Store(cfg.UploadPath+fileName, backup)
}
