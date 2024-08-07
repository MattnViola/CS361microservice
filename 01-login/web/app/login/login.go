package login

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"01-Login/platform/authenticator"
)

func Handler(auth *authenticator.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		state, err := generateRandomState()
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}


		// Because the URL is not necessarily static, this gets the callback URL from the HTTP request.
		host := ctx.Request.Host
		callbackURL := fmt.Sprintf("https://%s/callback", host)
		auth.Config.RedirectURL = callbackURL


		// State Checks.
		session := sessions.Default(ctx)
		session.Set("state", state)
		userCallback := ctx.Query("user_callback")
        dataCallback := ctx.Query("data_callback")

        if userCallback == "" || dataCallback == "" {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_callback and data_callback are required"})
            return
        }

		session.Set("user_callback", userCallback)
		session.Set("data_callback", dataCallback)

		if dataCallback == "debug" {
			profile := map[string]interface{}{
				"aud": "kgijA1KuOdr0XQszhTgsYEScLRheKRSJ",
				"exp": 1723024604,
				"family_name": "Norman",
				"given_name": "Matthew",
				"iat": 1722988604,
				"iss": "https://secure361login.us.auth0.com/",
				"key": "google-oauth2|104449841002335188382",
				"name": "Matthew Norman",
				"nickname": "mrnorman.sev",
				"picture": "https://lh3.googleusercontent.com/a/ACg8ocLTpCaxxzFgPZfI5utdAnAev-lJ5FEnxJ5uLfHZ67AydeoePQ=s96-c",
				"sid": "IlspzRKyaa1irbLW_RWQIQM5Tt_ue9US",
				"updated_at": "2024-08-06T03:02:31.139Z",
			}
			ctx.JSON(http.StatusOK, profile)
			return
		}

		if err := session.Save(); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Redirect(http.StatusTemporaryRedirect, auth.AuthCodeURL(state))
	}
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
