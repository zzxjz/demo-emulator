package glfw

import (
	"fmt"
	"log"
	"runtime"

	"github.com/zzxjz/demo-emulator/gb/core"
	"github.com/zzxjz/demo-emulator/gb/util"
	"github.com/go-gl/glfw/v3.2/glfw"
)

// multiplier to control the size of the Window
// Width = GbWidth * zoom, Height = GbHeight * zoom
const zoom = 5

var window *glfw.Window

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

// Setup initializes the windowing system
func Setup() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	fmt.Println("glfw init ok")

	// fixed size window
	glfw.WindowHint(glfw.Resizable, glfw.False)

	// sets the OpenGL version to 2.1
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	var err error
	window, err = glfw.CreateWindow(util.GbWidth*zoom, util.GbHeight*zoom, "Hack & Sangria - GB Emu", nil, nil)
	if err != nil {
		panic(err)
	}

	// sets the callback to receive input events
	window.SetKeyCallback(onKey)

	// move the rendering context to this thread
	window.MakeContextCurrent()

	// enable vsync, forces 60fps on most displays
	glfw.SwapInterval(1)

	fmt.Println("window ok")
}

// Update swaps the buffers and updates the input
func Update() {
	window.SwapBuffers()
	glfw.PollEvents()
}

// WindowClosed checks if the window has been closed by the user
func WindowClosed() bool {
	return window.ShouldClose()
}

// Teardown shut downs the windowing system
func Teardown() {
	glfw.Terminate()
}

func onKey(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {

	// check for key press / release events
	// and notifiy the emulator accordingly

	switch glfw.Key(k) {
	case glfw.KeyEnter:
		if action == glfw.Press {
			core.ButtonPressed(util.StartButton)
		} else if action == glfw.Release {
			core.ButtonReleased(util.StartButton)
		}
	case glfw.KeySpace:
		if action == glfw.Press {
			core.ButtonPressed(util.SelectButton)
		} else if action == glfw.Release {
			core.ButtonReleased(util.SelectButton)
		}
	case glfw.KeyS:
		if action == glfw.Press {
			core.ButtonPressed(util.AButton)
		} else if action == glfw.Release {
			core.ButtonReleased(util.AButton)
		}
	case glfw.KeyA:
		if action == glfw.Press {
			core.ButtonPressed(util.BButton)
		} else if action == glfw.Release {
			core.ButtonReleased(util.BButton)
		}
	case glfw.KeyUp:
		if action == glfw.Press {
			core.ButtonPressed(util.UpButton)
		} else if action == glfw.Release {
			core.ButtonReleased(util.UpButton)
		}
	case glfw.KeyDown:
		if action == glfw.Press {
			core.ButtonPressed(util.DownButton)
		} else if action == glfw.Release {
			core.ButtonReleased(util.DownButton)
		}
	case glfw.KeyLeft:
		if action == glfw.Press {
			core.ButtonPressed(util.LeftButton)
		} else if action == glfw.Release {
			core.ButtonReleased(util.LeftButton)
		}
	case glfw.KeyRight:
		if action == glfw.Press {
			core.ButtonPressed(util.RightButton)
		} else if action == glfw.Release {
			core.ButtonReleased(util.RightButton)
		}
	default:
		return
	}
}
