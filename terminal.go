package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var TCONF_RAW = []string{
	"-F", "/dev/tty",
	"-icanon", "-echo",
	"min", "0", "time", "0",
	"-isig", "-ixon",
}

var TCONF_NORMAL = []string{
	"-F", "/dev/tty", "icanon", "echo",
}

type Terminal struct {
}

func NewTerminal() *Terminal {
	return &Terminal{}
}

func (terminal *Terminal) configure() {
	exec.Command("stty", TCONF_RAW...).Run()
	fmt.Printf("\x1B\x5B?25l")
}

func (terminal *Terminal) restore() {
	exec.Command("stty", TCONF_NORMAL...).Run()
	fmt.Printf("\x1B\x5B?25h")
}

func (terminal *Terminal) commandOutput(name string, arg ...string) string {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	return string(out)
}

func (terminal *Terminal) get_size() (int, int) {
	output := terminal.commandOutput("stty", "size")
	output = strings.TrimSpace(output)
	size := strings.Split(output, " ")
	rows, _ := strconv.Atoi(size[0])
	cols, _ := strconv.Atoi(size[1])

	return cols, rows
}

func (terminal *Terminal) read_raw() []byte {
	var b []byte = make([]byte, 4096)
	os.Stdin.Sync()
	n, err := os.Stdin.Read(b)
	if err != nil {
		return []byte{}
	}
	b = b[:n]

	return b
}
