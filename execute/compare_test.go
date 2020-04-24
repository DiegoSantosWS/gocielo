package execute_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/DiegoSantosWS/gocielo/typescielo"

	// "github.com/andreyvit/diff"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type DiffReporter struct {
	path  cmp.Path
	diffs []string
}

func (r *DiffReporter) PushStep(ps cmp.PathStep) {
	r.path = append(r.path, ps)
}

func (r *DiffReporter) Report(rs cmp.Result) {
	if !rs.Equal() {
		vx, vy := r.path.Last().Values()
		r.diffs = append(r.diffs, fmt.Sprintf("%#v:\n\t-: %+v\n\t+: %+v\n", r.path, vx, vy))
	}
}

func (r *DiffReporter) PopStep() {
	r.path = r.path[:len(r.path)-1]
}

func (r *DiffReporter) String() string {
	return strings.Join(r.diffs, "\n")
}

func validatePayment(t *testing.T, name string, got, exp *typescielo.Payment) {
	diff := cmp.Diff(exp, got, cmpopts.IgnoreFields(typescielo.Payment{}, "ExtraDataCollection", "Links"))
	if len(diff) > 0 {
		t.Errorf("test [%s] mismatch (-want +got):\n%s", name, diff)
	}
}

func validateGetCreditCard(t *testing.T, name string, got typescielo.CreditCard, exp typescielo.CreditCard) {
	r := DiffReporter{}
	if !cmp.Equal(exp, got, cmp.Reporter(&r)) {
		t.Errorf("test [%s] mismatch (-want +got):\n%s", name, r.String())
	}
}

func validateAddCreditCard(t *testing.T, name, exp map[string]interface{}) {
	// TODO: CRIAR TEST
}
