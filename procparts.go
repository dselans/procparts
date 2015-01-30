package procparts

import (
	"io/ioutil"
	"strings"
)

type Partition struct {
	Major  string
	Minor  string
	Blocks string
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

		if len(fields) == 0 {
			continue
		}

		partition := Partition{
			Major:  fields[0],
			Minor:  fields[1],
			Blocks: fields[2],
			Name:   fields[3],
		}

		partitions = append(partitions, partition)
	}

	return partitions, nil
}
