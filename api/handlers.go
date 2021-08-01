package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}
func mainPage(w http.ResponseWriter, r *http.Request) {
	err2 := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err2 != nil {
		log.Fatalln(err2)
	}
}

func singUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("tsssss")
	fName := r.FormValue("first-name")
	lName := r.FormValue("last-name")
	uName := r.FormValue("user-name")
	bio := r.FormValue("bio")
	pass := r.FormValue("password")
	date := r.FormValue("birthdate")
	err := twitter.db.SignUp(fName, lName, uName, pass, bio, date)
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		fmt.Println(err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	uName := r.FormValue("user-name")
	pass := r.FormValue("password")
	err := twitter.db.Login(uName, pass)
	_ = tpl.ExecuteTemplate(w, "result.gohtml", err)
}

func showLogins(w http.ResponseWriter, r *http.Request) {
	logins, err := twitter.db.GetLogins()
	var err2 error
	if err == nil {
		err2 = tpl.ExecuteTemplate(w, "results.gohtml", logins)
	}
	if err2 != nil {
		log.Fatalln(err2)
	}
}

func tweet(w http.ResponseWriter, r *http.Request) {
	txt := r.FormValue("text")
	err := twitter.db.AddTweet(txt)
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		fmt.Println(err)
	}
}

func showOwnTweets(w http.ResponseWriter, r *http.Request) {
	tweets, err := twitter.db.GetOwnTweets()
	var err2 error
	if err == nil {
		err2 = tpl.ExecuteTemplate(w, "results.gohtml", tweets)
	}
	if err2 != nil {
		log.Fatalln(err2)
	}
}

func follow(w http.ResponseWriter, r *http.Request) {
	uName := r.FormValue("user-name")
	err := twitter.db.AddFollowing(uName)
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func unfollow(w http.ResponseWriter, r *http.Request) {
	uName := r.FormValue("user-name")
	err := twitter.db.RemoveFollowing(uName)
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func block(w http.ResponseWriter, r *http.Request) {
	uName := r.FormValue("user-name")
	err := twitter.db.AddBlocked(uName)
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func unblock(w http.ResponseWriter, r *http.Request) {
	uName := r.FormValue("user-name")
	err := twitter.db.RemoveBlocked(uName)
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func showFollowingsTweets(w http.ResponseWriter, r *http.Request) {
	tweets, err := twitter.db.GetFollowingTweets()
	var err2 error
	if err == nil {
		err2 = tpl.ExecuteTemplate(w, "results.gohtml", tweets)
	}
	if err2 != nil {
		log.Fatalln(err2)
	}
}

func timeLine(w http.ResponseWriter, r *http.Request) {
	tweets, err := twitter.db.GetTimeLine()
	var err2 error
	if err == nil {
		err2 = tpl.ExecuteTemplate(w, "results.gohtml", tweets)
	}
	if err2 != nil {
		log.Fatalln(err2)
	}
}

func comments(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	id := r.FormValue("id")
	err := twitter.db.AddComment(text, id)
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		fmt.Println(err)
	}
}

func showComments(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	comments, err := twitter.db.GetComments(id)
	var err2 error
	if err == nil {
		err2 = tpl.ExecuteTemplate(w, "results.gohtml", comments)
	}
	if err2 != nil {
		log.Fatalln(err2)
	}
}

/*func addHashtag(w http.ResponseWriter, r *http.Request) {
	e := r.FormValue("email")
	p := r.FormValue("password")
	err := twitter.db.AddHashtag(e, p)
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}*/

func showHashtag(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	tweets, err := twitter.db.GetHashtagTweets(text)
	var err2 error
	if err == nil {
		err2 = tpl.ExecuteTemplate(w, "results.gohtml", tweets)
	}
	if err2 != nil {
		log.Fatalln(err2)
	}
}

func like(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	err := twitter.db.AddLike(id)
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func likeCounts(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	like, err := twitter.db.GetCountLike(id)
	var err2 error
	if err == nil {
		err2 = tpl.ExecuteTemplate(w, "result.gohtml", like)
	}
	if err2 != nil {
		log.Fatalln(err2)
	}
}

func showLiker(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	likers, err := twitter.db.GetLiker(id)
	var err2 error
	if err == nil {
		err2 = tpl.ExecuteTemplate(w, "results.gohtml", likers)
	}
	if err2 != nil {
		log.Fatalln(err2)
	}
}

func trends(w http.ResponseWriter, r *http.Request) {
	trends, err := twitter.db.GetTrend()
	var err2 error
	if err == nil {
		err2 = tpl.ExecuteTemplate(w, "results.gohtml", trends)
	}
	if err2 != nil {
		log.Fatalln(err2)
	}
}

func message(w http.ResponseWriter, r *http.Request) {
	uName := r.FormValue("user-name")
	text := r.FormValue("text")
	id := r.FormValue("id")
	err := twitter.db.AddMessage(uName, text, id)
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func showMessages(w http.ResponseWriter, r *http.Request) {
	uName := r.FormValue("user-name")
	messages, err := twitter.db.GetMessages(uName)
	var err2 error
	if err == nil {
		err2 = tpl.ExecuteTemplate(w, "results.gohtml", messages)
	}
	if err2 != nil {
		log.Fatalln(err2)
	}
}
func showSenders(w http.ResponseWriter, r *http.Request) {
	senders, err := twitter.db.GetSenders()
	var err2 error
	if err == nil {
		err2 = tpl.ExecuteTemplate(w, "results.gohtml", senders)
	}
	if err2 != nil {
		log.Fatalln(err2)
	}
}
