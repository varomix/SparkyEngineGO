package gfx

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

// CreateWindow creates the window
func CreateWindow(width, height int, title string) *glfw.Window {
	win, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}
	win.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	win.SetSizeCallback(windowResize)

	return win
}

// Clear clears the window
func Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// Update s the window
func Update(window *glfw.Window) {
	window.SwapBuffers()
	glfw.PollEvents()
}

func windowResize(win *glfw.Window, width, height int) {
	// get the window size
	x, y := win.GetSize()
	// resize the gl Viewport when resizing the window
	gl.Viewport(0, 0, int32(x), int32(y))

}
