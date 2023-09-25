package zipdb

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func LoadLocations(file string) (locations map[string]Location, err error) {
	locations = make(map[string]Location)

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		l, err := parse(s.Text())
		if err != nil {
			return nil, err
		}
		locations[l.Zip] = l
	}

	return
}

func parse(s string) (Location, error) {
	parts := strings.Split(s, "\t")
	loc := Location{
		Country:   parts[0],
		Zip:       parts[1],
		City:      parts[2],
		StateLong: parts[3],
		State:     parts[4],
		County:    parts[5],
	}

	if l, err := strconv.ParseFloat(parts[9], 64); err != nil {
		return Location{}, err
	} else {
		loc.Lat = l
	}

	if l, err := strconv.ParseFloat(parts[10], 64); err != nil {
		return Location{}, err
	} else {
		loc.Long = l
	}

	return loc, nil
}
