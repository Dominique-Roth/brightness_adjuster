package main

import (
	"os"

    "brightness_adjuster/config"
    "brightness_adjuster/brightness"
)

func main() {
    args := os.Args[1:]
    if args[0] != "up" && args[0] != "down" {
        panic("No brightness direction provided!")
    }

    config := config.GetBrightnessConfiguration()
    switch args[0] {
        case "up":
            brightness.Increase_brightness(config)
        case "down":
            brightness.Decrease_brightness(config)
    }

}



