package config

import (
    "os"
    "path/filepath"

    "brightness_adjuster/utils"

    "github.com/pelletier/go-toml/v2"
)

type ConfigStruct struct {
    Brightness_file string
    Max_brightness_file string
    Step_count int
}


func GetBrightnessConfiguration() ConfigStruct {
    config := ConfigStruct{}
    config_file_path := "config.toml"
    err := toml.Unmarshal(getConfigFile(config_file_path), &config)
    utils.Check(err)
    return config
}

func getConfigFile(file_path string) []byte {
    executable, err := os.Executable()
    utils.Check(err)
    executable_path := filepath.Dir(executable)
    content, err := os.ReadFile(executable_path + "/" + file_path)
    utils.Check(err)
    return content
}
