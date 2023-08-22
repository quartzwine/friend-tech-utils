package main

import (
	"fmt"
	"github.com/quartzwine/friend-tech-utils/db"
	"github.com/quartzwine/friend-tech-utils/twitter"
    "github.com/pkg/profile"

	twitterscraper "github.com/n0madic/twitter-scraper"
    "time"
)

func main() {
	scraper := twitterscraper.New()
	defer profile.Start(profile.CPUProfile, profile.ProfilePath("profiler")).Stop()
	ticker := time.NewTicker(10 * time.Second)
    counter := 0
	for range ticker.C {
        if counter > 2 {
            panic("ended")
        }
		latestBlock := db.Get_latest_block_number()
		blockToQuery := latestBlock - 20

		fmt.Printf("processing block: %d\n", blockToQuery)

		addresses := db.Get_addresses_from_block(blockToQuery)
		for _, address := range addresses {
			username, err := twitter.GetTwitterUsernameFromAddress(address)
			if err != nil {
				panic(err)
			}
			profile, err := scraper.GetProfile(username)
			if err != nil {
				fmt.Printf("couldnt find username: %s\n", username)
				continue
			}
			if profile.FollowersCount > 1000 {
				colorCode := getColorCode(profile.FollowersCount)
				fmt.Printf("%sUsername: %s, Followers: %d, Address: %s\033[0m\n", colorCode, username, profile.FollowersCount, address)
			} else {
				fmt.Print("skipping user under 1k followers\n")
			}
		}
        counter +=1
	}
}


func getColorCode(followersCount int) string {
	switch {
	case followersCount <= 1000:
		return "\033[37m" // White for 1000 and under
	case followersCount > 1000 && followersCount <= 10000:
		return "\033[33m" // Yellow for 1000 - 10,000
	case followersCount > 10000 && followersCount <= 100000:
		return "\033[38;2;255;165;0m" // Orange (RGB) for 10,000 - 100,000
	case followersCount > 100000 && followersCount <= 1000000:
		return "\033[38;2;144;238;144m" // Light Green (RGB) for 100,000 - 1,000,000
	default:
		return "\033[32m" // Vibrant Green for above 1,000,000
	}
}