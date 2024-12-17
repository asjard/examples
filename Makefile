##env 示例列表
examples ?= cipher \
	fileupload \
	gw \
	mysql \
	server \
	readme

export PROJECT_NAME ?= asjard
export BIFROST_DIR ?= ./third_party/bifrost
export GEN_PROTO_GO_VALIDATE ?= true
export GEN_PROTO_GO_ASYNQ ?= true

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


gen_proto: all_env ## 生成protobuf目录下的协议
	$(ALL_ENV) /bin/bash scripts/gen_example_proto.sh

run_dep: ## 运行基础服务,例如数据库
	docker-compose -p $(PROJECT_NAME) up -d

run: run_dep $(examples) ## 本地运行

down: $(examples) ## 停止服务

clean: ## 清理服务
	docker-compose -p $(PROJECT_NAME) down

stats: ## 查看状态
	docker-compose -p $(PROJECT_NAME) stats

$(examples):
	SERVICE_NAME=$@ $(MAKE) -C $@ $(MAKECMDGOALS)
