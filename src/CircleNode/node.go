package CircleNode

type Node interface {
    InitNode(nodeId int)

    GetNodeId() int

    GetFrontNode() Node
    SetFrontNode(Node)

    GetBehindNode() Node
    SetBehindNode(Node)
}

type CircleNode struct {
    NodeId          int
    FrontNode       Node
    BehindNode      Node
}

func (n *CircleNode) InitNode(nodeId int) {
    n.NodeId = nodeId
}

func (n *CircleNode) GetNodeId() int {
    return n.NodeId
}

func (n *CircleNode) GetFrontNode() Node {
    return n.FrontNode
}

func (n *CircleNode) GetBehindNode() Node {
    return n.BehindNode
}

func (n *CircleNode) SetFrontNode(fNode Node) {
    n.FrontNode = fNode
}

func (n *CircleNode) SetBehindNode(bNode Node) {
    n.BehindNode = bNode
}
