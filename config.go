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

type Config struct {
	// read from file
	Port         string
	SongDir      string
	GoogleApiKey string

	// generated at startup
	YoutubeDl string
	Version   string
}

var (
	_template string = "port={port}\nsong_dir={song_dir}\ng_api_key=\n"
	_splitter        = "="

	_defaultPort    string = "6969"
	_defaultSongdir string = "tunes"

	_config Config
)

func InitConfig() {
	_config = Config{
		port(),
		songDir(),
		gApiKey(),
		youtubeDlPath(),
		Version(),
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
	return read("g_api_key")
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
		if line == "" || !strings.Contains(line, _splitter) {
			continue
		}

		values := strings.Split(line, _splitter)
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
	path := configPath()

	f, err := os.Open(path)
	if err != nil {
		_createConfig()
		log.Printf("created default config at %s\n", path)
	}

	f, err = os.Open(path)
	if err != nil {
		panic(err)
	}

	return f
}

func _createConfig() {
	f, err := os.Create(configPath())
	if err != nil {
		panic(err)
	}

	template := _template
	template = strings.Replace(template, "{port}", _defaultPort, 1)
	template = strings.Replace(template, "{song_dir}", _defaultSongdir, 1)

	f.WriteString(template)
}
