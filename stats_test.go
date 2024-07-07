package gocyclo_test

import (
	"reflect"
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

func TestSortAndFilter(t *testing.T) {
	tests := []struct {
		stats gocyclo.Stats
		top   int
		over  int
		want  gocyclo.Stats
	}{
		{
			stats: gocyclo.Stats{
				{Complexity: 1},
				{Complexity: 4},
				{Complexity: 2},
				{Complexity: 3},
			},
			top: -1, over: 0,
			want: gocyclo.Stats{
				{Complexity: 4},
				{Complexity: 3},
				{Complexity: 2},
				{Complexity: 1},
			},
		},
		{
			stats: gocyclo.Stats{
				{Complexity: 1},
				{Complexity: 2},
				{Complexity: 3},
				{Complexity: 4},
			},
			top: 2, over: 0,
			want: gocyclo.Stats{
				{Complexity: 4},
				{Complexity: 3},
			},
		},
		{
			stats: gocyclo.Stats{
				{Complexity: 1},
				{Complexity: 2},
				{Complexity: 4},
				{Complexity: 4},
				{Complexity: 5},
			},
			top: -1, over: 3,
			want: gocyclo.Stats{
				{Complexity: 5},
				{Complexity: 4},
				{Complexity: 4},
			},
		},
		{
			stats: gocyclo.Stats{
				{Complexity: 1},
				{Complexity: 2},
				{Complexity: 3},
				{Complexity: 4},
				{Complexity: 5},
			},
			top: 2, over: 2,
			want: gocyclo.Stats{
				{Complexity: 5},
				{Complexity: 4},
			},
		},
	}
	for _, tt := range tests {
		got := tt.stats.SortAndFilter(tt.top, tt.over)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Sort and filter (top %d over %d) for %q, got: %q, want: %q",
				tt.top, tt.over, tt.stats, got, tt.want)
		}
	}
}

func TestPercentile(t *testing.T) {
	tests := []struct {
		stats   gocyclo.Stats
		k       int
		want    int
		wantErr bool
	}{
		{
			// unsorted stats
			stats: gocyclo.Stats{
				{Complexity: 2},
				{Complexity: 4},
				{Complexity: 3},
				{Complexity: 1},
			},
			k:       -1,
			want:    -1,
			wantErr: true,
		},
		{
			// out-of-range k
			stats: gocyclo.Stats{
				{Complexity: 4},
				{Complexity: 3},
				{Complexity: 2},
				{Complexity: 1},
			},
			k:       -1,
			want:    -1,
			wantErr: true,
		},
		{
			// out-of-range k
			stats: gocyclo.Stats{
				{Complexity: 4},
				{Complexity: 3},
				{Complexity: 2},
				{Complexity: 1},
			},
			k:       100,
			want:    -1,
			wantErr: true,
		},
		{
			stats: gocyclo.Stats{
				{Complexity: 4},
				{Complexity: 3},
				{Complexity: 2},
				{Complexity: 1},
			},
			k:       50,
			want:    3,
			wantErr: false,
		},
		{
			stats: gocyclo.Stats{
				{Complexity: 4},
				{Complexity: 3},
				{Complexity: 2},
				{Complexity: 1},
			},
			k:       90,
			want:    4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := tt.stats.Percentile(tt.k)
		if (err != nil) != tt.wantErr {
			t.Errorf("expect Percentile error state %t got %v", tt.wantErr, err)
		}
		if tt.want != got {
			t.Errorf("expect Percentile returns %d got %d", tt.want, got)
		}
	}
}
