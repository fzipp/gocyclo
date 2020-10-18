package gocyclo_test

import (
	"testing"

	"github.com/fzipp/gocyclo"
)

func TestAverageComplexity(t *testing.T) {
	tests := []struct {
		stats gocyclo.Stats
		want  float64
	}{
		{gocyclo.Stats{
			{Complexity: 2},
		}, 2},
		{gocyclo.Stats{
			{Complexity: 2},
			{Complexity: 3},
		}, 2.5},
		{gocyclo.Stats{
			{Complexity: 2},
			{Complexity: 3},
			{Complexity: 4},
		}, 3},
		{gocyclo.Stats{
			{Complexity: 2},
			{Complexity: 3},
			{Complexity: 3},
			{Complexity: 3},
		}, 2.75},
	}
	for _, tt := range tests {
		got := tt.stats.AverageComplexity()
		if got != tt.want {
			t.Errorf("Average complexity for %q, got: %g, want: %g", tt.stats, got, tt.want)
		}
	}
}

func TestTotalComplexity(t *testing.T) {
	tests := []struct {
		stats gocyclo.Stats
		want  uint64
	}{
		{gocyclo.Stats{
			{Complexity: 2},
		}, 2},
		{gocyclo.Stats{
			{Complexity: 2},
			{Complexity: 3},
		}, 5},
		{gocyclo.Stats{
			{Complexity: 2},
			{Complexity: 3},
			{Complexity: 4},
		}, 9},
		{gocyclo.Stats{
			{Complexity: 2},
			{Complexity: 3},
			{Complexity: 3},
			{Complexity: 3},
		}, 11},
	}
	for _, tt := range tests {
		got := tt.stats.TotalComplexity()
		if got != tt.want {
			t.Errorf("Total complexity for %q, got: %d, want: %d", tt.stats, got, tt.want)
		}
	}
}
