swag:
	swag init -g cmd/app/main.go

mock:
	 mockgen -source=internal/service/service.go -destination=internal/service/mocks/mock.go