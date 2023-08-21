package twitter
import ("fmt"
	"encoding/json"
	"net/http"
)


func GetTwitterUsernameFromAddress(address string) (string,error) {
	url := fmt.Sprintf("https://prod-api.kosetto.com/users/%s", address)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("Error fetching URL: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		TwitterUsername string `json:"twitterUsername"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("Error decoding JSON: %v", err)
	}

	return result.TwitterUsername, nil

}