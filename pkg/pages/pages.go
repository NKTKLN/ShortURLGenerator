package pages

import (
	"fmt"
	"net/http"
	"os"

	"github.com/NKTKLN/ShortURLGenerator/pkg/others"
	"github.com/NKTKLN/ShortURLGenerator/pkg/pg"
	"github.com/NKTKLN/ShortURLGenerator/pkg/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Connecting page source
func Error404(c *gin.Context) {c.HTML(http.StatusOK, "pageError404.html", gin.H{"cookie": user.CheckCookie(c)})}

func SignIn(c *gin.Context) {c.HTML(http.StatusOK, "signIn.html", gin.H{"correct": true, "cookie": user.CheckCookie(c), "postUrl": "login"})}

func SignUp(c *gin.Context) {c.HTML(http.StatusOK, "signUp.html", gin.H{"correct": true, "cookie": user.CheckCookie(c)})}

func Reset(c *gin.Context) {c.HTML(http.StatusOK, "resetPassword.html", gin.H{"correct": true, "cookie": user.CheckCookie(c)})}

func TelegramLogin(c *gin.Context) {c.HTML(http.StatusOK, "signIn.html", gin.H{"correct": true, "cookie": user.CheckCookie(c), "postUrl": "telegramLogin"})}

func Home(c *gin.Context) {
	others.ErrorChecking(godotenv.Load("env/bot.env"))
	c.HTML(http.StatusOK, "home.html", gin.H{"correct": true, "bot": os.Getenv("BOT_USERNAME"), "cookie": user.CheckCookie(c)})
}

func Save(c *gin.Context) {
	c.Request.ParseForm()
	// Check the link to see if it works 
	if _, err := http.Get(c.Request.PostForm["InputUrl"][0]); err != nil {
		others.ErrorChecking(godotenv.Load("env/bot.env"))
		c.HTML(http.StatusOK, "home.html", gin.H{"correct": false, "bot": os.Getenv("BOT_USERNAME"), "cookie": user.CheckCookie(c)})
	} else {
		// Saving the link in the database
		userId := user.GetUserIdFromCookie(c)
		linkId := pg.UrlIdGenerator(c.Request.PostForm["InputUrl"][0], userId)
		// Link assembly to be added to the done page
		others.ErrorChecking(godotenv.Load("env/app.env"))
		url := fmt.Sprintf("https://%s/%s", os.Getenv("VIRTUAL_HOST"), linkId)
		c.HTML(http.StatusOK, "done.html", gin.H{"url": url, "cookie": user.CheckCookie(c)})
	}
}

func Redirect(c *gin.Context) {
	url := pg.GetInformationFromDB("url", fmt.Sprintf("id='%s'", c.Param("id")), "shorturl")
	if url == "" {
		c.HTML(http.StatusOK, "pageError404.html", gin.H{"cookie": user.CheckCookie(c)})
	} else {
		// Redirect the user to the site along with an increase in visits by 1 link in the database
		visits := pg.GetInformationFromDBI("visits", fmt.Sprintf("id='%s'", c.Param("id")), "shorturl")
		pg.ExecFromDB(fmt.Sprintf("UPDATE shorturl SET visits=%d WHERE id='%s';", visits + 1, c.Param("id")))
		c.Redirect(http.StatusFound, url)
	}
}

func Archive(c *gin.Context) {
	// Checking whether the user has a cookie
	if user.CheckCookie(c) {
		// Getting the data of all registered links by the user
		db := pg.DbConnect()
		var archiveLinks []others.Url
		others.ErrorChecking(db.Select(&archiveLinks, fmt.Sprintf("SELECT id, url, visits FROM shorturl WHERE userId=%d", user.GetUserIdFromCookie(c))))
		db.Close()
		others.ErrorChecking(godotenv.Load("env/app.env"))
		c.HTML(http.StatusOK, "archive.html", gin.H{"cookie": true, "url": os.Getenv("VIRTUAL_HOST"), "archive": archiveLinks})
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}

func Delete(c *gin.Context) {
	c.Request.ParseForm()
	// Checking cookies and user ID
	userId := pg.GetInformationFromDBI("userId", fmt.Sprintf("id='%s'", c.Request.PostForm["delete"][0]), "shorturl")
	if user.CheckCookie(c) && user.GetUserIdFromCookie(c) == userId {
		// Remove link from database
		pg.ExecFromDB(fmt.Sprintf("DELETE FROM shorturl WHERE id='%s';", c.Request.PostForm["delete"][0]))
		c.Redirect(http.StatusFound, "/archive")
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}
