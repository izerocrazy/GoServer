### GoLang 中关于 IP TCP UDP 等扩展知识###

#### 基础知识 ####

网络存在了很多年了，也有很多种。虽然已经定义了 OSI 模型，但最流行的网络协议栈是 TCP/IP 协议栈，IP 提供了第三层的 OSI 网络协议栈，TCP 和 UDP 提供了第四层。他们与 OSI 模型的大致关系如下：

TCP/IP     | OSI       
-----------|-----------
application| OSI 5 - 7
TCP/UDP    | OSI 4
IP         | OSI 3
h/w interface | OSI 1- 2

**IP层** 提供的是无连接的不可靠传输系统，**TCP** 构建在IP层之上，它是一个面向链接的协议，这个连接依赖于主机的端口号，但它依然是不可靠的，存在丢包的可能，**UDP** 则简单些，它依然是无连接不可靠的，但它也依赖端口号。对IP层来说，它自己的包头支持数据验证，在包头包含了源地址和目标地址，另外在由路由传导到因特网的时候，它还复杂将大包分解为小包，并在另一头（由因特网到路由）组合成大包。比起TCP和UDP而言，它的一个特点是不支持端口。

#### IP ####

IP 层支持的源地址和目标地址包括两种——IPv4 IPv6。这种统称为 IP 地址。IPv4 地址是一个32位无符号整数，拥有2^32个地址，它通常使用.符号分割成4个十进制数，如：127.0.0.1。其中又分为了两个部分，网段地址和网内地址。按照国际准则，网段地址只取前面一个的，称之为A类地址，而C类地址，当然是取了前三个，最后一个是网内地址，这个网络则只能容纳 2^8 台设备。为了标记出那些是网段地址，那些是网内地址，我们使用了『网络掩码』，它可以是 255.255.252.0，事实上，网段地址是可以指定到多少位的，而不是简单的根据「.」来取几位完成的。IPv6 地址的基本内容是类似的，它是由于 IPv4 没法满足需求而出现的。

在 Golang 中 net 包包含了这里面的一些相关的类型和函数用于编程。

	// 自动支持IPv4 IPv6
	type IP []byte			// IP 类型，也就是32位无符号整数
	func(i * IP) String()	// 它将返回 i 的字符串形式
	
	func ParseIP(string) IP	// 依据常见的字符串写法 127.0.0.1 返回一个 IP 类型的值
	
	//掩码略
	
	type IPAddr {
		IP IP				// 这样写大丈夫？
	}
	
	// 从功能上来说，它和 ParseIP 一样，net 是'ip','ip4','ip6'中的一个，而 addr 则可以是一个因特尔网址
	func ResolveIPAddr(net, addr string) (*IPAddr, os.Error)
	
#### 端口 ####
通过 IP 定位主机，但主机是可以提供多个服务。对于 TCP UDP 协议，它们是通过端口号来区分，一个服务都至少有一个端口。这里也存在一些标准端口，telnet 服务使用23的 TCP 协议，DNS 使用53的 TCP 或 UDP 服务。对于 Linux 而言，这些通常都写在 /etc/services 文件中，而 Golang 则提供 func LookupPort() 函数来做相同的事情。这些小于128的端口号的服务，由IETF标准化，统称位IETF服务。

#### TCP ####
它就不做过多的背景解释，见代码块吧。

	// TCPAddr 这一套基本和 IP 类似
	type TCPAddr struct {
		IP 		IP
		Port	int
	}
	
	func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error)
	
	// 找到主机+端口后，就是该建立和服务的连接了
	type TCPConn struct {
		…		//
	}
	// 它是全双工通讯的 Go 类型
	func (c *TCPConn) Write(b []byte) (n int, err os.Error)
	func (c *TCPConn) Read(b []byte) (n int, err os.Error)
	
	// 在客户端建立 TCPConn 的函数
	// net 同上，laddr 是本地地址，通常设置为 nil ,raddr 既为连接的服务器地址
	func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err os.Error)
	
	// 在服务端上通过 TCPListener 处理客户端的请求
	func ListenTCP(net string, laddr *TCPAddr) (l *TCPListener, err os.Error)
	func (l *TCPListener) Accept() (c Conn, err os.Error)
	// 通过 TCPListener 得到了 Conn，它和 TCPConn 都为ReadWriter，全双工
	
	// TCP 的其他常见的控制函数
	func (c *TCPConn) SetTimeOut(nsec int64) os.Error   // 设置超时	func (c *TCPConn) SetKeepAlive(keepalive bool) os.Error  // 保持存活状态（这个应用的场景？）
	
上面函数基本都是网络函数，所以都必须，注意是 **必须** 进行错误处理。	
#### UDP ####
UDP 和 TCP 函数格式差距并不大，主要调用的函数如下

	func ResolveUDPAddr(net, addr string)(*UDPAddr, os.Error)
	func DialUDP(net string, laddr, raddr *UDPAddr) (c *UDPConn, err os.Error)
	func ListenUDP(net string, laddr *UDPAddr)(c * UDPListener, err os.Error)
	func (c *UDPConn) ReadFromUDP(b []byte)(n int, addr *UDPAddr, err os.Error)
	func (c *UDPConn) WriteFromUDP(b []byte, addr *UPDAddr) (n int, err os.Error)
	
从这些函数基本能看出来，Conn 是一个接口，TCPConn 和 UDPConn 实现了这个接口，但是UDPConn 不使用 Write Read
接口来完成数据传输。实际上，在 Net 包中提供了更为简洁的，纯粹用接口来实现 TCP 和 UDP 的函数和类型，见下。

#### net 包中的接口 ####
Conn 接口和它相关的函数

	// 客户端
	func Dial(net, laddr, raddr string)(c Conn, err os.Error)
	// 服务端
	func Listen(net, laddr string)(l Listener, err os.Error)
	func (l Listener) Accept() (c Conn, err os.Error)
	
这里面非常有意思的是，其参数不再是指针了，这点和C++中的概念还是很接近的，一个接口相对于子类来说，像是一个指针

#### 其他　####
两点，一个是C中常见的 select，由于 goroutine 的原因，所以在 golang 中应该没有实现它（？）。二是 IP 协议也实现了。