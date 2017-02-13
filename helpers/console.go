package helpers

import (
	"os"
	"os/exec"
	"runtime"
)

// ClearConsole ..
func ClearConsole() {
	switch runtime.GOOS {
	case "linux":
	case "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		break
	case "windows":
		cmd := exec.Command("clr")
		cmd.Stdout = os.Stdout
		cmd.Run()
		break
	default:
		panic("unsupported platform for clear")
	}
}
