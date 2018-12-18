package store

import (
	"fmt"
	"testing"

	"github.com/edouardparis/spark/resources"
)

func TestListCharge(t *testing.T) {
	objects := make(map[string]*resources.Charge)
	for i := int64(0); i < 5; i++ {
		id := fmt.Sprintf("ch_%d", i+1)
		objects[id] = &resources.Charge{
			ID:      id,
			Created: i,
		}
	}

	tests := []struct {
		page     int
		size     int
		expected []string
	}{
		{
			page:     0,
			size:     10,
			expected: []string{"ch_5", "ch_4", "ch_3", "ch_2", "ch_1"},
		},
		{
			page:     0,
			size:     5,
			expected: []string{"ch_5", "ch_4", "ch_3", "ch_2", "ch_1"},
		},
		{
			page:     0,
			size:     4,
			expected: []string{"ch_5", "ch_4", "ch_3", "ch_2"},
		},
		{
			page:     1,
			size:     2,
			expected: []string{"ch_3", "ch_2"},
		},
		{
			page:     1,
			size:     3,
			expected: []string{"ch_2", "ch_1"},
		},
		{
			page:     2,
			size:     3,
			expected: []string{},
		},
	}

	for i := range tests {
		res := listCharges(tests[i].page, tests[i].size, objects)
		if len(res) != len(tests[i].expected) {
			t.Errorf("test #%d failed, length of res was not expected", i)
		}

		for j := range tests[i].expected {
			if tests[i].expected[j] != res[j].ID {
				t.Errorf("test #%d failed: wrong charge order", i)
				break
			}
		}
	}
}
