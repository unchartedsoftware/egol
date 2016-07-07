package sim

import (
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"
)

// Iterate applies one iteration of AI and returns the change in state as
// a map of changes.
func Iterate(organisms map[string]*Organism) map[string]*Update {
	updates := make(map[string]*Update)
	for _, organism := range organisms {
		updates[organism.ID] = &Update{
			ID: organism.ID,
			State: &State{
				Position: RandomPosition(),
			},
		}
	}

	for key, organism := range organisms {
		updates[organism.ID].State.Type = determineNextStateType(key, organism, organisms)
	}
	return updates
}

// RandomPosition returns a random vec3
func RandomPosition() mgl32.Vec3 {
	return mgl32.Vec3{
		rand.Float32(),
		rand.Float32(),
		rand.Float32(),
	}
}

func determineNextStateType(key string, organismOfInterest *Organism, organisms map[string]*Organism) string {
	positionOfInterest := organismOfInterest.State.Position

	if organismOfInterest.State.Type == "dead" {
		return "dead"
	}

	for iterKey, organism := range organisms {
		if iterKey == key {
			continue
		}
		closeBy := positionOfInterest.ApproxEqualThreshold(organism.State.Position, 0.6)
		if closeBy {
			if organismOfInterest.Attributes.Energy < organism.Attributes.Energy {
				return "dead"
			}
		}
	}

	return "alive"
}
