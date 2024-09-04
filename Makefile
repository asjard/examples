examples ?= cipher \
	fileupload \
	gw \
	mysql \
	server \
	readme

export PROJECT_NAME ?= asjard
export BIFROST_DIR ?= ./third_party/bifrost

-include ./third_party/bifrost/Makefile_base

.PHONY: build run $(examples)

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

run_dep: ## 运行基础服务,例如数据
	docker-compose -p $(PROJECT_NAME) up -d

run: run_dep $(examples) ## 本地运行

down: $(examples) ## 停止服务

clean: ## 清理服务
	docker-compose -p $(PROJECT_NAME) down

stats: ## 查看状态
	docker-compose -p $(PROJECT_NAME) stats

$(examples):
	SERVICE_NAME=$@ $(MAKE) -C $@ $(MAKECMDGOALS)
