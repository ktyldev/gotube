package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const _TEMPLATE = "port=\nsong_dir=\ng_api_key=\n"
const _SPLITTER = "="

type Config struct {
	// read from file
	Port         string
	SongDir      string
	GoogleApiKey string

	// generated at startup
	YoutubeDl string
}

var _config Config

func InitConfig() {
	_config = Config{
		port(),
		songDir(),
		gApiKey(),
		youtubeDlPath(),
	}
}

func GetConfig() Config {
	return _config
}

func port() string {
	port, err := strconv.Atoi(read("port"))
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf(":%d", port)
}

func songDir() string {
	dir := read("song_dir")
	if dir == "" {
		log.Fatalln("song_dir not set in config")
	}

	return dir
}

func gApiKey() string {
	key := read("g_api_key")
	if key == "" {
		log.Println("google api key not set - using slow search")
	} else {
		log.Printf("using api key: %s\n", key)
	}

	return key
}

func youtubeDlPath() string {
	bin := "youtube-dl"

	out, err := exec.
		Command("which", bin).
		CombinedOutput()

	if err != nil {
		log.Fatalf("couldn't find %s - are you sure it's installed?\n", bin)
	}

	path := fmt.Sprintf("%s", out)
	path = strings.TrimSuffix(path, "\n")

	log.Printf("found %s at %s", bin, path)

	return path
}

func configPath() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return filepath.Join(cwd, "config.txt")
}

func read(key string) string {
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
		log.Printf("no value set for key: %s\n", key)
	}

	return result
}

func openConfig() *os.File {
	f, err := os.Open(configPath())
	if err != nil {
		_createConfig()
		log.Fatalf("please fill in config file at %s\n", configPath())
	}

	return f
}

func _createConfig() {
	f, err := os.Create(configPath())
	if err != nil {
		panic(err)
	}

	f.WriteString(_TEMPLATE)
}
