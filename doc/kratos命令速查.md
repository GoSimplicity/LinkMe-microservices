# kratos命令速查
```bash
kratos new verify-code # 创建项目
cd verify-code
go mod tidy
go get github.com/google/wire/cmd/wire
go generate ./...      # 更新wire依赖
kratos run
kratos proto add api/xxx/xxx.proto # 生成proto文件
kratos proto client api/xxx/xxx.proto # 生成pb代码
kratos proto server api/xxx/xxx.proto -t internal/service # 生成服务代码
protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative xxx.proto # 生成proto文件到当前目录下
```