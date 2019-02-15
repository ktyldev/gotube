package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Configuration struct{}
type ConfigItem struct {
	Key          string
	DefaultValue string
}

const (
	CFG_PORT       = "port"
	CFG_SONG_DIR   = "song_dir"
	CFG_CACHE_SIZE = "cache_size"
	CFG_G_API_KEY  = "g_api_key"
)

var (
	Config *Configuration = &Configuration{}

	_splitter = "="

	_defaultPort      string = "6969"
	_defaultSongDir   string = "tunes"
	_defaultCacheSize string = "100M"

	_default = []ConfigItem{
		ConfigItem{
			CFG_PORT,
			_defaultPort,
		},
		ConfigItem{
			CFG_SONG_DIR,
			_defaultSongDir,
		},
		ConfigItem{
			CFG_CACHE_SIZE,
			_defaultCacheSize,
		},
		ConfigItem{
			CFG_G_API_KEY,
			"",
		},
	}
)

func (c *Configuration) Read(key string) string {
	hasKey := false
	for _, v := range _default {
		hasKey = v.Key == key
		if hasKey {
			break
		}
	}

	if !hasKey {
		panic(fmt.Sprintf("key: %s doesn't exist in config\n", key))
	}

	var result string = ""

	f := c.Open()
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

func (c *Configuration) SongDirPath() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return filepath.Join(cwd, c.Read(CFG_SONG_DIR))
}

func (c *Configuration) YoutubeDl() string {
	return _path("youtube-dl")
}

func (c *Configuration) Ffmpeg() string {
	return _path("ffmpeg")
}

func (c *Configuration) Version() string {
	return "4.20.69"
}

func (c *Configuration) Path() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return filepath.Join(cwd, "config.txt")
}

func (c *Configuration) Exists() bool {
	_, err := os.Open(Config.Path())
	return err == nil
}

func (c *Configuration) Open() *os.File {
	path := c.Path()

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return f
}

func (c *Configuration) Create() {
	path := c.Path()
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	template := ""
	for _, v := range _default {
		template += fmt.Sprintf("%s%s%s\n", v.Key, _splitter, v.DefaultValue)
	}

	f.WriteString(template)
	log.Printf("created config at %s\n", path)
}

func _path(bin string) string {
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
