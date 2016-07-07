package sim

import (
	"encoding/json"

	"github.com/go-gl/mathgl/mgl32"
)

// State represents the state of an organism.
type State struct {
	Type string `json:"type"`
	// attacking / defending / consuming
	Target uint32 `json:"target"`
	// seeking / fleeing
	Position mgl32.Vec3 `json:"position"`
}

// Attributes represents the attributes of an organism.
type Attributes struct {
	Family         uint32  `json:"family"`
	Hunger         float32 `json:"hunger"`
	Energy         float32 `json:"energy"`
	Offense        uint32  `json:"offense"`
	Defense        uint32  `json:"defense"`
	Agility        uint32  `json:"agility"`
	Range          float32 `json:"range"`
	Reproductivity uint32  `json:"reproductivity"`
}

// Organism represents a single autonomous organism.
type Organism struct {
	ID         uint32      `json:"id"`
	Position   mgl32.Vec3  `json:"position"`
	Rotation   float32     `json:"rotation"`
	State      *State      `json:"state"`
	Attributes *Attributes `json:"attributes"`
}

// Marshal returns the byte representation of an organism.
func (o *Organism) Marshal() ([]byte, error) {
	return json.Marshal(o)
}
