package main

import (
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

type Ssh struct {
	addr        string
	password    string
	user        string
	errIdentify string

	client     *ssh.Client
	connectErr error
}

type ISsh interface {
	Connect()
	RunCmd(...string) (string, error)
}

func NewSsh(ip, port, user, password string) *Ssh {
	return &Ssh{
		addr:        ip + ":" + port,
		password:    password,
		user:        user,
		errIdentify: fmt.Sprintf(" addr: %s ,user: %s", ip+":"+port, user),
	}
}

// Connect 建立SSH客户端连接
func (s *Ssh) Connect() {
	client, err := ssh.Dial("tcp", s.addr, &ssh.ClientConfig{
		User:            s.user,
		Auth:            []ssh.AuthMethod{ssh.Password(s.password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		s.connectErr = errors.Wrap(err, "connect fail;"+s.errIdentify)
		return
	}
	s.client = client
	return
}

// RunCmd 建立新会话并运行
func (s *Ssh) RunCmd(cmdList string) (string, error) {
	if s.connectErr != nil {
		return "", s.connectErr
	}
	session, err := s.client.NewSession()
	defer session.Close()
	if err != nil {
		return "", errors.Wrap(err, "Build session fail;"+s.errIdentify)
	}

	out, err := session.CombinedOutput(cmdList)

	return string(out), nil
}