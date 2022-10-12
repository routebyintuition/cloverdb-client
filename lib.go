package main

import "os"

func dirValid(path string) (bool, error) {
	loc, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, nil
	}

	if !loc.IsDir() {
		return false, nil
	}

	return true, nil
}
