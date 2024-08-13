package callback

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"01-Login/platform/authenticator"
	"01-Login/platform/storage"
)

// Handler for our callback.
func Handler(auth *authenticator.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		if ctx.Query("state") != session.Get("state") {
			ctx.String(http.StatusBadRequest, "Invalid state parameter.")
			return
		}


		// Exchange an authorization code for a token.
		token, err := auth.Exchange(ctx.Request.Context(), ctx.Query("code"))
		if err != nil {
			ctx.String(http.StatusUnauthorized, "Failed to convert an authorization code into a token.")
			return
		}

		idToken, err := auth.VerifyIDToken(ctx.Request.Context(), token)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Failed to verify ID Token.")
			return
		}

		var profile map[string]interface{}
		if err := idToken.Claims(&profile); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		session.Set("access_token", token.AccessToken)
		session.Set("profile", profile)
		if err := session.Save(); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		if sub, exists := profile["sub"]; exists {
			parts := strings.Split(sub.(string), "|")
			profile["key"] = parts[1]
			delete(profile, "sub")
		}
		

		storage.UploadJSON("cs361microservicedata", profile["key"].(string), profile)

		// Sends a post to the data callback, and redirects the user to the usercallback.
		userCallback, ok := session.Get("user_callback").(string)
		if !ok {
			ctx.String(http.StatusInternalServerError, "Failed to retrieve user_callback from session.")
			return
		}

		// dataCallback, ok := session.Get("data_callback").(string)
		// if !ok {
		// 	ctx.String(http.StatusInternalServerError, "Failed to retrieve data_callback from session.")
		// 	return
		// }

		// Clearing session here to get around problems of previous states.
		deleteSessionInfo(session)

		userCallback += "&key=" + profile["key"].(string)
		// profileData, err := json.Marshal(profile)
        // if err != nil {
        //     ctx.String(http.StatusInternalServerError, "Failed to marshal profile data.")
        //     return
        // }

        ctx.Redirect(http.StatusFound, userCallback)
	}
}

func deleteSessionInfo(session sessions.Session) {
	session.Delete("user_callback")
	// session.Delete("data_callback")
	session.Delete("state")
}