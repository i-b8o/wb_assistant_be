ENV	:= $(PWD)/.env

include $(ENV)

gen:
	protoc -I=proto/ --go_out=pb/ proto/*.proto
	protoc --go-grpc_out=pb/ proto/*.proto -I=proto/
	# protoc --dart_out=grpc:fe/welcome/packages/grpc_repository/lib/src/proto/ -Iproto proto/*.proto

git:
	git add .
	git commit -a -m "$m"
	git push -u origin main


swag:
	swag init -g cmd/main.go

mail:
	GOOS=linux GOARCH=amd64 go build -o mail_service -v services/mail/server/*
	scp mail_service root@$(MAIL_IP):/bin/mail_service
	rm mail_service

auth:
	GOOS=linux GOARCH=amd64 go build -o auth_service -v services/auth/server/*
	scp auth_service root@$(AUTH_IP):/bin/auth_service
	rm auth_service

api:
	GOOS=linux GOARCH=amd64 go build -o api_service -v cmd/*
	scp api_service root@$(API_IP):/bin/api_service
	rm api_service