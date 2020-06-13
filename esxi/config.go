package esxi

import (
	"fmt"
	"log"
)

type Config struct {
	esxiHostName string
	esxiHostPort string
	esxiUserName string
	esxiPassword string
}

func (c *Config) validateEsxiCreds() error {
	log.Printf("[validateEsxiCreds]\n")

	var remote_cmd string
	var err error

	remote_cmd = fmt.Sprintf("vmware --version")
	_, err = runRemoteSshCommand(c, remote_cmd, "Connectivity test, get vmware version")
	if err != nil {
		return fmt.Errorf("Failed to connect to esxi host: %s\n", err)
	}
	return nil
}
