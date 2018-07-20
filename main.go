package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

// #include <unistd.h>
import "C"

const (
	listenSocket = ":6666"

	stressBinary          = "stress"
	stressTimeout         = 60
	stressIterations      = 5
	stressDefaultMemoryMB = 256
	stressMemoryHangSec   = 2
)

// Common variables.
var (
	description = "Simple stress checker for Kubernetes nodes."
	gitCommit   = "n/a"
	name        = "kube-stresscheck"
	source      = "https://github.com/giantswarm/kube-stresscheck"
)

func main() {
	// Print version.
	if (len(os.Args) > 1) && (os.Args[1] == "version") {
		fmt.Printf("Description:    %s\n", description)
		fmt.Printf("Git Commit:     %s\n", gitCommit)
		fmt.Printf("Go Version:     %s\n", runtime.Version())
		fmt.Printf("Name:           %s\n", name)
		fmt.Printf("OS / Arch:      %s / %s\n", runtime.GOOS, runtime.GOARCH)
		fmt.Printf("Source:         %s\n", source)
		return
	}

	// Print usage.
	if (len(os.Args) > 1) && (os.Args[1] == "--help") {
		return
	}

	// CPU forks to start. By default allocate 2 times cores.
	var stressCPUForks = 2 * runtime.NumCPU()

	// The easiest way w/o tons of code and external modules to get system memory.
	var totalSystemMemoryMB = C.sysconf(C._SC_PHYS_PAGES) * C.sysconf(C._SC_PAGE_SIZE) / 1024 / 1024

	// Memory forks to start. By default aim to allocate total memory size.
	var stressMemoryForks = totalSystemMemoryMB / stressDefaultMemoryMB

	args := []string{
		"--cpu", fmt.Sprintf("%v", stressCPUForks),
		"--vm", fmt.Sprintf("%v", stressMemoryForks),
		"--vm-hang", fmt.Sprintf("%v", stressMemoryHangSec),
		"--timeout", fmt.Sprintf("%v", stressTimeout),
	}

	// Invoke stress command multiple times.
	for i := 0; i < stressIterations; i++ {
		log.Printf("Executing %s with %v", stressBinary, args)

		err := exec.Command(stressBinary, args...).Run()
		if err != nil {
			log.Printf("Command stress exited with: %s", err)
		}
	}

	// We can not say was it success or failure, we just gave a stress to the node.
	log.Printf("Stress testing executed.")
}
