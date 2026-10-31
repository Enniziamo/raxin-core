package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"raxin/core"
)

type Game struct {
	World       *core.WorldManager
	Inventory   *core.InventoryManager
	Interaction *core.InteractionManager

	Current   string
	History   []string
	FoundItem string
}

func NewGame() *Game {
	world := core.NewWorldManager()
	inventory := core.NewInventoryManager()

	return &Game{
		World:       world,
		Inventory:   inventory,
		Interaction: core.NewInteractionManager(world, inventory),
		Current:     "entrance",
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	game := NewGame()

	fmt.Println("welcome to the adventure!")
	fmt.Println("type 'help' to see available commands.\n")

	for {
		game.RenderCurrentLocation()

		fmt.Print("\n> ")
		if !scanner.Scan() {
			break
		}

		input := strings.ToLower(strings.TrimSpace(scanner.Text()))
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		command := parts[0]
		args := parts[1:]

		switch command {
		case "help":
			printHelp()
		case "look":
			game.Look()
		case "go":
			game.Move(args)
		case "back":
			game.GoBack()
		case "interact":
			game.Interact()
		case "take":
			game.Take(args)
		case "open":
			game.Open()
		case "talk":
			game.Talk()
		case "craft":
			game.Craft(args)
		case "inventory":
			game.Inventory.Print()
		case "map":
			fmt.Println(game.World.MapText())
		case "exit":
			fmt.Println("you exit the game. goodbye!")
			os.Exit(0)
		default:
			fmt.Println("unknown command. type 'help' to see commands.")
		}
	}
}

func printHelp() {
	fmt.Println("\ncommands:")
	fmt.Println("  look          - look around")
	fmt.Println("  go <place>    - move to a connected place (example: go city)")
	fmt.Println("  back          - return to the previous location")
	fmt.Println("  interact      - interact with the environment")
	fmt.Println("  take <item>   - take an item found in the current location")
	fmt.Println("  open          - open chests or containers")
	fmt.Println("  talk          - talk to people")
	fmt.Println("  craft <item>  - craft an item (example: craft fishingrod)")
	fmt.Println("  inventory     - show your items")
	fmt.Println("  map           - show the map")
	fmt.Println("  exit          - quit the game")
}

func (g *Game) RenderCurrentLocation() {
	location, ok := g.World.GetLocation(g.Current)
	if !ok {
		fmt.Println("unknown location.")
		return
	}

	fmt.Println(location.ASCIIArt)
	fmt.Printf("location: %s\n%s\n", location.Name, location.Description)
}

func (g *Game) Look() {
	location, ok := g.World.GetLocation(g.Current)
	if !ok {
		fmt.Println("unknown location.")
		return
	}

	fmt.Println(location.ASCIIArt)
	fmt.Printf("location: %s\n%s\n", location.Name, location.Description)
	fmt.Printf("you can go to: %s\n", g.World.ConnectionText(g.Current))
}

func (g *Game) Move(args []string) {
	if len(args) < 1 {
		fmt.Println("where do you want to go? (example: go city)")
		return
	}

	target := args[0]

	if target == g.Current {
		fmt.Println("you are already there.")
		return
	}

	if !g.World.CanMove(g.Current, target) {
		fmt.Println("you can't go there directly.")
		return
	}

	g.History = append(g.History, g.Current)
	g.Current = target
	g.FoundItem = ""
}

func (g *Game) GoBack() {
	if len(g.History) == 0 {
		fmt.Println("there is nowhere to go back.")
		return
	}

	g.Current = g.History[len(g.History)-1]
	g.History = g.History[:len(g.History)-1]
	g.FoundItem = ""
}

func (g *Game) Interact() {
	g.Interaction.Interact(g.Current, &g.FoundItem)
}

func (g *Game) Take(args []string) {
	if len(args) < 1 {
		fmt.Println("take what?")
		return
	}

	item := args[0]

	if g.FoundItem != item {
		fmt.Println("there is nothing like that to take here.")
		return
	}

	g.Inventory.AddItem(item)
	g.FoundItem = ""
	fmt.Printf("you picked up the %s!\n", item)
}

func (g *Game) Open() {
	if g.Current != "city" {
		fmt.Println("there is nothing to open here.")
		return
	}

	if g.Inventory.Has("hatenna") {
		fmt.Println("the chest is empty now.")
		return
	}

	g.Inventory.AddItem("hatenna")
	fmt.Println("you open the chest and find a pokemon. hatenna joins you!")
}

func (g *Game) Talk() {
	switch g.Current {
	case "city":
		fmt.Println("old man: 'i heard there is a tunnel under the glowing stone in the cavern.'")
	case "church":
		if g.Inventory.Has("gem") {
			fmt.Println("priest: 'my gem! thank you, traveler!'")
			fmt.Println("you are blessed with a divine light.")
			g.Inventory.AddItem("blessing")
		} else {
			fmt.Println("priest: 'i lost my precious gem somewhere in the forest.'")
		}
	default:
		fmt.Println("no one to talk to here.")
	}
}

func (g *Game) Craft(args []string) {
	if len(args) < 1 {
		fmt.Println("craft what? (example: craft fishingrod)")
		return
	}

	if args[0] != "fishingrod" {
		fmt.Println("you can't craft that.")
		return
	}

	required := []string{"stick", "rope", "nail"}

	if !g.Inventory.HasAll(required) {
		fmt.Println("you don't have all materials. you need: stick, rope, nail.")
		return
	}

	if g.Inventory.Has("fishingrod") {
		fmt.Println("you already made a fishing rod.")
		return
	}

	g.Inventory.AddItem("fishingrod")
	fmt.Println("you tie rope to the stick and use the nail as a hook. a fishing rod is made!")
}
