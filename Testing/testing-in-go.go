package sqrt

import (
	"math"
	"testing"
)

func almostSameValue(a float64, b float64) bool {
	return math.Abs(a-b) <= 0.0001
}

func Test_Sample1(t *testing.T) {
	val, err := Sqrt(5)

	if err != nil {
		t.Fatalf("Error in calculation: %s", err)
	}
	if !almostSameValue(val, 2.2360679775) {
		t.Fatalf("Error: Bad Value %f", val)
	}
}


.................. left



cmds: go test and go test -v