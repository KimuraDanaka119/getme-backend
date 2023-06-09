package handler_factory

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"getme-backend/internal/app"
	"getme-backend/internal/app/auth/delivery/http/handlers/login/simple_auth_handler"
	"getme-backend/internal/app/auth/delivery/http/handlers/login/telegram_auth_handler"
	logout_handler "getme-backend/internal/app/auth/delivery/http/handlers/logout"
	"getme-backend/internal/app/auth/delivery/http/handlers/register/simple_register_handler"
	"getme-backend/internal/app/auth/delivery/http/handlers/register/telegram_register_handler"
	"getme-backend/internal/app/offer/delivery/http/offer_handler"
	offer_id_accept_handler "getme-backend/internal/app/offer/delivery/http/offer_id_handler/accept_handler"
	"getme-backend/internal/app/plans/delivery/http/plans_handler"
	plans_task_handler "getme-backend/internal/app/plans/delivery/http/task_handler"
	skills_info_handler "getme-backend/internal/app/skill/delivery/http/skills/info_handler"
	skills_user_handler "getme-backend/internal/app/skill/delivery/http/skills/user_handler"
	"getme-backend/internal/app/task/delivery/http/task_handler"
	"getme-backend/internal/app/task/delivery/http/task_id_handler"
	"getme-backend/internal/app/token/delivery/http/handlers/token_handler"
	user_profile_handler "getme-backend/internal/app/user/delivery/http/handlers/profile"
	user_status_handler "getme-backend/internal/app/user/delivery/http/handlers/status"
	user_by_id_handler "getme-backend/internal/app/user/delivery/http/handlers/user_by_id"
	"getme-backend/internal/microservices/auth/delivery/grpc/client"
)

const (
	AUTH_TG = iota
	AUTH_SIMPLE
	AUTH_TOKEN
	REGISTER_TG
	REGISTER_SIMPLE
	LOGOUT
	SKILL_INFO
	PROFILE
	PROFILE_ID
	USER_SKILLS
	OFFERS
	USER_STATUS
	OFFER_ACCEPT
	TASK_CREATE
	TASK_APPLY
	PLANS
	TASK_GET
	ADMIN
)

type HandlerFactory struct {
	usecaseFactory    UsecaseFactory
	sessionClientConn *grpc.ClientConn
	logger            *logrus.Logger
	urlHandler        *map[string]app.Handler
}

func NewFactory(logger *logrus.Logger, sessionClientConn *grpc.ClientConn, ucFactory UsecaseFactory) *HandlerFactory {
	return &HandlerFactory{
		usecaseFactory:    ucFactory,
		logger:            logger,
		sessionClientConn: sessionClientConn,
	}
}

func (f *HandlerFactory) initAllHandlers() map[int]app.Handler {
	ucUsecase := f.usecaseFactory.GetUserUsecase()
	tokenUsecase := f.usecaseFactory.GetTokenUsecase()
	authUsecase := f.usecaseFactory.GetAuthUsecase()
	skillUsecase := f.usecaseFactory.GetSkillUsecase()
	offersUsecase := f.usecaseFactory.GetOfferUsecase()
	plansUsecase := f.usecaseFactory.GetPlansUsecase()
	taskUsecase := f.usecaseFactory.GetTaskUsecase()

	sClient := client.NewSessionClient(f.sessionClientConn)
	return map[int]app.Handler{
		LOGOUT:          logout_handler.NewLogoutHandler(f.logger, sClient),
		AUTH_TG:         telegram_auth_handler.NewAuthHandler(f.logger, sClient, tokenUsecase),
		REGISTER_TG:     telegram_register_handler.NewRegisterHandler(f.logger, sClient, authUsecase, ucUsecase),
		AUTH_SIMPLE:     simple_auth_handler.NewAuthHandler(f.logger, sClient, authUsecase),
		REGISTER_SIMPLE: simple_register_handler.NewRegisterHandler(f.logger, sClient, ucUsecase, authUsecase),
		AUTH_TOKEN:      token_handler.NewTokenHandler(f.logger, tokenUsecase, sClient),
		SKILL_INFO:      skills_info_handler.NewInfoHandler(f.logger, sClient, skillUsecase),
		PROFILE:         user_profile_handler.NewProfileHandler(f.logger, sClient, ucUsecase),
		USER_SKILLS:     skills_user_handler.NewUserHandler(f.logger, sClient, skillUsecase),
		PROFILE_ID:      user_by_id_handler.NewUserByIDHandler(f.logger, sClient, ucUsecase),
		OFFERS:          offer_handler.NewOfferHandler(f.logger, sClient, offersUsecase),
		USER_STATUS:     user_status_handler.NewStatusHandler(f.logger, sClient, ucUsecase),
		OFFER_ACCEPT:    offer_id_accept_handler.NewAcceptHandler(f.logger, sClient, offersUsecase),
		PLANS:           plans_handler.NewPlanHandler(f.logger, sClient, offersUsecase, plansUsecase),
		TASK_CREATE:     task_handler.NewTaskHandler(f.logger, sClient, taskUsecase),
		TASK_APPLY:      task_id_handler.NewTaskIdHandler(f.logger, sClient, taskUsecase),
		TASK_GET:        plans_task_handler.NewPlanIDTaskHandler(f.logger, sClient, plansUsecase),
	}
}

func (f *HandlerFactory) GetHandleUrls() *map[string]app.Handler {
	if f.urlHandler != nil {
		return f.urlHandler
	}

	hs := f.initAllHandlers()
	f.urlHandler = &map[string]app.Handler{
		//=============auth==============//
		"/logout":                 hs[LOGOUT],
		"/auth/telegram/login":    hs[AUTH_TG],
		"/auth/simple/login":      hs[AUTH_SIMPLE],
		"/auth/telegram/register": hs[REGISTER_TG],
		"/auth/simple/register":   hs[REGISTER_SIMPLE],
		"/auth/token":             hs[AUTH_TOKEN],
		//=============skills=============//
		"/skills":       hs[SKILL_INFO],
		"/skills/users": hs[USER_SKILLS],
		//=============user=============//
		"/user":        hs[PROFILE],
		"/user/:id":    hs[PROFILE_ID],
		"/user/status": hs[USER_STATUS],
		//=============offers=============//
		"/offers":           hs[OFFERS],
		"/offer/:id/accept": hs[OFFER_ACCEPT],
		//=============plans=============//
		"/plans":           hs[PLANS],
		"/plans/:id/task":  hs[TASK_CREATE],
		"/plans/:id/tasks": hs[TASK_GET],
		//=============tasks=============//
		"/tasks/:id/apply": hs[TASK_APPLY],
	}
	return f.urlHandler
}
