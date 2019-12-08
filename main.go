package main

import (
	"fmt"
	// "reflect"
	"net/http"
	"io/ioutil"
	// "encoding/json"
	// "time"

	"github.com/tidwall/gjson"
)

var link_github_release = "https://api.github.com/repos/rabbitmq/erlang-rpm/releases"

// type Struct_from_releases []struct {
// 	URL             string `json:"url"`
// 	AssetsURL       string `json:"assets_url"`
// 	UploadURL       string `json:"upload_url"`
// 	HTMLURL         string `json:"html_url"`
// 	ID              int    `json:"id"`
// 	NodeID          string `json:"node_id"`
// 	TagName         string `json:"tag_name"`
// 	TargetCommitish string `json:"target_commitish"`
// 	Name            string `json:"name"`
// 	Draft           bool   `json:"draft"`
// 	Author          struct {
// 		Login             string `json:"login"`
// 		ID                int    `json:"id"`
// 		NodeID            string `json:"node_id"`
// 		AvatarURL         string `json:"avatar_url"`
// 		GravatarID        string `json:"gravatar_id"`
// 		URL               string `json:"url"`
// 		HTMLURL           string `json:"html_url"`
// 		FollowersURL      string `json:"followers_url"`
// 		FollowingURL      string `json:"following_url"`
// 		GistsURL          string `json:"gists_url"`
// 		StarredURL        string `json:"starred_url"`
// 		SubscriptionsURL  string `json:"subscriptions_url"`
// 		OrganizationsURL  string `json:"organizations_url"`
// 		ReposURL          string `json:"repos_url"`
// 		EventsURL         string `json:"events_url"`
// 		ReceivedEventsURL string `json:"received_events_url"`
// 		Type              string `json:"type"`
// 		SiteAdmin         bool   `json:"site_admin"`
// 	} `json:"author"`
// 	Prerelease  bool      `json:"prerelease"`
// 	CreatedAt   time.Time `json:"created_at"`
// 	PublishedAt time.Time `json:"published_at"`
// 	Assets      []struct {
// 		URL      string `json:"url"`
// 		ID       int    `json:"id"`
// 		NodeID   string `json:"node_id"`
// 		Name     string `json:"name"`
// 		Label    string `json:"label"`
// 		Uploader struct {
// 			Login             string `json:"login"`
// 			ID                int    `json:"id"`
// 			NodeID            string `json:"node_id"`
// 			AvatarURL         string `json:"avatar_url"`
// 			GravatarID        string `json:"gravatar_id"`
// 			URL               string `json:"url"`
// 			HTMLURL           string `json:"html_url"`
// 			FollowersURL      string `json:"followers_url"`
// 			FollowingURL      string `json:"following_url"`
// 			GistsURL          string `json:"gists_url"`
// 			StarredURL        string `json:"starred_url"`
// 			SubscriptionsURL  string `json:"subscriptions_url"`
// 			OrganizationsURL  string `json:"organizations_url"`
// 			ReposURL          string `json:"repos_url"`
// 			EventsURL         string `json:"events_url"`
// 			ReceivedEventsURL string `json:"received_events_url"`
// 			Type              string `json:"type"`
// 			SiteAdmin         bool   `json:"site_admin"`
// 		} `json:"uploader"`
// 		ContentType        string    `json:"content_type"`
// 		State              string    `json:"state"`
// 		Size               int       `json:"size"`
// 		DownloadCount      int       `json:"download_count"`
// 		CreatedAt          time.Time `json:"created_at"`
// 		UpdatedAt          time.Time `json:"updated_at"`
// 		BrowserDownloadURL string    `json:"browser_download_url"`
// 	} `json:"assets"`
// 	TarballURL string `json:"tarball_url"`
// 	ZipballURL string `json:"zipball_url"`
// 	Body       string `json:"body"`
// }

func main() {

	// var object_struct_from_releases []Struct_from_releases
	
	response, err := http.Get(link_github_release)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	bodyString := string(body)

	value := gjson.Parse(bodyString).Get("assets_url")

	fmt.Println(value.String())

	// fmt.Println(bodyString)

    // if err := json.Unmarshal([]byte(bodyString), &object_struct_from_releases); err != nil {
    //     panic(err)
	// }

    // fmt.Printf("%+v\n", object_struct_from_releases)

}
