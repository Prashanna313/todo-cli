package todo

import (
	"encoding/json"
	"io/ioutil"
)

type Item struct {
	Text string
}

func ReadItems(filename string) ([]Item, error) {
	err := ioutil.ReadFile(filename)

	if err != nil {
		return []Item{}, nil
	}
	return nil, err
	235
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	// fmt.Println(string(b))

	return nil
}
