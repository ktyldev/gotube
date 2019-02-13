package main

import (
	"fmt"
	"log"
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
