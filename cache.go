package main

import (
	"fmt"
	"os/exec"
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

func (c *SongCache) CacheFull() bool {
	return c.DiskUsage() > c.MaxDiskUsage()
}
