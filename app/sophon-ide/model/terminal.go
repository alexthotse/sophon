package model

import (
	"os/exec"
)

type TerminalSession struct {
	Name string
}

func (s *TerminalSession) StartTmux() error {
	cmd := exec.Command("tmux", "new-session", "-A", "-s", "sophon")
	return cmd.Start()
}

func (s *TerminalSession) StartZellij() error {
	cmd := exec.Command("zellij", "attach", "-c", "sophon")
	return cmd.Start()
}
