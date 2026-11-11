package netboxapi

import "testing"

// TestListFilter: repeatable filters keep every value, single-value filters reject a second one,
// values that do not parse are reported with the filter's name, and a filter without values leaves
// its setter alone.
func TestListFilter(t *testing.T) {
	var ids []int64
	filter := FilterInt64s(func(v []int64) { ids = v })
	for _, value := range []string{"1", "2"} {
		if problem := filter.Add("id", value); problem != nil {
			t.Fatalf("Add(%q): %+v", value, problem)
		}
	}
	filter.Apply()
	if len(ids) != 2 || ids[0] != 1 || ids[1] != 2 {
		t.Errorf("ids = %v, want [1 2]", ids)
	}

	problem := FilterInt64s(func([]int64) {}).Add("id", "x")
	if problem == nil || problem.Summary != "Invalid filter value" || problem.Detail != `Filter 'id' takes an integer, got "x".` {
		t.Errorf("non-integer: %+v", problem)
	}

	var empty *bool
	single := FilterBool(func(v *bool) { empty = v })
	if problem := single.Add("tag__empty", "true"); problem != nil {
		t.Fatalf("first value: %+v", problem)
	}
	if problem := single.Add("tag__empty", "false"); problem == nil || problem.Summary != "Duplicate filter" {
		t.Errorf("second value: %+v", problem)
	}
	single.Apply()
	if empty == nil || !*empty {
		t.Errorf("empty = %v, want true", empty)
	}

	called := false
	FilterStrings(func([]string) { called = true }).Apply()
	if called {
		t.Error("a filter without values called its setter")
	}
}
