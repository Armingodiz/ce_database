package api

import (
	"github.com/gorilla/mux"
)

var twitter *Twitter

func SetTwitter(db DataBase) {
	twitter = NewTwitter(db)
}

type Twitter struct {
	router *mux.Router
	db     DataBase
}

func NewTwitter(db DataBase) *Twitter {
	return &Twitter{
		router: newRouter(),
		db:     db,
	}
}
func GetRouter() *mux.Router {
	return twitter.router
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", mainPage)
	r.HandleFunc("/signup", singUp)
	r.HandleFunc("/login", login)
	r.HandleFunc("/logins", showLogins)
	r.HandleFunc("/tweet", tweet)
	r.HandleFunc("/ownTweets", showOwnTweets)
	r.HandleFunc("/follow", follow)
	r.HandleFunc("/unfollow", unfollow)
	r.HandleFunc("/block", block)
	r.HandleFunc("/unblock", unblock)
	r.HandleFunc("/followingsTweet", showFollowingsTweets)
	r.HandleFunc("/timeLine", timeLine)
	r.HandleFunc("/comment", comments)
	r.HandleFunc("/showComments", showComments)
	//r.HandleFunc("/addHashtag", addHashtag)
	r.HandleFunc("/showHashtag", showHashtag)
	r.HandleFunc("/like", like)
	r.HandleFunc("/likeCounts", likeCounts)
	r.HandleFunc("/showLiker", showLiker)
	r.HandleFunc("/trends", trends)
	r.HandleFunc("/message", message)
	r.HandleFunc("/showMessages", showMessages)
	r.HandleFunc("/showSenders", showSenders)
	return r
}
