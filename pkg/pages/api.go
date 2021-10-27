package pages

import (
	"fmt"
	"net/http"

	"github.com/NKTKLN/ShortURLGenerator/pkg/others"
	"github.com/NKTKLN/ShortURLGenerator/pkg/pg"
	"github.com/gin-gonic/gin"
)

// @Summary Add url
// @Description Reference abbreviation
// @Accept  json
// @Produce  json
// @Success 200 {object} others.ApiUrl
// @Param url query string false "URL"
// @Param email query string false "Email"
// @Param password query string false "Password"
// @Router /add [post]
func ApiAdd(c *gin.Context) {
	if _, err := http.Get(c.Query("url")); err != nil {
		c.JSON(http.StatusOK, others.ApiUrl{Status: "Error! Incorrect url."})
	} else {
		passDb := pg.GetInformationFromDB("password", fmt.Sprintf("email='%s'", c.Query("email")), "users")
		if c.Query("email") != "" && c.Query("password") == passDb {
			// Adding a link with user data to the database through the api
			idDb := pg.GetInformationFromDBI("id", fmt.Sprintf("email='%s'", c.Query("email")), "users")
			link := pg.UrlIdGenerator(c.Query("url"), idDb)
			url := fmt.Sprintf("https://%s/%s", others.GetInfromationFromEnv("env/app.env", "VIRTUAL_HOST"), link)
			c.JSON(http.StatusOK, others.ApiUrl{Status: "Good.", Url: url})
		} else if c.Query("email") == "" && c.Query("password") == "" {
			// adding a link to the database through the api
			link := pg.UrlIdGenerator(c.Query("url"), 0)
			url := fmt.Sprintf("https://%s/%s", others.GetInfromationFromEnv("env/app.env", "VIRTUAL_HOST"), link)
			c.JSON(http.StatusOK, others.ApiUrl{Status: "Good.", Url: url})
		} else {
			c.JSON(http.StatusOK, others.ApiUrl{Status: "Error! Incorrect data."})
		}
	}
}

// @Summary Delete url
// @Description Delete url from db
// @Accept  json
// @Produce  json
// @Success 200 {object} others.Api
// @Param id query string false "URL id"
// @Param email query string false "Email"
// @Param password query string false "Password"
// @Router /delete [post]
func ApiDelete(c *gin.Context) {
	passDb := pg.GetInformationFromDB("password", fmt.Sprintf("email='%s'", c.Query("email")), "users")
	if c.Query("email") != "" && c.Query("password") == passDb {
		// Remove link from database
		userIdDb := pg.GetInformationFromDBI("id", fmt.Sprintf("email='%s'", c.Query("email")), "users")
		pg.ExecFromDB(fmt.Sprintf("DELETE FROM shorturl WHERE userId=%d AND id='%s';", userIdDb, c.Query("id")))
		c.JSON(http.StatusOK, gin.H{"type": "Good."})
	} else {
		c.JSON(http.StatusOK, gin.H{"type": "Error! Incorrect data."})
	}
}

// @Summary Archive
// @Description Reference abbreviation
// @Accept  json
// @Produce  json
// @Success 200 {object} others.Api
// @Param email query string false "Email"
// @Param password query string false "Password"
// @Router /archive [get]
func ApiArchive(c *gin.Context) {
	passDb := pg.GetInformationFromDB("password", fmt.Sprintf("email='%s'", c.Query("email")), "users")
	if c.Query("email") != "" && c.Query("password") == passDb {
		// Getting the data of all registered links by the user
		userIdDb := pg.GetInformationFromDBI("id", fmt.Sprintf("email='%s'", c.Query("email")), "users")
		db := pg.DbConnect()
		rows, _ := db.Query(fmt.Sprintf("SELECT id, url, visits FROM shorturl WHERE userId=%d", userIdDb))
		db.Close()
		var archiveLinks []others.Url
		for rows.Next() {
			var archive others.Url
			rows.Scan(&archive.Id, &archive.Url, &archive.Visits)
			archiveLinks = append(archiveLinks, archive)
		}
		c.JSON(http.StatusOK, gin.H{"type": "Good.", "archive": archiveLinks})
	} else {
		c.JSON(http.StatusOK, gin.H{"type": "Error! Incorrect data."})
	}
}
