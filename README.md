## 查看帮助

```bash
make
# 或
make help
```

```bash
Commands:
  update     更新本地代码
  gen_proto  生成protobuf目录下的协议
  run_dep    运行基础服务,例如数据
  run        本地运行
  down       停止服务
  clean      清理服务
  stats      查看状态
  help       使用帮助
Envs:
  PROJECT_NAME  项目名称   默认: asjard   当前: asjard
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

## 清理

```bash
make clean
```
