package main

import (
	controllerOrganizer "Backend/controller"
	pluginOrganizer "Backend/plugins"
	serviceOrganizer "Backend/services"
	"fmt"
)

func main() {
	controllerOrganizer.Testing()
	serviceOrganizer.Testing()
	pluginOrganizer.Testing()
	fmt.Println("vamo")
}
