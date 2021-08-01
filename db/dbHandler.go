package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

type DataBase struct {
	Client *sql.DB
}

///////////////////////////////////////////////////// new redis database :

func GetNewDatabase() *DataBase {
	dbDriver := "mysql"
	dbUser := "armin2"
	dbPass := "3011@Rmin3011="
	dbName := "ghanri"
	client, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err)
	}
	if err := client.Ping(); err != nil {
		fmt.Println(err)
		panic(err)
	}
	log.Print("dataBase successfully configured")
	return &DataBase{
		Client: client,
	}
}

func (db *DataBase) SignUp(firstName, lastName, username, password, bio, date string) error {
	fmt.Println("tss222")
	_, err := db.Client.Query("call signup(?, ?, ?, ?, ?, ?)", firstName, lastName, username, password, date, bio)
	return err
}

func (db *DataBase) Login(username, password string) error {
	//_, err := db.Client.Query("call login(?, ?)", username, password)
	result, err := db.Client.Exec("call login(?, ?)", username, password)
	if err != nil {
		return err
	}
	effected, _ := result.RowsAffected()
	if effected == 0 {
		return errors.New("something went wrong and did not logged in")
	} else {
		return errors.New("logged in")
	}
}

func (db *DataBase) GetLogins() (map[string]string, error) {
	userName := db.getLogged()
	fmt.Println(userName)
	rows, err := db.Client.Query("call getLogins(?);", userName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		logs := make(map[string]string)
		for rows.Next() {
			var userName string
			var time string
			err2 := rows.Scan(&userName, &time)
			if err2 != nil {
				return nil, err2
			} else {
				logs[time] = userName
			}
		}
		fmt.Println(logs)
		return logs, nil
	}
}

func (db *DataBase) AddTweet(text string) error {
	_, err := db.Client.Query("call addVoice(?, ?)", db.getLogged(), text)
	return err
}

func (db *DataBase) GetOwnTweets() ([]string, error) {
	userName := db.getLogged()
	fmt.Println(userName)
	rows, err := db.Client.Query("call getOwnTweets(?);", userName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		logs := make([]string, 0)
		for rows.Next() {
			var userName string
			var text string
			var time string
			err2 := rows.Scan(&userName, &text, &time)
			if err2 != nil {
				return nil, err2
			} else {
				logs = append(logs, userName+" said in "+time+" : "+text)
			}
		}
		fmt.Println(logs)
		return logs, nil
	}
}

func (db *DataBase) AddFollowing(username string) error {
	_, err := db.Client.Query("call follow(?, ?)", db.getLogged(), username)
	return err
}

func (db *DataBase) RemoveFollowing(username string) error {
	_, err := db.Client.Query("call unfollow(?, ?)", db.getLogged(), username)
	return err
}

func (db *DataBase) AddBlocked(username string) error {
	_, err := db.Client.Query("call block(?, ?)", db.getLogged(), username)
	return err
}

func (db *DataBase) RemoveBlocked(username string) error {
	_, err := db.Client.Query("call unblock(?, ?)", db.getLogged(), username)
	return err
}

func (db *DataBase) GetFollowingTweets() ([]string, error) {
	rows, err := db.Client.Query("call getFollowingTweets(?);", db.getLogged())
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		logs := make([]string, 0)
		for rows.Next() {
			var userName string
			var text string
			var time string
			err2 := rows.Scan(&userName, &text, &time)
			if err2 != nil {
				return nil, err2
			} else {
				logs = append(logs, userName+" said in "+time+" : "+text)
			}
		}
		return logs, nil
	}
}
func (db *DataBase) GetTimeLine() ([]string, error) {
	rows, err := db.Client.Query("call timeLine(?);", db.getLogged())
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		logs := make([]string, 0)
		for rows.Next() {
			var userName string
			var text string
			var time string
			err2 := rows.Scan(&userName, &text, &time)
			if err2 != nil {
				return nil, err2
			} else {
				logs = append(logs, userName+" said in "+time+" : "+text)
			}
		}
		return logs, nil
	}
	return nil, nil
}
func (db *DataBase) AddComment(text, tweetId string) error {
	intid, _ := strconv.Atoi(tweetId)
	_, err := db.Client.Query("call addComment(?, ?,?)", db.getLogged(), text, intid)
	return err
}

func (db *DataBase) GetComments(id string) ([]string, error) {
	intid, _ := strconv.Atoi(id)
	rows, err := db.Client.Query("call showComments(?,?);", db.getLogged(), intid)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		logs := make([]string, 0)
		for rows.Next() {
			var userName string
			var text string
			var time string
			err2 := rows.Scan(&userName, &text, &time)
			if err2 != nil {
				fmt.Println(err2)
				return nil, err2
			} else {
				logs = append(logs, userName+" said in "+time+" : "+text)
			}
		}
		fmt.Println(logs)
		return logs, nil
	}
}

