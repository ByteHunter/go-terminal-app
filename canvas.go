package main

import (
	"errors"
	"fmt"
	"strings"
)

var T_ESC = "\x1B"
var T_CSI = "\x1B\x5B"

type Color struct {
	r, g, b int
}

type Vector struct {
	x, y int
}

type Canvas struct {
	width, height int
	last_error    error
}

func NewCanvas(width, height int) *Canvas {
	return &Canvas{
		width:  width,
		height: height,
	}
}

func (canvas *Canvas) pos(x, y int) {
	fmt.Printf("\x1B\x5B%d;%dH", y, x)
}

func (canvas *Canvas) color(color Color) {
	fmt.Printf("\x1B\x5B38;2;%d;%d;%dm", color.r, color.g, color.b)
}

func (canvas *Canvas) background(color Color) {
	fmt.Printf("\x1B\x5B48;2;%d;%d;%dm", color.r, color.g, color.b)
}

func (canvas *Canvas) reset() {
	fmt.Printf("\x1B\x5B0m")
}

func (canvas *Canvas) clear_screen() {
	fmt.Printf("\x1B\x5B2J")
}

func (canvas *Canvas) draw_line_horizontal(x, y, l int, r rune) {
	canvas.pos(x, y)
	if l <= 0 {
		canvas.set_error(fmt.Sprintf("Line length must be bigger than 0, got(%d)", l))
	}
	fmt.Printf(strings.Repeat(string(r), l))
}

/*
 * Error messages
 */
func (canvas *Canvas) set_error(message string) {
	canvas.last_error = errors.New(message)
}

func (canvas *Canvas) get_last_error() error {
	defer canvas.clear_errors()
	return canvas.last_error
}

func (canvas *Canvas) clear_errors() {
	canvas.last_error = nil
}
