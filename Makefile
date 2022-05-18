gen:
	protoc -I=proto/ --go_out=pb/ proto/*.proto
	protoc --go-grpc_out=pb/ proto/*.proto -I=proto/
	# protoc --dart_out=grpc:fe/welcome/packages/grpc_repository/lib/src/proto/ -Iproto proto/*.proto

git:
	git add .
	git commit -a -m "$m"
	git push -u origin master