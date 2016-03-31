package main

import (
	"fmt"
	"runtime"

	"SparkyEngineGO/gfx"
	"SparkyEngineGO/input"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	// creates glfw window and inits OpenGL
	window := gfx.CreateWindow(960, 540, "Sparky")

	// Hook mouse and key events
	window.SetMouseButtonCallback(onMouseBtn)
	window.SetKeyCallback(input.OnKey)
	window.SetCloseCallback(onClose)

	gl.ClearColor(0.2, 0.3, 0.8, 1.0)
	gl.ClearDepth(1)
	// fmt.Printf("version_: %d", *gl.GetString(gl.VERSION))

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	for !window.ShouldClose() {
		// DO OpenGL Stuff
		gfx.Clear()

		gl.DrawArrays(gl.BUFFER, 0, 6)
		// old school gl drawing for testing
		gl.Begin(gl.QUADS)
		gl.Vertex2f(-0.5, 0.5)
		gl.Vertex2f(0.5, 0.5)
		gl.Vertex2f(0.5, -0.5)
		gl.Vertex2f(-0.5, -0.5)
		gl.End()

		gfx.Update(window)
	}
}

func onKey(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {
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

func onMouseBtn(window *glfw.Window, b glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Press {
		return
	}
	switch glfw.MouseButton(b) {
	case glfw.MouseButtonLeft:
		fmt.Println("You pressed the left mouse button")
	default:
		return

	}
}

func onClose(window *glfw.Window) {
	window.SetShouldClose(true)
}
