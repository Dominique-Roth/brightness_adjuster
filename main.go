package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    // ToDo: Outsource into multiple functions
    args := os.Args[1:]
    if args[0] != "up" && args[0] != "down" {
        panic("No brightness direction provided!")
    }

    // ToDo: Read brightness file from config file
    body, err := os.ReadFile("/sys/class/backlight/acpi_video0/brightness")
    if err != nil {
        panic(err)
    }
    input_brightness := strings.TrimSuffix(string(body), "\n")
    current_brightness, err := strconv.Atoi(input_brightness)
    check(err)

    fmt.Printf("Current brightness %d\n", current_brightness)
    new_brightness := current_brightness
    switch args[0] {
        case "up":
            new_brightness = new_brightness + 1
        case "down":
            new_brightness = new_brightness - 1
    }
    
    // ToDo: Read max brightness from file
    if new_brightness > 15 || new_brightness < 0 {
        panic("Brightness not supported!")
    }

    fmt.Printf("Writing new brightness %d\n", new_brightness)
    text_to_write_to_file := []byte(strconv.Itoa(new_brightness))
    // ToDo: Read brightness file from config file
    err = os.WriteFile("/sys/class/backlight/acpi_video0/brightness", text_to_write_to_file, 644)
    check(err)
}

func check(err error) {
    if err != nil {
        panic(err)
    }
}


