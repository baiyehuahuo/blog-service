# blog-service
Go 语言编程之旅第二章，博客程序

## 开启博客之旅
1. gin.Default 包含 Logger 和 Recovery 中间件
2. Logger 输出请求日志并标准化日志格式
3. Recovery 异常捕获，针对每个请求做 recovery 处理，防止出现 panic
4. gin 的路由也可以编辑测试代码
5. 注册路由时，会携带 Engine 的中间件
6. Engine 实现的 ServeHTTP 李永乐线程池的概念， 防止反复生成上下文对象