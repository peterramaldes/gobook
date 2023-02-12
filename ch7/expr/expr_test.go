package expr_test

import (
	"fmt"
	"math"
	"testing"

	"expr"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  expr.Env
		want string
	}{
		{"sqrt(A / pi)", expr.Env{"A": 87616, "pi": math.Pi}, "167"},
	}
	var prevExpr string
	for _, tt := range tests {
		// Just show the expr when change
		if tt.expr != prevExpr {
			fmt.Printf("\n%s\n", tt.expr)
			prevExpr = tt.expr
		}

		expr, err := Parse(tt.expr)
		if err != nil {
			t.Error(err) // err on parse
			continue
		}

		got := fmt.Sprintf("%.6g", expr.Eval(tt.env))
		fmt.Printf("\t%v => %s\n", tt.env, got)
		if got != tt.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				tt.expr, tt.env, got, tt.want)
		}

	}
}
