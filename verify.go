package main

import (
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"golang.org/x/oauth2"
)

func initVerify() {
	endpoints := map[string]oauth2.Endpoint{
		"github": {
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
		"google": {
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
	}
	authConfig = map[string]oauth2.Config{}
	for name, ep := range endpoints {
		client, ok := gConf.OAuth[name]
		if ok {
			authConfig[name] = oauth2.Config{
				ClientID:     client.ClientID,
				ClientSecret: client.ClientSecret,
				Endpoint:     ep,
				RedirectURL:  "http://" + gConf.Hostname + "/verify/" + name + "/redirect",
				Scopes:       []string{},
			}
		}
	}
}

var authConfig map[string]oauth2.Config

func verify(w http.ResponseWriter, r *http.Request, params martini.Params) {
	name := params["type"]
	cfg, ok := authConfig[name]
	if !ok {
		http.NotFound(w, r)
		return
	}
	url := cfg.AuthCodeURL(randomState())
	http.Redirect(w, r, url, http.StatusFound)
}

func verifyRedirect(w http.ResponseWriter, r *http.Request, params martini.Params) {
	name := params["type"]
	cfg, ok := authConfig[name]
	if !ok {
		http.NotFound(w, r)
		return
	}
	query := r.URL.Query()
	log.Println(query)
	if "" != query.Get("error") {
		http.Error(w, query.Get("error_description"), http.StatusBadRequest)
		return
	}
	token, _ := cfg.Exchange(oauth2.NoContext, query.Get("code"))
	http.Error(w, token.AccessToken, http.StatusAccepted)
	//more todo
}

func randomState() string {
	return "thecakeisalie"
}
