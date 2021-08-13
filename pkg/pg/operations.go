package pg

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/NKTKLN/ShortURLGenerator/pkg/others"
	_ "github.com/lib/pq"
)

// Getting information from db
func GetInformationFromDB(column, where, table string) (info string) {
	db := DbConnect()
	others.ErrorChecking(db.Get(&info, fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT 1;", column, table, where)))
	db.Close()
	return 
}

func GetInformationFromDBI(column, where, table string) (info int) {
	db := DbConnect()
	db.Get(&info, fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT 1;", column, table, where))
	db.Close()
	return 
}

func ExecFromDB(request string) error {
	db := DbConnect()
	_, err := db.Exec(request)
	db.Close()
	return err
}

// Generating a url id and adding it to the db
func UrlIdGenerator(url string, userId int) (link string) {
	link = GetInformationFromDB("id", fmt.Sprintf("url='%s' AND userId=%d", url, userId), "shorturl")
	// Checking for a link in the db
	if link == "" {
		for {
			// Generating a random id
			linkGen := ""
			words := strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789", "")
			for i:=0;i<=5;i++ {
				rand.Seed(time.Now().UTC().UnixNano())
				linkGen += words[rand.Intn(62)]
			}
			// Adding values to the db
			if ExecFromDB(fmt.Sprintf("INSERT INTO shorturl (id, url, userId, visits) VALUES ('%s', '%s', %d, 0);", linkGen, url, userId)) == nil {
				link = linkGen
				break
			}
		}
	} 
	return 
}

func GenerateUserId(email, password string) (hashId string) {
	for {
		// Generating a random id
		rand.Seed(time.Now().UTC().UnixNano())
		id := fmt.Sprintf("%d", rand.Intn(9999999) + 10000000)
		hashId = others.Encoding(id)
		// Adding user data to the database
		if ExecFromDB(fmt.Sprintf("INSERT INTO users (id, email, password, cookieId, telegramId) VALUES (%s, '%s', '%s', '%s', null);", id, email, password, hashId)) == nil {
			return
		}
	}
}
