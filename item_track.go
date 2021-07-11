package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func loadItemTrack() []item {
	data, err := ioutil.ReadFile("track.csv")
	check(err, "Failed to read track file")

	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	var its []item
	for _, line := range lines[1:] {
		con := strings.Split(line, ",")
		fmt.Println(con)
		target, err := strconv.ParseFloat(con[2], 32)
		check(err, "Failed to parse target price")

		it := item{
			email:       strings.TrimSpace(con[0]),
			url:         strings.TrimSpace(con[1]),
			targetPrice: float32(target),
		}

		it.validate()

		its = append(its, it)
	}

	return its
}

func (it item) validate() {
	assert(strings.Contains(it.email, "@"), "Invalid email")
	assert(strings.Contains(it.url, "https://www.amazon."), "Invalid URL")
	assert(it.targetPrice > 0, "Invalid target price")
}
