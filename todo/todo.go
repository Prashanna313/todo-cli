package todo

import (
	"encoding/json"
	"os"
	"strconv"
)

type Item struct {
	Done     bool
	Priority int
	Text     string
	position int
}

func ReadItems(filename string) ([]Item, error) {
	var items []Item
	b, err := os.ReadFile(filename)

	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, nil
	}

	for i := range items {
		items[i].position = i + 1
	}

	return items, err
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	// fmt.Println(string(b))

	return nil
}

func (i *Item) Label() string {
	return strconv.Itoa(i.position) + "."
}

func (i *Item) PrettyDone() string {
	if i.Done {
		return "X"
	}
	return " "
}

func (i *Item) PrettyP() string {
	if i.Priority == 1 {
		return "P1"
	}
	if i.Priority == 2 {
		return "P2"
	}
	if i.Priority == 3 {
		return "P3"
	}

	return " "
}

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 2:
		i.Priority = 2
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}
