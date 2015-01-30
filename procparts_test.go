package procparts

import (
	"testing"
)

func TestGetPartitions(t *testing.T) {
	partitions, err := GetPartitions()

	if err != nil {
		t.Error("Should not have any errors. Error:", err)
	}

	if len(partitions) == 0 {
		t.Error("Returned partition count should be higher than 0")
	}

	for _, part := range partitions {
		if part.Name == "" {
			t.Error("Name of the partition should not be blank")
		}
	}
}
