#### GoLang 中关于服务端和客户端的接口设计 ####

在网络编程中，存在服务端和客户端的设计是非常之多的（除此之外，还可以点对点，以及过滤器。后者类似与单向的中间层）。以下是我关于服务端和客户端的基本接口设计以及实现，使用的是 TCP 协议，还进行了序列化处理——GoLang 中似乎是必须进行序列化处理。

在设计过程中，我还在想对于一个进程而言，它们是可以作为『模块』处理的，换句话可以自己跑自己的。相对『模块』这个概念，还有一个『挂件』——它就必须依赖于主进程。这两者我也会同样设计其接口。

### **Server** ###
对一个 TCP Server，必须存储的是它自己的 IP 和 Port，在 GoLang 中 net.TCPAddr 包含了两者，当然，还有其他的必要数据，如端口的监听对象 net.TCPListen；以及不必要的数据，如对每一个存在的连接。具体代码见下

	type Server interface {
		Init(string)
		RunServer()
	}
	
	type EncodeServer struct {
		ServerListenInfo 	string		// 保存了我们常见的 IP 和 Port 格式：10.20.156.123:8080
		ServerTCPAddr 		*net.TCPAddr
		ServerListen		*net.TCPListener
	}
	
	// 完成对端口的绑定和监听逻辑
	func (es *EncodeServer) InitServer(str string)	// 处理客户端发上来连接服务端的请求
	func (es* EncodeServer) RunServer()
	
如你所见，此处只是完成部分，主要是对服务器的初始化和处理客户端的连接请求。对于一个完成的服务器而言，还缺少处理客户端具体请求的逻辑部分——这部分逻辑将会主要围绕 net.Conn 这个结构体，并且对于客户端和服务端而言，它都扮演的一样的角色，所以我的想法是尽可能的让服务端和客户端使用同样的代码——至少，暂不在服务端表现吧。

### **Client** ###
TCP Client 会记住连接的 Server 的 IP 和 Port，另外，此处我将 net.Conn 直接放在此中了，需要想个办法去实现上面提到的那个想法，也就是客户端和服务端使用同一段代码。

	type Client interface {
		Init()
		SendData(v interface{}) error
	}
	
	type EncodeClient struct {
		ClientTCPInfo 	string
		ClientConn		net.Conn
		ClientError	error
	}
	
	func (ec *EncodeClient) Init()
	func (ec *EncodeClient) SendData(v interface{}) error
	
虽然在结构体中包含了 net.Conn ，但在接口设计中却没有出现 Write 和 Read 对应的接口。**此处尚待补足。**