### Module 包
构建了一个简单 Module 接口

### proc 
一个简单的使用 Module 接口的范式例子

### CircleNode 
实现了 Module 的接口，它核心服务是 hash 一致性的算法，比如在服务端和客户端 hash 出一样的值，比如当服务器群组中有新的机器加入，或者有机器退出，依然 hash 出一样的值，它是分布式服务的一个基础

### EncodeNet
中间包含一个客户端和一个服务端包，均实现了 Module 接口

#### 可用上述四个包实现一个简单的 C/S 分布式服务

### ReflectMap
基于放射写的一个包，功能是，传入类名，即可生成一个对象（但比起 Java 等语言来说，它还需要提前主动注册好）

### RESTControl
遵守 REST 规则的一个接口

### HttpRouter
基于 ReflectMap 功能写了一个 Http 路由分发器，其路由规则暂时非常简陋，后续支持（: 符号，其他隐式规则，比如说调用函数），该路由器只接受实现 RESTControl 接口的对象

### yo + App/yo_server
一个山寨 yo app 应用的服务端，采用 MVC 结构，其 C 中对每一个对象都必须实现 RESTControl 接口

### Todo
未完成的一套管理系统


