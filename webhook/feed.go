package webhook

import (
	"fmt"
	"os"
	"github.com/mmcdole/gofeed"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sqlx.DB
	feedURL = os.Getenv("FEED_URL")
)

//Establish DBの初期化
func Establish() error {
	_db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("MARIADB_USERNAME"), os.Getenv("MARIADB_PASSWORD"), os.Getenv("MARIADB_HOSTNAME"), "3306", os.Getenv("MARIADB_DATABASE")))
	if err != nil {
		return err
	}
	db = _db
	return nil
}

//UpdateFeed feedの更新
func UpdateFeed() error {
	fp := gofeed.NewParser()

    feed, _ := fp.ParseURL(feedURL)
	items := feed.Items

	for _, item := range items {
		var link string
		err := db.Get(&link, "SELECT link FROM feed WHERE link=?", item.Link)
		if err != nil {
			body := "## [" + item.Title + "](" + item.Link + ")\n### " + item.Published + "\n" + item.Description
			err = postMessage(body)
			if err != nil {
				return err
			}
			_,err = db.Exec("INSERT INTO feed (title,link,published,description) VALUES (?,?,?,?)",item.Title,item.Link,item.Published,item.Description)
			if err != nil {
				return err
			}
		}
	}
	return nil
}