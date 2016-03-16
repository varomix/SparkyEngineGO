package main

import (
	"SparkyEngineGO/gfx"
	"runtime"

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

	gl.ClearColor(0.2, 0.3, 0.8, 1.0)
	gl.ClearDepth(1)

	for !window.ShouldClose() {
		// DO OpenGL Stuff
		gfx.Clear()

		// old school gl drawing for testing
		gl.Begin(gl.QUADS)
		gl.Vertex2f(-0.5, 0.5)
		gl.Vertex2f(0.5, 0.5)
		gl.Vertex2f(0.5, -0.5)
		gl.Vertex2f(-0.5, -0.5)
		gl.End()

		// window.SwapBuffers()
		// glfw.PollEvents()
		gfx.Update(window)
	}

}
