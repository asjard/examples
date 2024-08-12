-include ./Makefile_base

examples ?= cipher fileupload gw mysql server

.PHONY: $(examples)

gen_proto: ## 生成protobuf目录下的协议
	GEN_PROTO_GO=$(GEN_PROTO_GO) GEN_PROTO_GO_GRPC=$(GEN_PROTO_GO_GRPC) GEN_PROTO_GO_REST=$(GEN_PROTO_GO_REST) GEN_PROTO_GO_REST_GW=$(GEN_PROTO_GO_REST_GW) GEN_PROTO_TS=$(GEN_PROTO_TS) /bin/bash scripts/gen_example_proto.sh

run: $(examples) ## 运行example
	docker compose -p asjard up -d

down: ## 停止example
	docker compose -p asjard down

stats: ## 查看状态
	docker compose -p asjard stats

$(examples):
	${MAKE} -C $@ $(MAKECMDGOALS)
