## 查看帮助

```bash
make
# 或
make help
```

```bash
Commands:
  gen_proto  生成protobuf目录下的协议
  run        运行example
  down       停止example
  stats      查看状态
  help       使用帮助
  update     更新本地代码

Envs:
  GEN_PROTO_GO           是否根据protobuf文件生成*.pb.go文件 default: true
  GEN_PROTO_GO_GRPC      是否根据protobuf文件生成*_grpc.pb.go文件 default: true
  GEN_PROTO_GO_REST      是否根据protobuf文件生成*_rest.pb.go文件 default: true
  GEN_PROTO_GO_REST_GW   是否根据protobuf文件生成*_rest_gw.pb.go文件 default: true
  GEN_PROTO_TS           是否根据protobuf文件生成*_ts.pb.go文件 default: false
```

## 编译协议

```bash
make gen_proto
```

## 运行实例

```bash
## 只运行server
examples=server make run
## 运行多个实例
examples='cipher fileupload' make run
```
