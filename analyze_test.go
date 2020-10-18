package gocyclo_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/fzipp/gocyclo"
)

func TestAnalyze(t *testing.T) {
	tests := []struct {
		paths []string
		want  string
	}{
		{
			[]string{"testdata/ifs.go"},
			`3 testdata f3nested testdata/ifs.go:24:1
3 testdata f3 testdata/ifs.go:17:1
2 testdata f2else testdata/ifs.go:11:1
2 testdata f2 testdata/ifs.go:6:1
1 testdata f1 testdata/ifs.go:3:1`,
		},
		{
			[]string{"testdata/loops.go"},
			`4 testdata l4 testdata/loops.go:19:1
3 testdata l3 testdata/loops.go:8:1
2 testdata l2range testdata/loops.go:14:1
2 testdata l2 testdata/loops.go:3:1`,
		},
		{
			[]string{"testdata/cases.go"},
			`3 testdata c3default testdata/cases.go:32:1
3 testdata c3 testdata/cases.go:25:1
3 testdata c3nested testdata/cases.go:40:1
2 testdata c2multi testdata/cases.go:19:1
2 testdata c2default testdata/cases.go:12:1
2 testdata c2 testdata/cases.go:6:1
1 testdata c1 testdata/cases.go:3:1`,
		},
		{
			[]string{"testdata/comms.go"},
			`3 testdata comm3nested testdata/comms.go:33:1
3 testdata comm3default testdata/comms.go:25:1
3 testdata comm3 testdata/comms.go:18:1
2 testdata comm2default testdata/comms.go:11:1
2 testdata comm2 testdata/comms.go:5:1`,
		},
		{
			[]string{"testdata/methods.go"},
			`2 testdata (*S).m2ptr testdata/methods.go:16:1
2 testdata (S).m2 testdata/methods.go:8:1
1 testdata (*S).m1ptr testdata/methods.go:13:1
1 testdata (S).m1 testdata/methods.go:5:1`,
		},
		{
			[]string{"testdata/literals.go"},
			`3 testdata lit3 testdata/literals.go:13:12
2 testdata lit2 testdata/literals.go:8:12
1 testdata lit1 testdata/literals.go:5:12`,
		},
		{
			[]string{"testdata/ignores.go"},
			`1 testdata notIgnoredNotADirective testdata/ignores.go:13:1
1 testdata notIgnoredUnknownDirective testdata/ignores.go:10:1`,
		},
		{
			[]string{"testdata/operators.go"},
			`3 testdata op3mixed testdata/operators.go:11:1
2 testdata op2and testdata/operators.go:7:1
2 testdata op2or testdata/operators.go:3:1`,
		},
		{
			[]string{"testdata/directory"},
			`1 directory b testdata/directory/file2.go:3:1
1 directory a testdata/directory/file1.go:3:1`,
		},
	}

	for _, tt := range tests {
		stats := gocyclo.Analyze(tt.paths, nil).
			SortAndFilter(-1, 0)
		statLines := make([]string, len(stats))
		for i, s := range stats {
			statLines[i] = s.String()
		}
		got := strings.Join(statLines, "\n")
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Analyzed %q and got:\n%s\n\twant:\n%s", tt.paths, got, tt.want)
		}
	}
}
