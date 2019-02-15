package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type SongCache struct{}

var Cache *SongCache = &SongCache{}

func (c *SongCache) DiskUsage() uint64 {
	cmd := exec.Command(
		Config.Du(),
		"-b",
		"-0",
		Config.Read(CFG_SONG_DIR))

	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	s := fmt.Sprintf("%s", out)
	s = strings.Split(s, "\t")[0]

	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return i
}

func (c *SongCache) MaxDiskUsage() uint64 {
	size := Config.Read(CFG_CACHE_SIZE)

	// first figure out if a suffix is included
	suffix := size[len(size)-1:]
	var str string

	if suffix != "M" && suffix != "G" {
		// suffix doesn't exist, assume G
		suffix = "G"
		str = size
	} else {
		str = size[0 : len(size)-1]
	}

	n, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		panic(err)
	}

	var multiplier uint64
	switch suffix {
	case "M":
		multiplier = 1000000
		break
	case "G":
		fallthrough
	default:
		multiplier = 1000000000
		break
	}

	return n * multiplier
}

func (c *SongCache) Full() bool {
	return c.DiskUsage() > c.MaxDiskUsage()
}

type ByModTime []os.FileInfo

func (m ByModTime) Len() int           { return len(m) }
func (m ByModTime) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m ByModTime) Less(i, j int) bool { return m[i].ModTime().Unix() < m[j].ModTime().Unix() }

func (c *SongCache) Prune() {
	var currentUsage int64 = 0
	configUsage := c.MaxDiskUsage()
	songDir := Config.Read(CFG_SONG_DIR)

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	path := filepath.Join(cwd, songDir)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		currentUsage += f.Size()
	}

	// order songs by mod time, oldest -> youngest
	sort.Sort(ByModTime(files))

	for i, f := range files {
		log.Printf("%d\t%s\t%s\n", i, f.Name(), f.ModTime())
	}

	log.Printf("current usage: %d\n", currentUsage)
	log.Printf("config usage: %d\n", configUsage)

	for pruning := c.Full(); pruning; pruning = c.Full() {
		fPath := filepath.Join(path, files[0].Name())
		log.Printf("deleting: %s\n", fPath)
		err = os.Remove(fPath)
		if err != nil {
			panic(err)
		}

		// pop front item off, it's deleted now
		files = files[1:]
	}

	log.Printf("current usage: %d\n", c.DiskUsage())
}

func (c *SongCache) Update(s *Song) {
	cmd := exec.Command(
		Config.Touch(),
		s.Path())

	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	log.Printf("update song mod time: %s\n", s.Filename())
}
