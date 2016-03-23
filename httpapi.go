// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	github "github.com/kkdai/githubrss"
)

// func ServeFollowed(w http.ResponseWriter, r *http.Request) {
// 	key := r.RequestURI
// 	switch {
// 	case r.Method == "GET":
// 		//call api
// 		key = key[1:]
// 		log.Println("[GET]FOLLOWED ID", key)
// 		rss := github.NewGithubRss(key)
// 		ret, _ := rss.GetFollower(5)
// 		fmt.Fprintf(w, ret)
// 		// log.Println("[GET] ID", ret)
// 	default:
// 		w.Header().Add("Allow", "GET")
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	}
// }

func starred(w http.ResponseWriter, r *http.Request) {
	id := strings.Trim(r.RequestURI, "/starred")
	id = strings.Trim(id, "?")
	rss := github.NewGithubRss(id)
	ret, _ := rss.GetStarred(10)
	io.WriteString(w, ret)
}

func follower(w http.ResponseWriter, r *http.Request) {
	id := strings.Trim(r.RequestURI, "/follower")
	id = strings.Trim(id, "?")
	rss := github.NewGithubRss(id)
	ret, _ := rss.GetFollower(10)
	io.WriteString(w, ret)
}

func following(w http.ResponseWriter, r *http.Request) {
	id := strings.Trim(r.RequestURI, "/following")
	id = strings.Trim(id, "?")
	rss := github.NewGithubRss(id)
	ret, _ := rss.GetFollowing(10)
	io.WriteString(w, ret)
}

func serveHttpAPI(port string, existC chan bool) {
	// exit when receivee error channel
	go func() {
		if err, ok := <-existC; ok {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/starred", starred)
	mux.HandleFunc("/follower", follower)
	mux.HandleFunc("/following", following)
	http.ListenAndServe(":"+port, mux)
}
