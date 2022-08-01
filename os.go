package main

import "strings"

func osCommand() []string {
	return []string{
		"[ -f /etc/os-release ] && awk -F'[= \"]' '/PRETTY_NAME/{print $3,$4,$5}' /etc/os-release",
		"[ -f /etc/lsb-release ] && awk -F'[=\"]+' '/DESCRIPTION/{print $2}' /etc/lsb-release",
		"[ -f /etc/redhat-release ] && awk '{print $0}' /etc/redhat-release"}
}

type iLinuxOS interface {
	osString() string
	getCommands() []string
}

type LinuxOS struct {
	name     string
	commands []string
}

func (s *LinuxOS) osString() string {
	return s.name
}

func (s *LinuxOS) getCommands() []string {
	return s.commands
}

type Centos struct {
	LinuxOS
}
type Ubuntu struct {
	LinuxOS
}
type OpenEuler struct {
	LinuxOS
}

func NewOS(osStr string) iLinuxOS {
	osStr = strings.ToLower(osStr)
	switch {

	case strings.IndexAny(osStr, "openeuler") != -1:
		return &OpenEuler{LinuxOS{
			name:     osStr,
			commands: global.Os.OpenEuler,
		}}
	case strings.IndexAny(osStr, "centos") != -1:
		return &Centos{LinuxOS{
			name: osStr,
		}}
	case strings.IndexAny(osStr, "ubuntu") != -1:
		return &Ubuntu{LinuxOS{
			name: osStr,
		}}
	default:
		return &LinuxOS{
			name:     osStr,
			commands: global.Os.Base,
		}
	}

}
