package session

import (
	"github.com/labstack/echo/v4"
	"echosample/pkg/log"
	"echosample/pkg/util/crypto"
	"echosample/pkg/db/repository"
	"fmt"
	"echosample/pkg/db/model"
	"time"
)

var (
	logger = log.GetInstance("session")
)

func Auth(c echo.Context) error {
	_, err := c.Cookie("sessID")
	if err != nil {
		err = fmt.Errorf("Auth Fail!! Wrap: %+w", err)
		logger.Error(err, "Cookie is None!!", nil)
		return err
	}

	sess := GetSession(c)
	user := repository.FindUserByID(sess.UserID)
	maxAge := sess.MaxAge
	lastLogin := user.LastLogin

	if !isEnable(lastLogin, maxAge) {
		err := fmt.Errorf("expored session, SessionID=%s", sess.ID)
		return err
	}

	return nil
}

func GetSession(c echo.Context) model.Session {
	var sess model.Session

	cookie, _ := c.Cookie("sessID")
	fmt.Printf("Cookie:\n%+w", cookie)
	if cookie != nil {
		fmt.Println("GetCookie!")
		sessID := cookie.Value
		sess = repository.FindSessionByID(sessID)
	} else {
		fmt.Println("Cookie is none!")
		sess = StartSession(c)
	}

	return sess
}

func CreateSessionID() string {
	return crypto.GetSHA256()
}

func RegisterSession(sessID string, name string) model.Session {
	user := repository.FindUserByName(name)

	userID := user.ID
	sess := repository.RegisterSession(sessID, userID)

	return sess
}

func StartSession(c echo.Context) model.Session {
	sessID := CreateSessionID()
	sess := RegisterSession(sessID, c.FormValue("name"))
	return sess
}

func isEnable(lastLogin time.Time, maxAge int) bool {
	now := time.Now()
	return now.Before(lastLogin.Add(time.Second * time.Duration(maxAge)))
}

