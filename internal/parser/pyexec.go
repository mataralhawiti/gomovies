package parser

import "os/exec"

func RunPyScript(pythonBinPth string) {
	cmd := exec.Command(pythonBinPth, "py/test.py")

	err := cmd.Start()
	if err != nil {
		panic(err)
	}
}
