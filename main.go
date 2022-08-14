package main

import (
	"fmt"
	"path/filepath"
	"os"

	"github.com/NeowayLabs/drm"
)

const (
	driPath = "/dev/dri"
	baseRender = 128
)


func main() {
	for i, v := range drm.ListDevices() {
		renderPath := filepath.Join(driPath, fmt.Sprintf("renderD%d", baseRender+i))
		if _, err := os.Stat(renderPath); err != nil {
			continue
		}

		linkPath := filepath.Join(driPath, fmt.Sprintf("renderN%s", v.Name))
		if _, err := os.Stat(linkPath); err != nil {
			fmt.Printf("Skipping linking %s to %s as that filename already exists\n", renderPath, linkPath)
			continue
		}

		err := os.Link(renderPath, linkPath)
		if err != nil {
			fmt.Printf("Failed to link %s to %s: %v\n", renderPath, linkPath, err)
		}
	}
}
