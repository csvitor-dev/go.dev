package cookies

import (
	"net/http"

	"github.com/csvitor-dev/go.dev/internal/config/env"
	"github.com/gorilla/securecookie"
)

var cookie *securecookie.SecureCookie

func Active() {
	cookie = securecookie.New(env.WebEnv.HASH_KEY, env.WebEnv.BLOCK_KEY)
}

func Save(w http.ResponseWriter, id uint64, token string) error {
	secureData, err := cookie.Encode("auth_token",
		map[string]any{
			"user_id": id,
			"token":   token,
		})

	if err != nil {
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "auth_token",
		Value: secureData,
		Path:  "/",
		HttpOnly: true,
	})

	return nil
}
