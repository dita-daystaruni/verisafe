package app

import (
	"context"
	"fmt"

	"github.com/dita-daystaruni/verisafe/internal/configs"
	"github.com/dromara/carbon/v2"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Server struct {
	*gin.Engine
	*configs.Config
	*pgx.Conn
}

func NewServer() (*Server, error) {
	cfg, err := configs.LoadConfig()
	if err != nil {
		return nil, err
	}

	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DatabaseConfig.DatabaseUser,
		cfg.DatabaseConfig.DatabasePassword,
		cfg.DatabaseConfig.DatabaseHost,
		cfg.DatabaseConfig.DatabasePort,
		cfg.DatabaseConfig.DatabaseName,
	))
	if err != nil {
		return nil, err
	}

	return &Server{server, cfg, conn}, nil
}

// Runs the server
func (s *Server) RunServer() {
	carbon.SetDefault(carbon.Default{
		Layout:       carbon.ISO8601ZuluMicroLayout,
		Timezone:     carbon.UTC,
		WeekStartsAt: carbon.Sunday,
		Locale:       "en", // value range: translate file name in the lang directory, excluding file suffix
	})
	RegisterHandlers(s)
	s.Run(fmt.Sprintf(":%d", s.AppConfig.Port))
}
