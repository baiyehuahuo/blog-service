# blog-service
Go 语言编程之旅第二章，博客程序

## 开启博客之旅
1. gin.Default 包含 Logger 和 Recovery 中间件
2. Logger 输出请求日志并标准化日志格式
3. Recovery 异常捕获，针对每个请求做 recovery 处理，防止出现 panic
4. gin 的路由也可以编辑测试代码
5. 注册路由时，会携带 Engine 的中间件
6. Engine 实现的 ServeHTTP 李永乐线程池的概念， 防止反复生成上下文对象

## 项目设计
1. 目录结构设计
   - configs: 配置文件
   - docs: 文档集合
   - global: 全局变量
   - internal: 内部模块
     - dao: 数据访问层， 所有数据库有关操作在此执行
     - middleware: HTTP 中间件
     - model: 模型层， 用于存放对象
     - routers: 路由相关逻辑处理
     - service: 项目核心业务逻辑
   - pkg: 项目相关模块包
   - storage: 项目生成的临时文件
   - scripts: 各类构建， 安装， 分析等操作的脚本
   - third_party: 第三方资源工具
2. 数据库的表设计， 先设计 SQL 语句， SQL 语句可以通过在线工具转换为 golang 结构体
3. 所有相关结构体封装到 model 包中
4. 从上到下， 先设定接口功能
5. 定制处理接口，先空置，处与 router 目录下
6. 测试路由是否能正常调用

## 编写公共组件
1. 五个模块
   - 错误码标准化
   - 配置管理
   - 数据库连接
   - 日志写入
   - 响应处理
2. 错误码标准化
   - 内部建立一套错误码
   - 通过单独的函数将错误码映射到对应的 HTTP 状态码
3. 配置管理
   - 启动时： 做一些初始化操作
   - 运行时： 监听文件的变更来实现在线更新配置
   - 写在单独的 yaml 文件中, 通过第三方库 viper 读取配置文件
   - 封装配置文件的读取
4. 数据库连接 注意可拓展性和编码规范
5. 日志写入
   - 第三方开源库 lumberjack 核心功能是将日志写入滚动文件
   - 日志所需的文件信息如文件名、行数需要从 runtime 获取
   - 其他主要还是在规范与封装上
6. 响应处理
   - 与错误码标准化相对应
7. **早期做标准化，后期省心省力**

## 生成接口文档
1. Swagger 扫描注解生成 OpenAPI 规范化文档
2. OpenAPI 规范
   - 有关 API 的描述
   - API 可用路径
   - 每个路径上的可用操作
   - 每个操作的输入输出格式
3. Swagger 需要的注解
   - @Summary 摘要
   - @Produce 响应类型
   - @Param 参数格式 从左到右参数为参数名，入参类型，数据类型，是否必填，注释
   - @Success 响应成功 从左到右参数为状态码，入参类型，数据类型，注释
   - @Failure 响应失败 从左到右参数为状态码，入参类型，数据类型，注释
   - @Router 路由 从左到右为路由地址 HTTP 方法
4. 主目录下执行 swag init
5. 路由
   - 在路由中注册 `r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))`
   - 同时要 `import` 对应的 `docs` 目录， 通过 `_` 引入不调用即可
   - 访问 `ip:port/swagger/index.html` 即可看到对应的OpenAPI接口文档
6. 隐藏字段不易于展示， 建议新建一个针对 `Swagger` 的对象