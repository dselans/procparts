package procparts

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Partition struct {
	Major  uint64
	Minor  uint64
	Blocks uint64
	Name   string
}

func GetPartitions() ([]Partition, error) {
	partitionsFile := "/proc/partitions"

	data, err := ioutil.ReadFile(partitionsFile)
	if err != nil {
		return nil, err
	}

	partitions := make([]Partition, 0)

	for _, line := range strings.Split(string(data), "\n") {
		fields := strings.Fields(line)

		if len(fields) == 0 || fields[0] == "major" {
			continue
		}

		major, err := strconv.ParseUint(fields[0], 0, 64)
		if err != nil {
			return nil, err
		}

		minor, err := strconv.ParseUint(fields[1], 0, 64)
		if err != nil {
			return nil, err
		}

		blocks, err := strconv.ParseUint(fields[2], 0, 64)
		if err != nil {
			return nil, err
		}

		partition := Partition{
			Major:  major,
			Minor:  minor,
			Blocks: blocks,
			Name:   fields[3],
		}

		partitions = append(partitions, partition)
	}

	return partitions, nil
}
