package app

import (
	"context"
	"fmt"
	"time"

	"github.com/dita-daystaruni/verisafe/internal/configs"
	"github.com/dromara/carbon/v2"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	*gin.Engine
	*configs.Config
	*pgxpool.Pool
}

func NewServer() (*Server, error) {
	cfg, err := configs.LoadConfig()
	if err != nil {
		return nil, err
	}

	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	dbConfig, err := pgxpool.ParseConfig(fmt.Sprintf(
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

	dbConfig.MaxConns = 10
	dbConfig.MinConns = 5
	dbConfig.MaxConnLifetime = time.Minute * 30

	connPool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		return nil, err
	}

	
	return &Server{server, cfg, connPool}, nil
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
