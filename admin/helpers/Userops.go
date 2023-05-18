package helpers

import (
	"go_blog/admin/models"
	"net/http"
)

func SetUser(w http.ResponseWriter, r *http.Request, username string, password string) error {
	session, err := store.Get(r, "blog-admin")
	if err != nil {
		return err
	}
	session.Values["username"] = username
	session.Values["password"] = password
	return session.Save(r, w)
}

func CheckUser(w http.ResponseWriter, r *http.Request) bool {
	session, err := store.Get(r, "blog-admin")
	if err != nil {
		return false
	}
	username := session.Values["username"]
	password := session.Values["password"]
	if username != nil && password != nil {
		user := models.User_model{}.Get("username = ? AND password = ?", username, password)
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}

func RemoveUser(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, "blog-admin")
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1

	return session.Save(r, w)
}
