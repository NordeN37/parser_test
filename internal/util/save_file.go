package util

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Exists reports whether the named file or directory exists.
func Exists(path, name string) error {
	if _, err := os.Stat(path + name); err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(path, 0755); err != nil {
				return err
			}
		}
	}
	return nil
}

func CreateFileResultJson(result interface{}, path, fileName string) error {
	resultJson, err := json.Marshal(result)
	if err != nil {
		return err
	}
	err = Exists(path, fileName+".json")
	if err != nil {
		return err
	}
	err = os.WriteFile(fmt.Sprintf("%s%s_%d.%s", path, fileName, time.Now().Unix(), "json"), resultJson, 0644)
	if err != nil {
		return err
	}
	return nil
}
