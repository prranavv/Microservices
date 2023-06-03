install:
	GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger
swagger:
	GO111MODULE=off swagger generate spec -0 ./swagger.yaml --scan-models