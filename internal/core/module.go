package core

import (
	"github.com/zeusro/system/internal/core/config"
	"github.com/zeusro/system/internal/core/logprovider"
	"github.com/zeusro/system/internal/core/webprovider"
	"go.uber.org/fx"
)

var CoreModule = fx.Options(
	fx.Provide(config.NewFileConfig),
	fx.Provide(logprovider.GetLogger),
	//todo 集成数据库
	// fx.Provide(NewDatabase),
	fx.Provide(webprovider.NewGinEngine),
)
