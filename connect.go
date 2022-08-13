package main

import (
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
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
	RunShell() (string, error)
	UploadShell() error
	ScanOS() string
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

func (s *Ssh) RunShell() (string, error) {
	if s.connectErr != nil {
		return "", s.connectErr
	}
	session, err := s.client.NewSession()
	if err != nil {
		return "", errors.Wrap(err, "Build session fail;"+s.errIdentify)
	}
	out, _ := session.CombinedOutput("sh " + global.Shell.AbsPath + "/" + global.Shell.Name)
	fmt.Printf("out: %s\n", out)
	return string(out), err
}

func (s *Ssh) UploadShell() error {
	if s.connectErr != nil {
		return s.connectErr
	}
	session, err := s.client.NewSession()
	defer session.Close()
	if err != nil {
		return errors.Wrap(err, "Build session fail;"+s.errIdentify)
	}

	filename := global.Shell.Name
	File, _ := os.Open(global.Shell.LocalPath + "/" + filename)
	dirname := global.Shell.AbsPath
	size, _ := File.Stat()
	go func() {
		w, _ := session.StdinPipe()
		fmt.Fprintln(w, "C0644", size.Size(), filename)
		io.CopyN(w, File, size.Size())
		fmt.Fprint(w, "\x00")
		w.Close()
	}()

	if err = session.Run(fmt.Sprintf("/usr/bin/scp -qrt %s/", dirname)); err != nil {
		return err
	}
	return nil
}

// ScanOS 针对操作系统执行不同的命令，并返回运行结果
func (s *Ssh) ScanOS() string {
	var res string
	// upload sh script
	_ = s.UploadShell()
	// run sh script
	res, _ = s.RunShell()
	// return response
	return res
}
