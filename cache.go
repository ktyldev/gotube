package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"
)

type SongCache struct{}

var Cache *SongCache = &SongCache{}

func (c *SongCache) DiskUsage() uint64 {
	var usage uint64 = 0

	for _, f := range c.Files() {
		usage += uint64(f.Size())
	}

	return usage
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

func (c *SongCache) Files() []os.FileInfo {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	path := filepath.Join(cwd, Config.Read(CFG_SONG_DIR))
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	return files
}

type ByModTime []os.FileInfo

func (m ByModTime) Len() int           { return len(m) }
func (m ByModTime) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m ByModTime) Less(i, j int) bool { return m[i].ModTime().Unix() < m[j].ModTime().Unix() }

func (c *SongCache) Prune() {
	var err error
	files := c.Files()
	path := Config.SongDirPath()

	// order songs by mod time, oldest -> youngest
	sort.Sort(ByModTime(files))
	// exit early, no need to prune
	if !c.Full() {
		return
	}

	c.LogUsage()
	log.Println("cache full - pruning oldest files")
	for pruning := c.Full(); pruning; pruning = c.Full() {
		name := files[0].Name()

		fPath := filepath.Join(path, name)
		log.Printf("removing: %s\n", name)
		err = os.Remove(fPath)
		if err != nil {
			panic(err)
		}

		// pop front item off, it's deleted now
		files = files[1:]
	}

	c.LogUsage()
}

func (c *SongCache) LogUsage() {
	percentage := float64(c.DiskUsage()) / float64(c.MaxDiskUsage()) * 100

	log.Printf(
		"cache usage:\t%d/%d\t(%.1f%%)\n",
		c.DiskUsage(),
		c.MaxDiskUsage(),
		percentage)
}

func (c *SongCache) Update(s *Song) {
	now := time.Now().Local()
	err := os.Chtimes(s.Path(), now, now)
	if err != nil {
		panic(err)
	}

	log.Printf("update song mod time: %s\n", s.Filename())
}
