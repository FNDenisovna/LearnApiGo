package apis

import (
	"context"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var usersPasswords = map[string][]byte{
	"joe":  []byte("$2a$12$aMfFQpGSiPiYkekov7LOsu63pZFaWzmlfm1T8lvG6JFj2Bh4SZPWS"),
	"mary": []byte("$2a$12$l398tX477zeEBP6Se0mAv.ZLR8.LZZehuDgbtw2yoQeMjIyCNCsRW"),
}

// verifyUserPass проверяет, что имя пользователя и пароль соответствуют друг другу,
// сверяясь с нашей "базой данных" userPasswords.
func verifyUserPass(username, password string) bool {
	wantPass, hasUser := usersPasswords[username]
	if !hasUser {
		return false
	}
	if cmperr := bcrypt.CompareHashAndPassword(wantPass, []byte(password)); cmperr == nil {
		return true
	}
	return false
}

const UserContextKey = "user"

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		user, pass, ok := req.BasicAuth()
		if ok && verifyUserPass(user, pass) {
			newctx := context.WithValue(req.Context(), UserContextKey, user)
			next.ServeHTTP(w, req.WithContext(newctx))
		} else {
			w.Header().Set("WWW-Authenticate", `Basic realm="api"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}
