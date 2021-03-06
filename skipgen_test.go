package main

import (
	"testing"
)

func TestStringInSlicePositive(t *testing.T) {
	slice := []string{"one", "two", "three", "four", "five"}

	result := stringInSlice("one", slice)
	if ! result {
		t.Errorf("'one' not found in %s", slice)
	}
}
func TestStringInSliceNegative(t *testing.T) {
	slice := []string{"one", "two", "three", "four", "five"}

	result := stringInSlice("zero", slice)
	if result {
		t.Errorf("'zero' found in %s", slice)
	}
}

func TestSkipAll(t *testing.T) {

	skipAll := `
skiplist:
  - reason: kernel tests baselining
    url: https://projects.linaro.org/projects/CTT/queues/issue/CTT-585
    environments: all # Test this form
    boards:
      - all # And test this form
    branches:
      - all
    tests:
      - test_maps
      - test_lru_map
      - test_lpm_map
      - test_progs
  - reason: "LKFT: linux-next: vm compaction_test : ERROR: Less that 1/730 of memory
             is available"
    url: https://bugs.linaro.org/show_bug.cgi?id=3145
    environments:
      - production
      - staging
    boards:
      - x15
      - juno
      - hikey
    branches:
      - "4.4"
      - 4.9
      - mainline
    tests:
      - run_vmtests
  - reason: duplicate skip for specific board combination
    url: https://bugs.linaro.org/show_bug.cgi?id=3145
    environments:
      - production
    boards:
      - x86
    branches:
      - 4.14
    tests:
      - run_vmtests
`

	skips, err := parseSkipfile([]byte(skipAll))
	if err != nil {
		t.Errorf("Unexpected error parsing yaml, %s", err)
	}

	t.Run("parseSkipfile spotcheck", func(t *testing.T) {
		if skips.Skiplist[0].Reason != "kernel tests baselining" {
			t.Errorf("Parsing error, skiplist is wrong")
		}
	})

    t.Run("getSkipfileContents", func(t *testing.T) {
		if getSkipfileContents("x15", "4.4", "production", skips) !=
`run_vmtests
test_lpm_map
test_lru_map
test_maps
test_progs
`{
			t.Errorf("Incorrect Skipfile Contents")
		}
	})

	// Test deduplication, and 'all'
    t.Run("getSkipfileContents", func(t *testing.T) {
		if getSkipfileContents("all", "all", "all", skips) !=
`run_vmtests
test_lpm_map
test_lru_map
test_maps
test_progs
`{
			t.Errorf("Incorrect Skipfile Contents")
		}
	})

}

func TestSkipMinimum(t *testing.T) {

	skipAll := `
skiplist:
  - reason: Some test reason
    url: https://bugs.linaro.org/show_bug.cgi?id=3145
    environments: production # Test UnmarshallYAML
    boards: x15 # Test UnmarshallYAML
    branches: "4.4" # Test UnmarshallYAML
    tests: run_vmtests # Test UnmarshallYAML
`
	skips, err := parseSkipfile([]byte(skipAll))
	if err != nil {
		t.Errorf("Unexpected error parsing yaml, %s", err)
	}

	t.Run("parseSkipfile spotcheck", func(t *testing.T) {
		if skips.Skiplist[0].Reason != "Some test reason" {
			t.Errorf("Parsing error, skiplist is wrong")
		}
	})

    t.Run("getSkipfileContents positive", func(t *testing.T) {
		if getSkipfileContents("x15", "4.4", "production", skips) !=
`run_vmtests
` {
			t.Errorf("Incorrect Skipfile Contents")
		}
	})

    t.Run("getSkipfileContents empty", func(t *testing.T) {
		if getSkipfileContents("x15", "4.9", "production", skips) != "" {
			t.Errorf("Incorrect Skipfile Contents")
		}
	})
}
