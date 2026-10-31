package core

import "fmt"

type InventoryManager struct {
	items []string
}

func NewInventoryManager() *InventoryManager {
	return &InventoryManager{
		items: []string{},
	}
}

func (i *InventoryManager) Has(item string) bool {
	for _, current := range i.items {
		if current == item {
			return true
		}
	}

	return false
}

func (i *InventoryManager) HasAll(items []string) bool {
	for _, item := range items {
		if !i.Has(item) {
			return false
		}
	}

	return true
}

func (i *InventoryManager) AddItem(item string) {
	if i.Has(item) {
		return
	}

	i.items = append(i.items, item)
}

func (i *InventoryManager) RemoveItem(item string) bool {
	for index, current := range i.items {
		if current == item {
			i.items = append(i.items[:index], i.items[index+1:]...)
			return true
		}
	}

	return false
}

func (i *InventoryManager) Items() []string {
	items := make([]string, len(i.items))
	copy(items, i.items)
	return items
}

func (i *InventoryManager) IsEmpty() bool {
	return len(i.items) == 0
}

func (i *InventoryManager) Print() {
	if i.IsEmpty() {
		fmt.Println("your inventory is empty.")
		return
	}

	fmt.Println("inventory:", i.items)
}
