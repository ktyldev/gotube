package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const _TEMPLATE = "port=\nsong_dir=\n"
const _SPLITTER = "="

type Config struct {
	Port    string
	SongDir string
}

var _config Config

func InitConfig() {
	_config = Config{
		getPort(),
		getSongDir(),
	}
}

func GetConfig() Config {
	return _config
}

func getPort() string {
	port, err := strconv.Atoi(_read("port"))
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf(":%d", port)
}

func getSongDir() string {
	return _read("song_dir")
}

func _configPath() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return filepath.Join(cwd, "config.txt")
}

func _read(key string) string {
	var result string = ""

	f := openConfig()
	defer f.Close()

	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || !strings.Contains(line, _SPLITTER) {
			continue
		}

		values := strings.Split(line, _SPLITTER)
		if len(values) != 2 {
			log.Fatalf("malformed config entry: %s\n", line)
		}

		if values[0] == key {
			result = values[1]
		}
	}

	if result == "" {
		log.Fatalf("no value set for key: %s\n", key)
	}

	return result
}

func openConfig() *os.File {
	f, err := os.Open(_configPath())
	if err != nil {
		_createConfig()
		log.Fatalf("please fill in config file at %s\n", _configPath())
	}

	return f
}

func _createConfig() {
	f, err := os.Create(_configPath())
	if err != nil {
		panic(err)
	}

	f.WriteString(_TEMPLATE)
}
