package main_test

import (
	"fmt"
	"testing"

	"github.com/quartzwine/friend-tech-utils/friendtech"
)

// Utility function
func doSomethingHelper(input string) string {
	// ...
	return "expected"
}

// Actual test case
func TestSomething(t *testing.T) {
	result := doSomethingHelper("input")
	if result != "expected" {
		t.Errorf("Expected 'expected', got '%s'", result)
	}
	fmt.Print("boom\n")
}

func TestGetRecentlyJoinedv1(t *testing.T) {
	users, err := friendtech.GetRecentlyJoinedv1()
	if err != nil {
		t.Errorf("Error getting recently joined %v", err)
	}

	for _, user := range users.Users {
		fmt.Printf("address: %s\n", user.Address)

	}
}
