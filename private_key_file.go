package sshtunnel

import (
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

func ParsePrivateKeyFile(file string, passphrase []byte) (ssh.AuthMethod, error) {

	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return ParsePrivateKey(buffer, passphrase)

}

func ParsePrivateKey(buffer []byte, passphrase []byte) (ssh.AuthMethod, error) {

	var signer ssh.Signer
	var err error

	if passphrase != nil {
		signer, err = ssh.ParsePrivateKeyWithPassphrase(buffer, passphrase)
	} else {
		signer, err = ssh.ParsePrivateKey(buffer)
	}

	if err != nil {
		return nil, err
	}

	return ssh.PublicKeys(signer), nil

}
