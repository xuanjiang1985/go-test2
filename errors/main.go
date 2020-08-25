package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func main() {
	err := readFile("../config/config.yml1")

	if err != nil {
		fmt.Printf("%+v", err)
	}

	// s, err := ioutil.ReadAll(r)
	// if err != nil {
	// 	fmt.Printf("%s\n", errors.Wrap(err, "read failed"))
	// }
	// fmt.Printf("%s\n", string(s))
}

func readFile(file string) error {
	_, err := os.OpenFile(file, 0, os.ModeAppend)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
