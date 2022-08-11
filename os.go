package main

import "strings"

type iLinuxOS interface {
	osString() string
}

type LinuxOS struct {
	name string
}

func (s *LinuxOS) osString() string {
	return s.name
}

func NewOS(osStr string) iLinuxOS {
	osStr = strings.ToLower(osStr)
	return &LinuxOS{
		name: osStr,
	}
}
