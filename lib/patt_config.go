package command

import (
	"os"
	"fmt"
	"bufio"
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	Source string `json:"source"`
	Destination string `json:"destination"`
}

func configFilePath() string {
	return os.Getenv("HOME") + "/.patt.d/config.json"
}

func ReadConfig() map[string]Config {
	var configs map[string]Config

	fp, err := os.Open(configFilePath())
	if err != nil {
		fmt.Println(err)
	}
	defer fp.Close()

	reader := bufio.NewReader(fp)
	jsonBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(jsonBytes, &configs)
	if err != nil {
		fmt.Println(err)
	}
	return configs
}

func WriteConfig(configs map[string]Config) error {
	err := os.Remove(configFilePath())
	if err != nil {
		return err
	}

	jsonBytes, err := json.MarshalIndent(configs, "", " ")
	if err != nil {
		return err
	}

	fp, err := os.Create(configFilePath())

	if err != nil {
		fmt.Println(err)
	}
	defer fp.Close()
	writer := bufio.NewWriter(fp)
	_, err = writer.Write(jsonBytes)
	if err != nil {
		return err
	}
	writer.Flush()

	return nil
}