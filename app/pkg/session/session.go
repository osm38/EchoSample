package session

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-contrib/session"
	"github.com/gorilla/sessions"
	"echosample/pkg/log"
	"echosample/pkg/util/crypto"
	"echosample/pkg/db/repository"
)

var (
	logger = log.GetInstance("session")
)

func Auth(sessID string) error {
	return nil
}

func GetSession(c echo.Context) *sessions.Session {
	sess, err := session.Get("GoSampleSession", c)
	if err != nil {
		logger.Error(err, "GetSession is Failed!!", nil)
	}

	si := sess.Values["SessionID"]
	if si == nil {
		sess.Options = &sessions.Options{
			MaxAge: 86400 * 7,
			HttpOnly: false,
		}
		si = StartSession(c)
		sess.Values["SessionID"] = si
	}
	return sess
}

func CreateSessionID() string {
	return crypto.GetSHA256()
}

func RegisterSession(sessID string, name string) {
	user := repository.FindUserByName(name)

	userID := user.ID
	repository.RegisterSession(sessID, userID)
}

func StartSession(c echo.Context) string {
	si := CreateSessionID()
	RegisterSession(si, c.FormValue("name"))
	return si
}