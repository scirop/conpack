// conpack/main.go

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	runtime     string
	packageName string
)

func init() {
	flag.StringVar(&packageName, "p", "", "Package name to search for")
	flag.StringVar(&packageName, "package", "", "Package name to search for")
	flag.StringVar(&runtime, "r", "docker", "Container runtime to use (e.g., docker, podman, finch)")
	flag.StringVar(&runtime, "runtime", "docker", "Container runtime to use (e.g., docker, podman, finch)")
	flag.Usage = usage
}

func usage() {
	fmt.Println("Usage: \n  conpack [-p|--package <package_name>] [-r|--runtime <runtime>]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -p, --package <package_name>  Specify package name to search for")
	fmt.Println("  -r, --runtime <runtime>	Specify package manager to use (e.g., docker, podman, finch). Default is docker.")
	fmt.Println("      --help                    Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  conpack -p curl")
	fmt.Println("  conpack -p curl -r podman")
	fmt.Println()
}

func main() {
	flag.Parse()

	if packageName == "" {
		usage()
		os.Exit(1)
	}

	cmd := exec.Command(runtime, "ps", "-q")
	containersBytes, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("No running containers found")
		log.Fatal(err)
	}
	containers := strings.Fields(string(containersBytes))
	if err != nil {
		log.Fatal(err)
	}
	found := false
	// Colorful rotating dots
	animations := []string{
		"⠋",
		"⠙",
		"⠹",
		"⠸",
		"⠼",
		"⠴",
		"⠲",
		"⠳",
		"⠊",
	}
	colors := []string{
		"\x1b[31m", // Red
		"\x1b[32m", // Green
		"\x1b[33m", // Yellow
		"\x1b[34m", // Blue
		"\x1b[35m", // Magenta
		"\x1b[36m", // Cyan
	}

	fmt.Printf("Checking %d containers for package %s...\n", len(containers), packageName)

	var foundContainers = []string{}

	for i, container := range containers {
		// Print wait indicator animation
		color := colors[i%len(colors)]
		animation := animations[i%len(animations)]
		fmt.Printf("\rChecking containers... [%s%s%s] (%d/%d)", color, animation, "\x1b[0m", i+1, len(containers))
		cmd := exec.Command(runtime, "exec", container, packageName, "--version")
		output, _ := cmd.CombinedOutput()
		if !strings.Contains(string(output), "exec failed") {
			foundContainers = append(foundContainers, string(container))
		}
		if len(containers) < 20 {
			time.Sleep(50 * time.Millisecond)
		}
	}
	if len(foundContainers) > 0 {
		fmt.Printf("\n\nPackage \x1b[35m%s\x1b[0m found in:\n\n", packageName)
		fmt.Printf("CONTAINER NAME\tCONTAINER ID\n")
		for _, container := range foundContainers {
			found = true
			cmd := exec.Command(runtime, "inspect", "--format='{{.Name}}'", container)
			containerName, _ := cmd.CombinedOutput()
			// remove line breaks in container name
			containerName = containerName[1 : len(containerName)-2]
			fmt.Printf("%s\t%s\n", string(containerName), container)
		}
	}
	// Clear wait indicator animation
	fmt.Println()

	if !found {
		fmt.Printf("Package %s not found on any running container\n", packageName)
	}
}
