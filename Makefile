
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	#protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/test2/test2.proto
	protoc --proto_path=${GOPATH}/src:. --go_out=. proto/model/*.proto
	protoc --proto_path=${GOPATH}/src:. --micro_out=. proto/rpcapi/*.proto

.PHONY: build
build:
	go build -o goMicroDemoSrv ./srv/server.go
	go build -o goMicroDemoCli ./cli/client.go
	go build -o goMicroDemoApi ./api/api.go

.PHONY: move
move: build
	cp ${GOPATH}/bin/micro ./

.PHONY: tar
tar: move
	tar -czf micro.tar.gz goMicroDemoSrv goMicroDemoCli goMicroDemoApi micro Makefile
	rm goMicroDemoCli
	rm goMicroDemoApi
	rm goMicroDemoSrv
	rm micro

.PHONY: untar
untar:
	tar -xzvf micro.tar.gz
	rm micro.tar.gz

.PHONY: delete
delete:
	-rm micro
	-rm goMicroDemoCli
	-rm goMicroDemoApi
	-rm goMicroDemoSrv
	-rm micro.tar.gz

.PHONY:stop
stop:
	#-pgrep goMicroDemoSrv && killall goMicroDemoSrv
	#-pgrep goMicroDemoCli && killall goMicroDemoCli
	#-pgrep goMicroDemoApi && killall goMicroDemoApi
	#-pgrep micro && killall micro
	pkill goMicroDemoSrv || true
	pkill goMicroDemoCli || true
	pkill goMicroDemoApi || true
	pkill micro || true


