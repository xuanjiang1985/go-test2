package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
)

// Monster struct
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

// Store is a method of Monster
func (m *Monster) Store() error {
	data, err := json.Marshal(m)

	if err != nil {
		return errors.WithStack(err)
	}

	ioutil.WriteFile("monstor.str", data, 0666)
	return nil
}

// Restore data from file
func (m *Monster) Restore() error {
	byte1, err := ioutil.ReadFile("monstor.str1")

	if err != nil {
		return errors.WithStack(err)
	}

	err = json.Unmarshal(byte1, m)

	if err != nil {
		return errors.WithStack(err)
	}
	fmt.Printf("%#v\n", m)
	return nil
}
