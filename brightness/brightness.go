package brightness 

import (
    "os"
    "fmt"
    "strings"
    "strconv"
    
    "brightness_adjuster/utils"
    "brightness_adjuster/config"
)

func Increase_brightness(config config.ConfigStruct) {
    adjust_brightness("up", config)
}

func Decrease_brightness(config config.ConfigStruct) {
    adjust_brightness("down", config)
}

func adjust_brightness(
    direction string,
    config config.ConfigStruct,
) {
    body, err := os.ReadFile(config.Brightness_file)
    utils.Check(err)

    input_brightness := strings.TrimSuffix(string(body), "\n")
    current_brightness, err := strconv.Atoi(input_brightness)
    utils.Check(err)

    fmt.Printf("Current brightness %d\n", current_brightness)
    new_brightness := current_brightness
    switch direction {
        case "up":
            new_brightness = new_brightness + config.Step_count
        case "down":
            new_brightness = new_brightness - config.Step_count
    }
    max_brightness := get_max_brightness(config.Max_brightness_file)
    
    // ToDo: Read max brightness from file
    if new_brightness > max_brightness {
        new_brightness = max_brightness
    } else if new_brightness < 0 {
        // ToDo: Add options to prevent 0 and default to 1
        new_brightness = 0
    }

    fmt.Printf("Writing new brightness %d\n", new_brightness)
    text_to_write_to_file := []byte(strconv.Itoa(new_brightness))
    // ToDo: Read brightness file from config file
    err = os.WriteFile(config.Brightness_file, text_to_write_to_file, 644)
    utils.Check(err)
}

func get_max_brightness(max_brightness_file string) int {
    content, err := os.ReadFile(max_brightness_file)
    utils.Check(err)
    stripped_content := strings.TrimSuffix(string(content), "\n")
    max_brightness, err := strconv.Atoi(stripped_content)
    utils.Check(err)
    return max_brightness
}

