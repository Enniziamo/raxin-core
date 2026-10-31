package core

import (
	"fmt"
	"os"
	"time"
)

type InteractionManager struct {
	World     *WorldManager
	Inventory *InventoryManager
}

func NewInteractionManager(world *WorldManager, inventory *InventoryManager) *InteractionManager {
	return &InteractionManager{
		World:     world,
		Inventory: inventory,
	}
}

func (m *InteractionManager) Interact(currentLocation string, foundItem *string) {
	switch currentLocation {
	case "entrance":
		m.InteractEntrance()
	case "riverside":
		m.InteractRiverside()
	case "beach":
		m.InteractBeach(foundItem)
	case "cavern":
		m.InteractCavern()
	case "church":
		m.InteractChurch()
	default:
		fmt.Println("nothing special to interact with here.")
	}
}

func (m *InteractionManager) InteractEntrance() {
	if m.Inventory.Has("stick") {
		fmt.Println("you already took a stick.")
		return
	}

	fmt.Println("you pull a sturdy stick from a fallen branch.")
	m.Inventory.AddItem("stick")
}

func (m *InteractionManager) InteractRiverside() {
	if !m.Inventory.Has("fishingrod") {
		fmt.Println("you see something shiny in the water, but it is too far to reach by hand.")
		return
	}

	if m.Inventory.Has("gem") {
		fmt.Println("the river is calm now.")
		return
	}

	fmt.Println("something shines under the water.")
	fmt.Println("you cast your fishing rod and pull out a glowing gem.")
	m.Inventory.AddItem("gem")
}

func (m *InteractionManager) InteractBeach(foundItem *string) {
	if !m.Inventory.Has("hammer") {
		fmt.Println("you see something shiny buried in the sand. maybe you can 'take hammer'.")
		*foundItem = "hammer"
	} else {
		fmt.Println("you already took the hammer.")
	}

	if !m.Inventory.Has("rope") {
		fmt.Println("you find an old rope tangled in the sand.")
		m.Inventory.AddItem("rope")
	} else {
		fmt.Println("the rope is already taken.")
	}
}

func (m *InteractionManager) InteractCavern() {
	if !m.Inventory.Has("hammer") {
		fmt.Println("you need something strong to break the stone.")
		return
	}

	fmt.Println("you break the glowing stone. a passage opens to a hidden church.")
	m.World.AddConnection("cavern", "church")
}

func (m *InteractionManager) InteractChurch() {
	if !m.Inventory.Has("nail") {
		fmt.Println("you notice a bent nail on the floor and pick it up.")
		m.Inventory.AddItem("nail")
	} else {
		fmt.Println("there is nothing else of interest here.")
	}

	if !m.Inventory.Has("gem") {
		fmt.Println("priest: 'have you found my lost gem?'")
		return
	}

	fmt.Println("priest: 'my gem. thank you, traveler.'")
	time.Sleep(1 * time.Second)
	fmt.Println("'fool. you freed me!' the priest transforms into a witcher.")
	fmt.Println("the church trembles as darkness spreads. demo ends here.")
	fmt.Println("----------------------------------------------------------")
	os.Exit(0)
}
