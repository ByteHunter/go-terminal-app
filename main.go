package main

import (
	"bytes"
	"fmt"
)

var terminal Terminal
var canvas Canvas
var chronos Chronos

func main() {
	// Configure and ensure the configuration of the terminal is restored
	terminal = *NewTerminal()
	terminal.configure()
	defer terminal.restore()

	// Initialize some required stuff
	chronos = *NewChronos()
	canvas = *NewCanvas(terminal.get_size())
	channel := make(chan []byte)

	// Launch input goroutine
	go input_loop(channel)

	// First drawing
	chronos.begin()
	simple_draw([]byte{})
	chronos.measure()

	// Begin the main loop here
	main_loop(channel)

	// Clear the canvas before exiting
	canvas.reset()
	canvas.clear_screen()
	canvas.pos(1, 1)
}

// This goroutine reads the input
func input_loop(buffer chan<- []byte) {
	for {
		r := terminal.read_raw()
		if len(r) > 0 {
			buffer <- r
		}
	}
}

// This is the main loop for the program to react to input
func main_loop(channel <-chan []byte) {
	for {
		r := <-channel

		if bytes.Equal(r, K_ESC) {
			break
		}

		if bytes.Equal(r, K_CTRL_C) {
			canvas.set_error("Ctrl+C is disabled in this program use 'ESC' instead")
		}

		if bytes.Equal(r, KEY_a) {
			canvas.draw_line_horizontal(1, 1, 0, ' ')
		}

		// An example of drawing and measuring performance
		chronos.begin()
		simple_draw(r)
		chronos.measure()
	}
}

func simple_draw(r []byte) {
	canvas.clear_screen()
	canvas.background(Color{50, 50, 50})
	canvas.draw_line_horizontal(1, canvas.height, canvas.width, ' ')
	canvas.pos(1, canvas.height)
	canvas.color(Color{240, 240, 240})
	// fmt.Printf("Received an input: %s (%.2f ms) FPS: %.2f", r, chronos.ms, chronos.fps)
	fmt.Printf("Received an input: (%s) %+v (%.2f ms) FPS: %.2f", r, r, chronos.ms, chronos.fps)
	if canvas.last_error != nil {
		canvas.background(Color{200, 0, 0})
		canvas.color(Color{250, 250, 250})
		canvas.draw_line_horizontal(canvas.width/2, canvas.height, canvas.width-canvas.width/2+1, ' ')
		canvas.pos(canvas.width/2+1, canvas.height)
		fmt.Printf(" " + canvas.get_last_error().Error())
	}
	canvas.reset()
}
