package main_test

import (
	"fmt"
	"testing"

	"github.com/quartzwine/friend-tech-utils/friendtech"
	"github.com/quartzwine/friend-tech-utils/db"
	"github.com/quartzwine/friend-tech-utils/twitter"

	twitterscraper "github.com/n0madic/twitter-scraper"
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

func TestGetRecentlyJoinedv2(t *testing.T) {
	scraper := twitterscraper.New()
	addresses := db.Get_addresses()

	//fmt.Println("Addresses:")
	for _, address := range addresses {
		username, err := twitter.GetTwitterUsernameFromAddress(address)
		if (err != nil) {
			panic(err)
		}
		profile, err := scraper.GetProfile(username)
		if (err != nil) {
			fmt.Printf("couldnt find username: %s\n", username)
			continue
		}
		if (profile.FollowersCount > 1000) {
			fmt.Printf("Address: %s Username: %s, Followers: %d\n", address, username, profile.FollowersCount)
		} else{
			fmt.Print("skipping user under 1k followers\n")
		}
		
	}

}

func TestDbSetup(t *testing.T) {
	addresses := db.Get_addresses()

	for _, address := range addresses{
		fmt.Println(address)
	}

	fmt.Println("boom")

}
