package api

type DataBase interface {
	SignUp(firstName, lastName, username, password, bio, date string) error
	Login(username, password string) error
	GetLogins() (map[string]string, error)
	AddTweet(text string) error
	GetOwnTweets() ([]string, error)
	AddFollowing(username string) error
	RemoveFollowing(username string) error
	AddBlocked(username string) error
	RemoveBlocked(username string) error
	GetFollowingTweets() ([]string, error)
	GetTimeLine() ([]string, error)
	AddComment(text, tweetId string) error
	GetComments(id string) ([]string, error)
	//AddHashtag(username, password string) error
	GetHashtagTweets(hashtag string) ([]string, error)
	AddLike(id string) error
	GetCountLike(id string) (int, error)
	GetLiker(id string) ([]string, error)
	GetTrend() ([]string, error)
	AddMessage(username, text, id string) error
	GetMessages(userName string) ([]string, error)
	GetSenders() ([]string, error)
}
