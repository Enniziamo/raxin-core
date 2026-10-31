package core

import "strings"

type WorldManager struct {
	locations map[string]Location
}

func NewWorldManager() *WorldManager {
	return &WorldManager{
		locations: BuildLocations(),
	}
}

func (w *WorldManager) GetLocation(name string) (Location, bool) {
	location, ok := w.locations[name]
	return location, ok
}

func (w *WorldManager) CanMove(current, target string) bool {
	location, ok := w.locations[current]
	if !ok {
		return false
	}

	for _, connection := range location.Connections {
		if connection == target {
			return true
		}
	}

	return false
}

func (w *WorldManager) AddConnection(from, to string) {
	location, ok := w.locations[from]
	if !ok {
		return
	}

	for _, connection := range location.Connections {
		if connection == to {
			return
		}
	}

	location.Connections = append(location.Connections, to)
	w.locations[from] = location
}

func (w *WorldManager) MapText() string {
	return `map:

entrance
cavern
riverside
riverbank
city
beach`
}

func (w *WorldManager) ConnectionText(locationName string) string {
	location, ok := w.locations[locationName]
	if !ok {
		return ""
	}

	return strings.Join(location.Connections, ", ")
}
