package main

import (
	"reflect"
	"testing"
)

func TestSCCs(t *testing.T) {
	// 0->1->2->0 is one SCC; 3 and 4 are singletons.
	got := SCCs(5, [][2]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}, {3, 4}})
	want := [][]int{{0, 1, 2}, {3}, {4}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SCCs = %v, want %v", got, want)
	}

	// Two separate cycles joined by a one-way bridge stay separate SCCs.
	got = SCCs(6, [][2]int{{0, 1}, {1, 0}, {1, 2}, {2, 3}, {3, 2}, {4, 5}})
	want = [][]int{{0, 1}, {2, 3}, {4}, {5}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SCCs = %v, want %v", got, want)
	}

	// A DAG: every vertex its own component.
	got = SCCs(3, [][2]int{{0, 1}, {1, 2}})
	want = [][]int{{0}, {1}, {2}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("DAG SCCs = %v, want %v", got, want)
	}

	if got := SCCs(0, nil); len(got) != 0 {
		t.Errorf("empty graph = %v", got)
	}
}
