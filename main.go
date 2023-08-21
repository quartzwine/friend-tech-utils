package main

import (
	"fmt"
	"github.com/quartzwine/friend-tech-utils/db"
	"github.com/quartzwine/friend-tech-utils/twitter"
    "github.com/pkg/profile"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

func main() {
    scraper := twitterscraper.New()
    addresses := db.Get_addresses()

    defer profile.Start(profile.CPUProfile).Stop()

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