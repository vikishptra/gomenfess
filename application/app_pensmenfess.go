package application

import (
	"vikishptra/domain_pensfess/controller/menfessapi"
	"vikishptra/domain_pensfess/gateway/withgorm"
	"vikishptra/domain_pensfess/usecase/runregisteruser"
	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/config"
	"vikishptra/shared/infrastructure/logger"
	"vikishptra/shared/infrastructure/server"
	"vikishptra/shared/infrastructure/token"
)

type pensmenfess struct{}

func NewPensmenfess() gogen.Runner {
	return &pensmenfess{}
}

func (pensmenfess) Run() error {

	const appName = "pensmenfess"

	cfg := config.ReadConfig()

	appData := gogen.NewApplicationData(appName)

	log := logger.NewSimpleJSONLogger(appData)

	jwtToken := token.NewJWTToken(cfg.JWTSecretKey)

	datasource := withgorm.NewGateway(log, appData, cfg)

	httpHandler := server.NewGinHTTPHandler(log, cfg.Servers[appName].Address, appData)

	x := menfessapi.NewGinController(log, cfg, jwtToken)
	x.AddUsecase(
		//
		runregisteruser.NewUsecase(datasource),
	)
	x.RegisterRouter(httpHandler.Router)

	httpHandler.RunWithGracefullyShutdown()

	return nil
}
