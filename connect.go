package main

import (
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"time"
)

type Ssh struct {
	addr        string
	password    string
	user        string
	errIdentify string
	os          iLinuxOS
	osFlag      bool

	client     *ssh.Client
	connectErr error
}

type ISsh interface {
	Connect()
	RunCmd(string) (string, error)
	GetOS() error
	getOS(cmd string) error
	Save() []string
	ScanOS() []string
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
		Timeout:         time.Second * time.Duration(global.Connect.Timeout),
	})
	if err != nil {
		s.connectErr = errors.Wrap(err, "connect fail;"+s.errIdentify)
		return
	}
	s.client = client
	return
}

// RunCmd 建立新会话并运行
func (s *Ssh) RunCmd(cmd string) (string, error) {
	if s.connectErr != nil {
		return "", s.connectErr
	}
	session, err := s.client.NewSession()
	defer session.Close()
	if err != nil {
		return "", errors.Wrap(err, "Build session fail;"+s.errIdentify)
	}

	out, err := session.CombinedOutput(cmd)
	if err != nil {
		return "", errors.Wrap(err, "Run command fail;"+s.errIdentify)
	}
	return string(out), nil
}

// GetOS 识别 linux 发行版本
func (s *Ssh) GetOS() error {
	commands := osCommand()
	var err error
	for _, command := range commands {
		err = s.getOS(command)
	}
	return err
}

func (s *Ssh) getOS(cmd string) error {
	if s.osFlag {
		return nil
	}
	osStr, err := s.RunCmd(cmd)
	if err != nil {
		return err
	}
	s.os = NewOS(osStr)
	s.osFlag = true
	return nil
}

// Save 返回需要保存的内容
func (s *Ssh) Save() []string {
	return []string{s.addr, s.user, s.password, s.os.osString()}
}

// ScanOS 针对操作系统执行不同的命令，并返回运行结果
func (s *Ssh) ScanOS() []string {
	var res []string
	commands := s.os.getCommands()
	for _, command := range commands {
		out, err := s.RunCmd(command)
		res = append(res, command)
		if err != nil {
			res = append(res, err.Error())
		} else {
			res = append(res, out)
		}
	}
	return res
}
