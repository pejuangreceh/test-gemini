package config

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	DB  *gorm.DB
	Cfg *Config
}

func NewServer(cfg *Config) *Server {
	var (
		dbConnection *gorm.DB

		// redisConnection *redisPackage.Client
	)

	// if cfg.UseDB {
	// 	dbConn := db.Init(cfg)
	// 	// dbAll := db.Use()
	// 	if cfg.Env == "development" {
	// 		dbConnection = dbConn.Debug()
	// 		dbCustomer = dbCust.Debug()
	// 		dbEkyc = dbKyc.Debug()
	// 		dbUmiFact = dbUmi.Debug()
	// 		dbConfig = dbCfg.Debug()
	// 	} else {
	// 		dbConnection = dbConn.Debug()
	// 		dbCustomer = dbCust.Debug()
	// 		dbEkyc = dbKyc.Debug()
	// 		dbUmiFact = dbUmi.Debug()
	// 		dbConfig = dbCfg.Debug()
	// 	}
	// }

	// if cfg.UseRedis {
	// 	redisConnection = redis.Init(cfg)
	// }

	return &Server{
		DB: dbConnection,
		// Redis:      redisConnection,
		// Service:    services.New(cfg),
		Cfg: cfg,
	}
}

func (server *Server) Start(e *echo.Echo) error {
	return e.Start(":" + server.Cfg.Port)
}
