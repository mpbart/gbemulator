package main

import (
	"fmt"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type display struct {
	window *glfw.Window
}

type Display interface {
	Render()
	Start()
}

func CreateDisplay() {
	err := glfw.Init()
	if err != nil {
		fmt.Println(err)
	}
	defer glfw.Terminate()

	if err := gl.Init(); err != nil {
		fmt.Println(err)
	}

	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
    if err != nil {
        fmt.Println(err)
    }

	window.MakeContextCurrent()

	gl.Viewport(0, 0, 640, 480)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(640), float64(480), 0, -1, 1)
	gl.ClearColor(0.255, 0.255, 0.255, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	window.SetPos(0, 0)

	d := display{window}
	d.Start()
}

func (d *display) Render() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.Disable(gl.DEPTH_TEST)
    gl.PointSize(1.0)
    gl.Begin(gl.POINTS)
    for y := 0; y < 480; y++ {
        for x := 0; x < 640; x++ {
			if x % 4 == 0 {
				gl.Color3ub(30, 20, 10)
				gl.Vertex2i(int32(x), int32(y))
			}
        }
    }
    gl.End()
    d.window.SwapBuffers()
}

func (d *display) Start() {
	for !d.window.ShouldClose() {
		d.Render()
		d.window.SwapBuffers()
		glfw.PollEvents()
	}
}

