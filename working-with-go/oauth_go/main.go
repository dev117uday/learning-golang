package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	GOOGLE_Client_ID     string
	GOOGLE_Client_Secret string
	googleOauthConfig           = &oauth2.Config{}
	randomState          string = "random"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	GOOGLE_Client_ID = os.Getenv("GOOGLE_Client_ID")
	GOOGLE_Client_Secret = os.Getenv("GOOGLE_Client_Secret")

	googleOauthConfig.RedirectURL = "http://localhost:8080/callback"
	googleOauthConfig.ClientID = GOOGLE_Client_ID
	googleOauthConfig.ClientSecret = GOOGLE_Client_Secret
	googleOauthConfig.Scopes = []string{"https://www.googleapis.com/auth/userinfo.profile"}
	googleOauthConfig.Endpoint = google.Endpoint

}

func main() {

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

func handleHome(w http.ResponseWriter, r *http.Request) {
	var htmlIndex = `<html>
		<body>
			<a href="/login">Google Log In</a>
		</body>
		</html>
	`

	fmt.Fprintf(w, htmlIndex)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(randomState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintf(w, "Content: %s\n", content)
}

func getUserInfo(state string, code string) ([]byte, error) {
	if state != randomState {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}
