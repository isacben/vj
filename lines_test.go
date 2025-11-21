package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateVisibleLines(t *testing.T) {
	tests := []struct {
		name     string
        firstLine int
		total    int
        content []line
		expected []line
	}{
		{
			"all lines visible",
            0,
			10,
            []line{{0, "line0"}, {1, "line1"}, {2, "line2"}},
            []line{{0, "line0"}, {1, "line1"}, {2, "line2"}},
		},
		{
			"two lines visible",
            1,
			2,
            []line{{0, "line0"}, {1, "line1"}, {2, "line2"}},
            []line{{1, "line1"}, {2, "line2"}},
		},
		{
			"four lines visible",
            2,
			4,
            []line{
                {0, "line0"}, {1, "line1"}, {2, "line2"},
                {3, "line3"}, {4, "line4"}, {5, "line5"},
                {6, "line6"}, {7, "line7"}, {8, "line7"},
            },
            []line{
                {2, "line2"}, {3, "line3"}, {4, "line4"},
                {5, "line5"},
            },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vl := NewVisibleLines(tt.firstLine, tt.total, tt.content)
			assert.Equal(t, tt.expected, vl.linesOnScreen)
		})
	}
}
