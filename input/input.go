package input

import (
	"fmt"

	"github.com/go-gl/glfw/v3.1/glfw"
)

// OnKey process keyboard input
func OnKey(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Press {
		return
	}
	switch glfw.Key(k) {
	case glfw.KeyEscape:
		window.SetShouldClose(true)
	case glfw.KeySpace:
		fmt.Println("You pressed space")
	default:
		return
	}
}