func (db *DataBase) GetHashtagTweets(hashtag string) ([]string, error) {
	rows, err := db.Client.Query("call showHashtagTweets(?);", hashtag)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		logs := make([]string, 0)
		for rows.Next() {
			var userName string
			var text string
			var time string
			err2 := rows.Scan(&userName, &text, &time)
			if err2 != nil {
				fmt.Println(err2)
				return nil, err2
			} else {
				logs = append(logs, userName+" said in "+time+" : "+text)
			}
		}
		fmt.Println(logs)
		return logs, nil
	}
}

func (db *DataBase) AddLike(id string) error {
	intid, _ := strconv.Atoi(id)
	_, err := db.Client.Query("call likeTweet(?, ?)", db.getLogged(), intid)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (db *DataBase) GetCountLike(id string) (int, error) {
	intid, _ := strconv.Atoi(id)
	rows, err := db.Client.Query("call showLike(?,?);", db.getLogged(), intid)
	count := 0
	if err != nil {
		fmt.Println(err)
		return 0, err
	} else {
		for rows.Next() {
			err2 := rows.Scan(&count)
			if err2 != nil {
				fmt.Println(err2)
				return 0, err2
			}
		}
	}
	return count, nil
}

func (db *DataBase) GetLiker(id string) ([]string, error) {
	intid, _ := strconv.Atoi(id)
	rows, err := db.Client.Query("call showLikers(?,?);", db.getLogged(), intid)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		logs := make([]string, 0)
		for rows.Next() {
			var userName string
			err2 := rows.Scan(&userName)
			if err2 != nil {
				fmt.Println(err2)
				return nil, err2
			} else {
				logs = append(logs, userName)
			}
		}
		fmt.Println(logs)
		return logs, nil
	}
}

func (db *DataBase) GetTrend() ([]string, error) {
	rows, err := db.Client.Query("call trends(?);", db.getLogged())
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		logs := make([]string, 0)
		for rows.Next() {
			var userName string
			var text string
			var time string
			var count int
			err2 := rows.Scan(&userName, &text, &time, &count)
			if err2 != nil {
				fmt.Println(err2)
				return nil, err2
			} else {
				logs = append(logs, userName+" said in "+time+" : "+text+" count likes : "+strconv.Itoa(count))
			}
		}
		return logs, nil
	}
}
func (db *DataBase) AddMessage(username, text, id string) error {
	intid, _ := strconv.Atoi(id)
	fmt.Println(text)
	if intid != 0 {
		_, err := db.Client.Query("call voiceMessage(?, ?,?)", db.getLogged(), username, intid)
		if err != nil {
			fmt.Println(err)
		}
		return err
	} else {
		_, err := db.Client.Query("call textMessage(?, ?,?)", db.getLogged(), username, text)
		if err != nil {
			fmt.Println(err)
		}
		return err
	}
}
func (db *DataBase) GetSenders() ([]string, error) {
	rows, err := db.Client.Query("call showSenders(?);", db.getLogged())
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		logs := make([]string, 0)
		for rows.Next() {
			var userName string
			err2 := rows.Scan(&userName)
			if err2 != nil {
				fmt.Println(err2)
				return nil, err2
			} else {
				logs = append(logs, userName)
			}
		}
		fmt.Println(logs)
		return logs, nil
	}
}

func (db *DataBase) GetMessages(userName string) ([]string, error) {
	messages := make([]string, 0)
	messages = append(messages, "Voice messages :")
	voiceMessages, _ := db.GetVoiceMessages(userName)
	messages = append(messages, voiceMessages...)
	messages = append(messages, "text messages :")
	textMessages, _ := db.GetTextMessages(userName)
	messages = append(messages, textMessages...)
	fmt.Println(textMessages)
	return messages, nil
}

func (db *DataBase) GetVoiceMessages(userName string) ([]string, error) {
	rows, err := db.Client.Query("call showVoiceMessages(?,?);", db.getLogged(), userName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		logs := make([]string, 0)
		for rows.Next() {
			var date string
			var voice int
			var username string
			var text string
			err2 := rows.Scan(&date, &voice, &username, &text)
			if err2 != nil {
				fmt.Println(err2)
				return nil, err2
			} else {
				logs = append(logs, userName+" in  "+date+"  sent tweet : "+strconv.Itoa(voice)+"  ( "+username+" : "+text+" )")
			}
		}
		return logs, nil
	}
}

func (db *DataBase) GetTextMessages(userName string) ([]string, error) {
	rows, err := db.Client.Query("call showTextMessages(?,?);", db.getLogged(), userName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		logs := make([]string, 0)
		for rows.Next() {
			var date string
			var text string
			err2 := rows.Scan(&date, &text)
			if err2 != nil {
				fmt.Println(err2)
				return nil, err2
			} else {
				logs = append(logs, userName+" in  "+date+"  said : "+text)
			}
		}
		fmt.Println(logs)
		fmt.Println(userName)
		return logs, nil
	}
}

func (db *DataBase) getLogged() string {
	rows, err := db.Client.Query("call getLogged();")
	logs := make([]string, 0)
	if err != nil {
		fmt.Println(err)
		return ""
	} else {
		for rows.Next() {
			var userName string
			err2 := rows.Scan(&userName)
			if err2 != nil {
				return ""
			} else {
				logs = append(logs, userName)
			}
		}
	}
	return logs[0]
}
