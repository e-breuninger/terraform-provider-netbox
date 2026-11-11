// Code generated automatically. DO NOT EDIT.

// The filters of the plural data sources: parsing the values configured for a filter and passing
// them to the matching go-netbox list parameter.

package netboxapi

import (
	"fmt"
	"strconv"
)

// ListFilter collects the values one entry of a plural data source's filters set gives a list
// query parameter, and hands them to the parameter's go-netbox setter. The generated read maps
// each supported filter name to one ListFilter, so a filter costs one table entry, not code.
type ListFilter interface {
	// Add parses value for the filter called name and keeps it.
	Add(name, value string) *FilterValueError
	// Apply passes the kept values to the setter; a filter without values leaves it alone.
	Apply()
}

// FilterValueError is a filter value the parameter does not take: one that does not parse, or a
// second value for a single-value parameter.
type FilterValueError struct {
	Summary, Detail string
}

// FilterStrings, FilterInt64s, FilterFloat64s and FilterBools serve a repeatable parameter: every
// value is kept, and NetBox ORs them.
func FilterStrings(set func([]string)) ListFilter {
	return &repeatableFilter[string]{set, parseFilterString, nil}
}
func FilterInt64s(set func([]int64)) ListFilter {
	return &repeatableFilter[int64]{set, parseFilterInt64, nil}
}
func FilterFloat64s(set func([]float64)) ListFilter {
	return &repeatableFilter[float64]{set, parseFilterFloat64, nil}
}
func FilterBools(set func([]bool)) ListFilter {
	return &repeatableFilter[bool]{set, parseFilterBool, nil}
}

// FilterString, FilterInt64, FilterFloat64 and FilterBool serve a parameter NetBox takes once.
func FilterString(set func(*string)) ListFilter {
	return &singleFilter[string]{set, parseFilterString, nil}
}
func FilterInt64(set func(*int64)) ListFilter {
	return &singleFilter[int64]{set, parseFilterInt64, nil}
}
func FilterFloat64(set func(*float64)) ListFilter {
	return &singleFilter[float64]{set, parseFilterFloat64, nil}
}
func FilterBool(set func(*bool)) ListFilter { return &singleFilter[bool]{set, parseFilterBool, nil} }

// repeatableFilter keeps every value and passes them as one slice.
type repeatableFilter[T any] struct {
	set    func([]T)
	parse  func(name, value string) (T, *FilterValueError)
	values []T
}

func (f *repeatableFilter[T]) Add(name, value string) *FilterValueError {
	v, problem := f.parse(name, value)
	if problem != nil {
		return problem
	}
	f.values = append(f.values, v)
	return nil
}

func (f *repeatableFilter[T]) Apply() {
	if len(f.values) > 0 {
		f.set(f.values)
	}
}

// singleFilter keeps one value and rejects a second.
type singleFilter[T any] struct {
	set   func(*T)
	parse func(name, value string) (T, *FilterValueError)
	kept  *T
}

func (f *singleFilter[T]) Add(name, value string) *FilterValueError {
	if f.kept != nil {
		return &FilterValueError{"Duplicate filter", fmt.Sprintf("Filter '%s' takes a single value.", name)}
	}
	v, problem := f.parse(name, value)
	if problem != nil {
		return problem
	}
	f.kept = &v
	return nil
}

func (f *singleFilter[T]) Apply() {
	if f.kept != nil {
		f.set(f.kept)
	}
}

func parseFilterString(_, value string) (string, *FilterValueError) { return value, nil }

func parseFilterInt64(name, value string) (int64, *FilterValueError) {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, &FilterValueError{"Invalid filter value", fmt.Sprintf("Filter '%s' takes an integer, got %q.", name, value)}
	}
	return v, nil
}

func parseFilterFloat64(name, value string) (float64, *FilterValueError) {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, &FilterValueError{"Invalid filter value", fmt.Sprintf("Filter '%s' takes a number, got %q.", name, value)}
	}
	return v, nil
}

func parseFilterBool(name, value string) (bool, *FilterValueError) {
	v, err := strconv.ParseBool(value)
	if err != nil {
		return false, &FilterValueError{"Invalid filter value", fmt.Sprintf("Filter '%s' takes a boolean, got %q.", name, value)}
	}
	return v, nil
}
