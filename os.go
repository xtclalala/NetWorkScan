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

func NewOS(osStr string) iLinuxOS {
	osStr = strings.ToLower(osStr)

	for osName, commands := range OSConfig {
		if strings.IndexAny(osStr, strings.ToUpper(osName)) != -1 {
			return &LinuxOS{
				name:     osName,
				commands: commands,
			}
		}
	}
	return &LinuxOS{
		name:     osStr,
		commands: OSConfig["base"],
	}
}
