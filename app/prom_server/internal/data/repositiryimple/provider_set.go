package repositiryimple

import (
	"github.com/google/wire"

	"prometheus-manager/app/prom_server/internal/data/repositiryimple/alarmhistory"
	"prometheus-manager/app/prom_server/internal/data/repositiryimple/alarmintervene"
	"prometheus-manager/app/prom_server/internal/data/repositiryimple/alarmpage"
	"prometheus-manager/app/prom_server/internal/data/repositiryimple/alarmrealtime"
	"prometheus-manager/app/prom_server/internal/data/repositiryimple/alarmsuppress"
	"prometheus-manager/app/prom_server/internal/data/repositiryimple/alarmupgrade"
	"prometheus-manager/app/prom_server/internal/data/repositiryimple/api"
	"prometheus-manager/app/prom_server/internal/data/repositiryimple/cache"
	"prometheus-manager/app/prom_server/internal/data/repositiryimple/captcha"
	"prometheus-manager/app/prom_server/internal/data/repositiryimple/chatgroup"
	"prometheus-manager/app/prom_server/internal/data/repositiryimple/dataimpl"
	"prometheus-manager/app/prom_server/internal/data/repositiryimple/endpoint"
	"prometheus-manager/app/prom_server/internal/data/repositiryimple/msg"
	"prometheus-manager/app/prom_server/internal/data/repositiryimple/notify"
	"prometheus-manager/app/prom_server/internal/data/repositiryimple/ping"
	"prometheus-manager/app/prom_server/internal/data/repositiryimple/promdict"
	"prometheus-manager/app/prom_server/internal/data/repositiryimple/role"
	"prometheus-manager/app/prom_server/internal/data/repositiryimple/strategy"
	"prometheus-manager/app/prom_server/internal/data/repositiryimple/strategygroup"
	"prometheus-manager/app/prom_server/internal/data/repositiryimple/user"
)

// ProviderSetRepository 注入repository依赖
var ProviderSetRepository = wire.NewSet(
	ping.NewPingRepo,
	promdict.NewPromDictRepo,
	strategy.NewStrategyRepo,
	alarmpage.NewAlarmPageRepo,
	alarmhistory.NewAlarmHistoryRepo,
	strategygroup.NewStrategyGroupRepo,
	role.NewRoleRepo,
	user.NewUserRepo,
	cache.NewCacheRepo,
	endpoint.NewEndpointRepo,
	api.NewApiRepo,
	captcha.NewCaptchaRepo,
	chatgroup.NewChatGroupRepo,
	notify.NewNotifyRepo,
	dataimpl.NewDataRepo,
	alarmrealtime.NewAlarmRealtime,
	alarmintervene.NewAlarmIntervene,
	alarmsuppress.NewAlarmSuppress,
	alarmupgrade.NewAlarmUpgrade,
	msg.NewMsgRepo,
)
