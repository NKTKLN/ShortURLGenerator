package pages

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/NKTKLN/ShortURLGenerator/pkg/others"
	"github.com/NKTKLN/ShortURLGenerator/pkg/pg"
	"github.com/NKTKLN/ShortURLGenerator/pkg/redis"
	"github.com/NKTKLN/ShortURLGenerator/pkg/user"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.Request.ParseForm()
	// Check for mail in the database
	passDb := pg.GetInformationFromDB("password", fmt.Sprintf("email='%s'", c.Request.PostForm["email"][0]), "users")
	if passDb == c.Request.PostForm["password"][0] {
		// Create a cookie for the logined in user
		idDb := pg.GetInformationFromDB("id", fmt.Sprintf("email='%s'", c.Request.PostForm["email"][0]), "users")
		hashId := others.Encoding(idDb)
		pg.ExecFromDB(fmt.Sprintf("UPDATE users SET cookieId='%s' WHERE id='%s';", hashId, idDb))
		user.AddCookie(hashId, 60*60*24*365, c)
		c.Redirect(http.StatusFound, "/")
	} else {
		c.HTML(http.StatusOK, "signIn.html", gin.H{"correct": false, "cookie": user.CheckCookie(c), "postUrl": "login"})
	}
}

func ResetPassword(c *gin.Context) {
	c.Request.ParseForm()
	// Checking for mail in the database
	idDb := pg.GetInformationFromDBI("id", fmt.Sprintf("email='%s'", c.Request.PostForm["email"][0]), "users")
	if idDb != 0 {
		// Create a temporary access key to verify mail fidelity and then send it
		rand.Seed(time.Now().UTC().UnixNano())
		confirmationLinkHash := hex.EncodeToString([]byte(others.Encoding(fmt.Sprintf("%d", rand.Intn(9999999) + 10000000))))
		redis.RedisKeyGenerate(c.Request.PostForm["email"][0], "reset", confirmationLinkHash, idDb)
		c.HTML(http.StatusOK, "correctReset.html", gin.H{"cookie": user.CheckCookie(c)})
	} else {
		c.HTML(http.StatusOK, "resetPassword.html", gin.H{"correct": false, "cookie": user.CheckCookie(c)})
	}
}

func ResetPasswordVerification(c *gin.Context) {
	// Checking that the key is in working order
	client := redis.RedisClientConnect()
	if _, err := client.Get(c.Param("hash")).Result(); err != nil {
		c.HTML(http.StatusOK, "verification.html", gin.H{"correct": false, "cookie": user.CheckCookie(c)})
	} else {
		c.HTML(http.StatusOK, "verificationReset.html", gin.H{"correct": true, "cookie": true, "hash": c.Param("hash")})
	}
}

func ResetPasswordSave(c *gin.Context) {
	c.Request.ParseForm()
	// Checking that the key is in working order
	client := redis.RedisClientConnect()
	if val, err := client.Get(c.Param("hash")).Result(); err != nil {
		c.HTML(http.StatusOK, "verification.html", gin.H{"correct": false, "cookie": user.CheckCookie(c)})
	} else {
		// Adding a new password for a user
		if c.Request.PostForm["passwordRep"][0] == c.Request.PostForm["password"][0] {
			pg.ExecFromDB(fmt.Sprintf("UPDATE users SET password='%s' WHERE id='%s';", c.Request.PostForm["password"][0], val))
			c.Redirect(http.StatusFound, "/login")
		} else {
			c.HTML(http.StatusOK, "verificationReset.html", gin.H{"correct": false, "cookie": user.CheckCookie(c), "hash": c.Param("hash")})
		}
	}
}

func Register(c *gin.Context) {
	c.Request.ParseForm()
	// Database user check
	passDb := pg.GetInformationFromDB("password", fmt.Sprintf("email='%s'", c.Request.PostForm["email"][0]), "users")
	if passDb == "" {
		// Create a temporary access key to verify mail fidelity and then send it
		rand.Seed(time.Now().UTC().UnixNano())
		confirmationLinkHash := hex.EncodeToString([]byte(others.Encoding(fmt.Sprintf("%d", rand.Intn(9999999) + 10000000))))
		json, err := json.Marshal(others.UserReg{Email: c.Request.PostForm["email"][0], Password: c.Request.PostForm["password"][0]})
		others.ErrorChecking(err)
		redis.RedisKeyGenerate(c.Request.PostForm["email"][0], "verification", confirmationLinkHash, json)
		c.HTML(http.StatusOK, "correctRegister.html", gin.H{"cookie": user.CheckCookie(c)})
	} else if c.Request.PostForm["password"][0] == passDb {
		c.HTML(http.StatusOK, "signUp.html", gin.H{"correct": false, "cookie": user.CheckCookie(c)})
	}
}

func VerificationLogin(c *gin.Context) {
	// Checking that the key is in working order
	client := redis.RedisClientConnect()
	if val, err := client.Get(c.Param("hash")).Result(); err != nil {
		c.HTML(http.StatusOK, "verification.html", gin.H{"correct": false, "cookie": user.CheckCookie(c)})
	} else {
		// Adding a new user to the database and generating a cookie for it
		var userParam others.UserReg
		json.Unmarshal([]byte(val), &userParam)
		if pg.GetInformationFromDB("password", fmt.Sprintf("email='%s'", userParam.Email), "users") == "" {
			hashId := pg.GenerateUserId(userParam.Email, userParam.Password)
			user.AddCookie(hashId, 60*60*24*365, c)
		}
		c.HTML(http.StatusOK, "verification.html", gin.H{"correct": true, "cookie": true})	
	}
}

func Logout(c *gin.Context) {
	// Deleting cookies from a user
	userId := user.GetUserIdFromCookie(c)
	pg.ExecFromDB(fmt.Sprintf("UPDATE users SET cookieId=null WHERE id='%d';", userId))
	user.AddCookie("", -1, c)
	c.Redirect(http.StatusFound, "/")
}

func TelegramLoginDone(c *gin.Context) {
	c.Request.ParseForm()
	// Check for mail in the database
	passDb := pg.GetInformationFromDB("password", fmt.Sprintf("email='%s'", c.Request.PostForm["email"][0]), "users")
	if passDb == c.Request.PostForm["password"][0] {
		// Create a temporary access key to verify telegram and then send it
		rand.Seed(time.Now().UTC().UnixNano())
		idDb := pg.GetInformationFromDB("id", fmt.Sprintf("email='%s'", c.Request.PostForm["email"][0]), "users")
		redis.RedisKeyGenerate(c.Request.PostForm["email"][0], "telegram", fmt.Sprintf("%d", rand.Intn(9999999) + 10000000), idDb)
		c.HTML(http.StatusOK, "correctRegister.html", gin.H{"cookie": user.CheckCookie(c)})
	} else {
		c.HTML(http.StatusOK, "signIn.html", gin.H{"correct": false, "cookie": user.CheckCookie(c), "postUrl": "telegramLogin"})
	}
}
