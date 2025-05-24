package sshclient

import (
	"golang.org/x/crypto/ssh"
	"strings"
)

type SSHClient struct {
	Address  string
	User     string
	Password string
	client   *ssh.Client
}

func New(addr, user, pass string) *SSHClient {
	return &SSHClient{Address: addr, User: user, Password: pass}
}

func (c *SSHClient) Connect() error {
	config := &ssh.ClientConfig{
		User:            c.User,
		Auth:            []ssh.AuthMethod{ssh.Password(c.Password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", c.Address+":22", config)
	if err != nil {
		return err
	}
	c.client = client
	return nil
}

func (c *SSHClient) RunCommand(cmd string) error {
	session, err := c.client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	return session.Run(strings.ReplaceAll(cmd, "\n", " ; "))
}

func (c *SSHClient) Close() {
	if c.client != nil {
		c.client.Close()
	}
}
