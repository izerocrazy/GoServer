### Hash 一致性的接口设计 ###

首先，掌握一下 Hash 一致性的背景知识，以下节选自『wiki百科』以及『百度百科』等网站。

**设计目标**是为了解决因特网中的 Hot Spot 问题。个人理解的Hot Spot 问题是指，一个服务进程由于内容等原因，导致其负载远超其他服务进程。同样可以解决这个问题的算法还有一个 CARP ，这个算法基本思路是让多台机器共享一个IP地址，多台服务器提供的内容相同，这对Web服务也还是够用的。

但与 CARP 比起来，Hash 一致性能够提供更多的便利之处。比如说热拔插。总的来说，Hash 一致性能满足/呈现出以下下条件/特性：**分散性**，这个是必须，它的工作环境是多个服务进程中，内容有可能会同时出现在不同的进程中，分散性即是描述这种情况出现的概率；**平衡性**，Hash的结果能够均匀分布到所有的空间/服务进程中；**单调性**，它不仅强调 Hash 结果的唯一性，而且要求满足有新的进程进入，或者有进程退出，并不影响 Hash 结果，它是 Hash 一致性的关键所在；**平衡**，它则是要求退出或进入的服务，能够顺利的迁移到还在运行的其他进程上，保证整个服务群能够提供完整的服务；**负载**，它和分散性有类似，它是指一个服务进程，被不同的用户识别为不同的内容，这也是应该避免/减低的一个参数。

以下是用GoLang实现的一个满足 Hash 一致性中的平衡性和单调性的算法，而其他的则在具体的服务应用中体现。它的核心数据结构为 Circle，核心算法为 Circle 中的左右查找。代码中不包含 Hash 算法，推荐使用CRC32。

#### **HashNode** ####

每个节点知道自己左右两个节点，原本的设计会一个指向 Circle 的指针。这两个点都是可以变动的，比如说你可以知道前X个，后Y个，或者你有全局的节点数据等等。这种设计和具体的应用靠的比较紧凑，但实现的功能大致是：1、Init；2、找到附近的节点。

	type Node interface {
		InitNode (nodeID int)
		GetFrontNode () *Node
		GetBehindNode () *Node
	}
	
	type HashNode struct {
		NodeID 		int
		FrontNode	*HashNode
		BehindNode	*HashNode
	}
	
	func (n *HashNode) InitNode(nodeID int)
	func (n *HashNode) GetFrontNode() *Node
	func (n *HashNode) GetBehindNode() *Node
	
#### **NodeCircle** ####

Circle 数据结构，代码实现如下，其接口 AddNode 是假设传入的 NodeID 自身保证了单调性，这个 NodeID 也就是用户生成的 Hash 值。

	type Circle interface {
		InitCircle()
		
		AddNode(nodeID int)
		ReduceNode(nodeID int)
		GetNode(nodeID int) *Node
	}
	
	type NodeCircle struct {
		HeadNode	*Node
		Count		int
	}
	
	func (c* NodeCircle) InitCircle()
	func (c* NodeCircle) AddNode(nodeID int)
	func (c* NodeCircle) ReduceNode(nodeID int)
	func (c* NodeCircle) GetNode(nodeID int) *Node
	// GetNode 实际调用的就是下面两个中的一个（具体看设计），如果没有直接命中的话
	func (c* NodeCircle) GetFrontNode(nodeID int) *Node
	func (c* NodeCircle) GetBehindNode(nodeID int) *Node
	
#### Hash 数的生成 ####

这个就直接用 GoLang 标准库中的 Hash 就可以了：）	