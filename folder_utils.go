package main

import "os"

func CreateFolder(name string) error {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		if err := os.Mkdir(name, os.ModeSticky|os.ModePerm); err != nil {
			println("Failed to create", name, "folder", err.Error())
			return err
		}
	}
	return nil
}
