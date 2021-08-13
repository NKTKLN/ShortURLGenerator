package user

import (
	"fmt"
	"net/http"

	"github.com/NKTKLN/ShortURLGenerator/pkg/pg"
	"github.com/gin-gonic/gin"
)

func AddCookie(id string, time int, c *gin.Context) {
	val := &http.Cookie{
		Name: "id",
		Value: id,
		MaxAge: time,
		Path: "/",
	}
	http.SetCookie(c.Writer, val)
}

func CheckCookie(c *gin.Context) bool {
	_, err := c.Cookie("id")
	return err == nil
}

func GetUserIdFromCookie(c *gin.Context) (userId int) {
	if cookieId, err := c.Cookie("id"); err == nil {
		userId = pg.GetInformationFromDBI("id", fmt.Sprintf("cookieId='%s'", cookieId), "users")
	}
	return
}
