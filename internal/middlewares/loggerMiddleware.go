package middlewares

import (
	"time"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

func LoadLoggerMiddleware(logger *zap.Logger) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			req := c.Request()

			err := next(c)

			fields := []zap.Field{
				zap.String("method", req.Method),
				zap.String("path", req.URL.Path),
				zap.Int("status", c.Response().Status),
				zap.Duration("latency", time.Since(start)),
			}

			if err != nil {
				fields = append(fields, zap.Error(err))
				logger.Debug("request failed", fields...)
			} else {
				logger.Debug("request complete", fields...)
			}

			return err
		}
	}
}
