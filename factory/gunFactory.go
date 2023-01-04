package main

import "fmt"

func getGun(gunType string) (igun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	} else if gunType == "musket" {
		return newMusket(), nil
	} else {
		return nil, fmt.Errorf("Wrong gun type passed")
	}
}
