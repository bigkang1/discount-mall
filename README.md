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

## 页面展示

以下为新蜂商城 Vue 版本的页面预览：

- 登录页

![](static-files/登录.png)

- 首页

![](static-files/首页.png)

- 商品搜索

![](static-files/商品搜索.png)

- 商品详情页

![](static-files/详情页.png)

- 购物车

![](static-files/购物车.png)

- 生成订单

![](static-files/生成订单.png)

- 地址管理

![](static-files/地址管理.png)

- 订单列表

![](static-files/订单列表.png)

- 订单详情

![](static-files/订单详情.png)

## 感谢

- [newbee-ltd](https://github.com/newbee-ltd)

- [gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin)
