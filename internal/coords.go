package internal

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Provider interface {
	GetCoordinates(ctx context.Context) (Coordinates, error)
}

type CoreLocationCLI struct {
	binary string
}

func (p CoreLocationCLI) GetCoordinates() (Coordinates, error) {
	//CoreLocationCLI gives us our coordinates on MacOS
	bin := p.binary
	if bin == "" {
		bin = "CoreLocationCLI"
	}

	out, err := exec.Command(bin, "--format", "%latitude %longitude").Output()
	if err != nil {
		return Coordinates{}, err
	}
	parts := strings.Fields(strings.TrimSpace(string(out)))
	if len(parts) != 2 {
		return Coordinates{}, fmt.Errorf("unexpected coordinates output: %w", string(out))
	}

	latitude, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return Coordinates{}, err
	}

	longitude, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return Coordinates{}, err
	}
	return Coordinates{latitude: latitude, longitude: longitude}, nil
}
