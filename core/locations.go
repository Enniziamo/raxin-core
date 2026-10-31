package core

type Location struct {
	Name        string
	Description string
	Connections []string
}

func BuildLocations() map[string]Location {
	return map[string]Location{
		"entrance": {
			Name:        "entrance",
			Description: "there is a stick. you stand at the forest entrance. paths lead forward into the woods.",
			Connections: []string{"riverside", "cavern"},
		},
		"riverside": {
			Name:        "riverside",
			Description: "you reach the riverside. the sound of water calms you. maybe you can fish here.",
			Connections: []string{"entrance", "city", "riverbank"},
		},
		"city": {
			Name:        "city",
			Description: "you arrive at a noisy city. there is a chest near a fountain. an old man is here.",
			Connections: []string{"riverside", "beach", "cavern"},
		},
		"beach": {
			Name:        "beach",
			Description: "you stand on a sunny beach. the sand hides something shiny.",
			Connections: []string{"city"},
		},
		"cavern": {
			Name:        "cavern",
			Description: "you enter a dark cavern. a glowing stone blocks your way.",
			Connections: []string{"riverside"},
		},
		"church": {
			Name:        "church",
			Description: "a secret church beyond the cavern. a priest stands here, looking sad.",
			Connections: []string{"cavern"},
		},
		"riverbank": {
			Name:        "riverbank",
			Description: "you stand at the edge of a sparkling river. a brown raccoon is here, sipping on a soda. it looks at you curiously.",
			Connections: []string{"riverside"},
		},
	}
}
