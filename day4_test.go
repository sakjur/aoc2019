package aoc2019

import "testing"

func TestPasswordCandidates(t *testing.T) {
	//t.Logf("Task 1: %d", len(passwordCandidates(231832, 767346)))
	if n := len(passwordCandidates(231832, 767346)); n != 1330 {
		t.Errorf("expected password candidates for task 1 to be 1330, got %d", n)
	}

	//t.Logf("Task 2: %d", len(PasswordCandidates(231832, 767346)))
	if n := len(PasswordCandidates(231832, 767346)); n != 876 {
		t.Errorf("expected password candidates for task 2 to be 876, got %d", n)
	}
}
