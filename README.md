## 使用模板初始化步骤如下

> Author	Alex Xiang
>
> Date	2024.07.04

1. 新建项目

   ```go
   go mod init <project_name>
   ```

2. 全局替换 `singapore` -> `<project_name>`

3. 项目初始化

   ```go
   go mod tidy
   go mod vendor
   ```

4. 试运行

   ```go
   go run ./src/server/router/main.go
   ```

5. 编译打包

   ```bash
   ./build.sh server router local
   ```

