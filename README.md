![](static-files/newbee-mall.png)

### discount-mall 项目

本项目为限时折扣商城后端接口。技术栈为 Go + Gin。


**如果觉得项目还不错的话可以给项目一个 Star 吧。**

### 本地启动

#### 后端项目启动

首先导入 static-files 中的 sql 文件。

```bash
# 克隆项目
git clone https://github.com/newbee-ltd/newbee-mall-api-go

# 使用 go mod 并安装go依赖包
go generate
# 编译 
go build -o server main.go (windows编译命令为go build -o server.exe main.go )
# 运行二进制
./server (windows运行命令为 server.exe)
```

#### 前端项目启动

测试用户名：admin  测试密码：123456

## 感谢

- [gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin)
