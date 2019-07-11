package square

import (
	"encoding/json"
	"strings"
)

type Circle struct {
	Cord
	Radius json.Number `json:"radius"`
}

type Square struct {
	LeftBottom Cord `json:"left_bottom"`
	RightTop   Cord `json:"right_top"`
}

type Cord struct {
	X json.Number `json:"x"`
	Y json.Number `json:"y"`
}

// IsFloat if there is at least one coordinate with a point return true
func IsFloat(task []Circle) bool {
	for _, c := range task {
		if strings.Contains(c.Radius.String(), ".") || strings.Contains(c.Y.String(), ".") || strings.Contains(c.X.String(), ".") {
			return true
		}
	}
	return false
}
