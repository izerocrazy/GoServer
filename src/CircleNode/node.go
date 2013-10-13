package CircleNode

import (
    "fmt"
)

type Node struct {
    NodeID          int
    FrontNode       *Node
    BehindNode      *Node

    MyCircle        *NodeCircle
}

func (n *Node) InitNode(nodeID int, circle *NodeCircle) {
    n.NodeID = nodeID
    n.MyCircle = circle
}

//============

type NodeCircle struct {
    HeadNode    *Node
    Count       int
    NodeMap     map[int]string
}

func (c *NodeCircle) InitCircle() {
    c.Count = 0
    c.NodeMap = make(map[int]string)
}

func (c *NodeCircle) AddNode(nodeID int, str string){
    var node Node
    node.InitNode(nodeID, c)

    if c.Count == 0 {
        c.HeadNode = &node
        node.FrontNode = &node
        node.BehindNode = &node
    } else {
        frontNode := c.FindFrontNode(nodeID)
        behindNode := c.FindBehindNode(nodeID)

        frontNode.BehindNode = &node
        behindNode.FrontNode = &node
        node.FrontNode = frontNode
        node.BehindNode = behindNode

        if c.HeadNode == node.FrontNode && node.NodeID < c.HeadNode.NodeID {
            c.HeadNode = &node
        }
    }

    c.NodeMap[nodeID] = str

    c.Count = c.Count + 1
}

func (c *NodeCircle) FindFrontNode(nodeID int) *Node {
    if c.Count == 0 {
        return nil
    }

    node := c.HeadNode
    for {
        fmt.Printf("FindFrontNode :%d, %d, %d \n", node.NodeID, node.BehindNode.NodeID, nodeID);
        if node.NodeID == node.BehindNode.NodeID {
            return node
        }

        // last one
        if node.NodeID > node.FrontNode.NodeID && node.NodeID < nodeID {
            return node.FrontNode
        }

        // first one
        if node.NodeID < node.BehindNode.NodeID && node.NodeID > nodeID && node.BehindNode.NodeID > nodeID{
            return node
        }

        if node.NodeID > node.BehindNode.NodeID && node.NodeID > nodeID && node.BehindNode.NodeID < nodeID {
            return node
        }

        node = node.FrontNode
        if node == c.HeadNode {
            return nil
        }
    }
}

func (c *NodeCircle) FindBehindNode(nodeID int) *Node {
    if c.Count == 0 {
        return nil
    }

    node := c.HeadNode
    for {
        fmt.Printf("FindBehindNode: %d, %d, %d \n", node.NodeID, node.FrontNode.NodeID, nodeID);
        if node.NodeID == node.FrontNode.NodeID {
            return node
        }

        // last one
        if node.NodeID > node.FrontNode.NodeID && node.NodeID < nodeID && node.FrontNode.NodeID < nodeID {
            return node
        }

        // first one
        if node.NodeID < node.BehindNode.NodeID && node.NodeID > nodeID {
            return node.BehindNode
        }

        // normal one
        if node.NodeID < node.FrontNode.NodeID && node.NodeID < nodeID && node.FrontNode.NodeID > nodeID{
            return node
        }

        node = node.FrontNode
        if node == c.HeadNode {
            return nil
        }
    }
}

func (c *NodeCircle) Show() {
    if c.Count == 0 {
        fmt.Printf("this circle is empty \n")
        return
    }

    node := c.HeadNode
    for {
        fmt.Printf("%d\n", node.NodeID)

        node = node.FrontNode
        if node == c.HeadNode {
            return 
        }
    }
}


/*func main() {
    var circle NodeCircle
    circle.InitCircle()
    circle.Show()

    //var node1 Node
    circle.AddNode(1, "Node")
    circle.Show()
    fmt.Printf("++++\n")

    circle.AddNode(3, "Node")
    circle.Show()
    fmt.Printf("++++\n")

    circle.AddNode(4, "Node")
    circle.Show()
    fmt.Printf("++++\n")

    circle.AddNode(2, "Node")
    circle.Show()
    fmt.Printf("++++\n")
}*/
