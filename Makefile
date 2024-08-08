examples ?= cipher \
			fileupload \
			gw \
			mysql \
			server

all: help

.PHONY: $(examples)

##env 是否根据protobuf文件生成*.pb.go文件
GEN_PROTO_GO ?= true
##env 是否根据protobuf文件生成*_grpc.pb.go文件
GEN_PROTO_GO_GRPC ?= true
##env 是否根据protobuf文件生成*_rest.pb.go文件
GEN_PROTO_GO_REST ?= true
##env 是否根据protobuf文件生成*_rest_gw.pb.go文件
GEN_PROTO_GO_REST_GW ?= true
##env 是否根据protobuf文件生成*_ts.pb.go文件
GEN_PROTO_TS ?= false

help: ## 使用帮助
	@echo "Commands:"
	@echo "$$(grep -hE '^\S+:.*##' $(MAKEFILE_LIST) | sed -e 's/:.*##\s*/:/' -e 's/^\(.\+\):\(.*\)/  \\033[35m\1\\033[m:\2/' | column -c2 -t -s :)"
	@echo
	@echo "Envs:"
	@echo "$$(grep -A1 -hE '^##env' Makefile |sed 's/##env//' |sed 'N;s/\n/ =/'|awk -F '=' '{print $$2"|"$$1"default:"$$3}'|sed -e 's/ .*|/|/' -e 's/^\(.\+\)|\(.*\)/  \\033[32m\1\\033[m|\2/'  |column -c2 -t -s '|')"

update: .gitmodules ## 更新本地代码
	git submodule sync
	git submodule foreach --recursive git reset --hard
	git submodule foreach --recursive git clean -fdx
	git submodule init
	git submodule update
	git submodule update --remote
	git submodule foreach  --recursive 'tag="$$(git config -f $$toplevel/.gitmodules submodule.$$name.tag)";[ -n $$tag ] && git reset --hard  $$tag || echo "this module has no tag"'

gen_proto: ## 生成protobuf目录下的协议
	GEN_PROTO_GO=$(GEN_PROTO_GO) GEN_PROTO_GO_GRPC=$(GEN_PROTO_GO_GRPC) GEN_PROTO_GO_REST=$(GEN_PROTO_GO_REST) GEN_PROTO_GO_REST_GW=$(GEN_PROTO_GO_REST_GW) GEN_PROTO_TS=$(GEN_PROTO_TS) /bin/bash scripts/gen_example_proto.sh

run: $(examples) ## 运行example

$(examples):
	ASJARD_CONF_DIR=$(PWD)/$@/conf go run $@/main.go
