package middleware

import (
	"encoding/gob"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type AuthnMiddlewareBuilder struct {
}

func (b AuthnMiddlewareBuilder) Authn() gin.HandlerFunc {
	gob.Register(time.Now())
	return func(ctx *gin.Context) {
		sess := sessions.Default(ctx)
		userID := sess.Get("userID")
		if userID == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		now := time.Now()
		const updatedTimeKey = "updatedTime"
		val := sess.Get(updatedTimeKey)
		lastUpdatedTime, ok := val.(time.Time)
		if !ok || val == nil || now.Sub(lastUpdatedTime) > time.Second*10 {
			sess.Set(updatedTimeKey, now)
			sess.Set("userID", userID)
			err := sess.Save()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
