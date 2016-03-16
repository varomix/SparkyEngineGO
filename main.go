package main

import (
	"SparkyEngineGO/gfx"
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
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

	gl.ClearColor(0.2, 0.3, 0.8, 1.0)
	gl.ClearDepth(1)
	fmt.Printf("version_: %d", *gl.GetString(gl.VERSION))

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	for !window.ShouldClose() {
		// DO OpenGL Stuff
		gfx.Clear()

		gl.DrawArrays(gl.BUFFER, 0, 6)

		gfx.Update(window)
	}

}
