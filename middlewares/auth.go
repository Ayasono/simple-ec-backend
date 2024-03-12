package middlewares

import (
  "fmt"
  "net/http"
  "time"

  "github.com/Ayasono/simple-kins-backend/utils"
  "github.com/dgrijalva/jwt-go"
  "github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 检查并验证JWT的中间件
func jwtMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    tokenString, err := c.Cookie("jwt") // 假设token保存在cookie中
    if err != nil {
      c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
      return
    }

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
      }
      jwtToken := utils.LoadEnvVariables()
      return []byte(jwtToken.Token), nil
    })

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
      // 检查Token是否过期
      if exp, ok := claims["exp"].(float64); ok {
        // jwt-go库中，exp是以秒为单位的时间戳
        now := time.Now().Unix()
        if int64(exp) < now {
          c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token is expired"})
          return
        }
      }
    } else {
      c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
      return
    }

    c.Next() // Token验证通过，继续处理请求
  }
}
