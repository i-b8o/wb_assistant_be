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
	GOOS=linux GOARCH=amd64 go build -o c -v services/mail/server/*
	scp c root@$(SEND_MAIL_IP):/root/c
	rm c