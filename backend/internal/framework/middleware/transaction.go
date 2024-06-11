package middleware

import (
	"github.com/gin-gonic/gin"
	mydb "myapp/internal/framework/db"
)

func Transaction() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := mydb.DB.Begin()
		defer func() {
			if 400 <= ctx.Writer.Status() {
				db.Rollback()
				return
			}
			db.Commit()
		}()
		ctx.Set("db", db)
		ctx.Next()
	}
}
