package auth

import (
	"net/http"

	godd "github.com/pagongamedev/go-dd"
	goddAPI "github.com/pagongamedev/go-dd/api"
	"golang.org/x/crypto/bcrypt"
)

type requestRegisterCreate struct {
	Username   string `json:"username"     swaggertype:"string"     validate:"required,email"`
	Password   string `json:"password"     swaggertype:"string"     validate:"required,min=8"`
	RePassword string `json:"re_password"  swaggertype:"string"     validate:"required,eqfield=Password"`
	Phone      string `json:"phone"        swaggertype:"string"     validate:"required"`
}

func handlerRegisterCreate() *goddAPI.HTTP {
	api := goddAPI.NewAPIHTTP()

	// api.LifeCycle.ValidateAuth
	// api.LifeCycle.ValidateRole

	api.LifeCycle.ParseRequest(func(context godd.InterfaceContext) (requestMapping interface{}, goddgoddErr *godd.Error) {
		requestMapping = new(requestRegisterCreate)

		if err := context.BodyParser(requestMapping); err != nil {
			return nil, godd.ErrorNew(http.StatusBadRequest, err)
		}
		return requestMapping, nil
	})

	api.LifeCycle.ValidateRequest(func(context godd.InterfaceContext, requestMappingBody interface{}) (requestValidatedBody interface{}, goddErr *godd.Error) {
		request := requestMappingBody.(*requestRegisterCreate)

		// Validate Request
		errors := context.ValidateStruct(*request, godd.Map{"requestRegisterCreate": &requestRegisterCreate{}})
		if errors != nil {
			return nil, errors
		}
		return request, nil
	})

	api.LifeCycle.HandlerLogic(func(context godd.InterfaceContext, requestValidatedBody, requestValidatedParam, requestValidatedQuery interface{}) (code int, responseRaw interface{}, responsePagination *godd.ResponsePagination, goddErr *godd.Error) {
		request := requestValidatedBody.(*requestRegisterCreate)

		svc := context.GetService().(*service)

		response, err := svc.registerCreate(request.Username, request.Password, request.Phone)

		if err != nil {
			return err.Code, nil, nil, err
		}

		return http.StatusOK, response, responsePagination, nil
	})

	// api.LifeCycle.MappingResponse(func(context godd.InterfaceContext, code int, responseRaw interface{}, responsePagination *godd.ResponsePagination) (codeOut int, responseMapping interface{}, responsePaginationOut *godd.ResponsePagination, goddErr *godd.Error) {
	// 	return code, []MappingBookBalance{*responseRaw.(*MappingBookBalance)}, responsePagination, nil
	// })

	return api
}

func (svc *service) registerCreate(username string, password string, phone string) (interface{}, *godd.Error) {

	hashsaltPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)

	response, goddErr := svc.repo.UserRegister(username, string(hashsaltPassword), phone)
	if goddErr != nil {
		return nil, goddErr
	}
	return response, nil
}
