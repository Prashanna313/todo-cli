package todo

import (
	"encoding/json"
	"io/ioutil"
)

type Item struct {
	Priority int
	Text     string
}

func ReadItems(filename string) ([]Item, error) {
	var items []Item
	b, err := ioutil.ReadFile(filename)

	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, nil
	}
	return items, err
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	// fmt.Println(string(b))

	return nil
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
