package CircleNode

import (
    "fmt"
    "hash/crc32"
)

type Circle interface {
    Init()

    AddNode(nodeId int)
    ReduceNode(nodeId int)

    GetFrontNode(nodeId int) Node
    GetBehindNode(nodeId int) Node
}

//=======
type CircleModuler struct{
    HeadNode    Node
    Count       int
}

//=======
// Module
func (cm *CircleModuler) Init() {
    fmt.Print("CircleModuler.Init()\n")

    cm.Count = 0
}

func (cm *CircleModuler) Breath(){
    fmt.Print("CircleModuler.Breath()\n")

    cm.Show()

    //var node1 Node
    cm.AddHashNode("1")
    cm.Show()
    fmt.Printf("++++\n")

    cm.AddHashNode("3")
    cm.Show()
    fmt.Printf("++++\n")

    cm.AddNode(4)
    cm.Show()
    fmt.Printf("++++\n")

    cm.AddNode(2)
    cm.Show()
    fmt.Printf("++++\n")
}

func (cm *CircleModuler) Run() {
    fmt.Print("CircleModuler.Run()\n")
}

func (cm *CircleModuler) Stop() {
    fmt.Print("CircleModuler.Stop()\n")
}

func (cm *CircleModuler) IsSelfRun() bool {
    return false
}

func (cm *CircleModuler) Load()  error {
    cm.Init()
    if cm.IsSelfRun() == true {
        cm.Run()
    }

    return nil
}

func (cm *CircleModuler) Unload() error {
    if cm.IsSelfRun() == true {
        cm.Stop()
    }

    return nil
}

//=========
// Circle
func (c *CircleModuler) AddNode(nodeId int){
    var node Node
    node = new (CircleNode)
    node.InitNode(nodeId)

    if c.Count == 0 {
        c.HeadNode = node
        node.SetFrontNode(node)
        node.SetBehindNode(node)
    } else {
        frontNode := c.FindFrontNode(nodeId)
        behindNode := c.FindBehindNode(nodeId)

        frontNode.SetBehindNode(node)
        behindNode.SetFrontNode(node)
        node.SetFrontNode(frontNode)
        node.SetBehindNode(behindNode)

        if c.HeadNode == node.GetFrontNode() && node.GetNodeId() < c.HeadNode.GetNodeId() {
            c.HeadNode = node
        }
    }

    c.Count = c.Count + 1
}

func (c *CircleModuler) AddHashNode(hashStr string) int {
    key := []byte(hashStr)
    var v int
    v = int(crc32.ChecksumIEEE(key))

    c.AddNode(v)

    return v
}

func (c *CircleModuler) FindFrontNode(nodeID int) Node {
    if c.Count == 0 {
        return nil
    }

    node := c.HeadNode
    for {
        fmt.Printf("FindFrontNode :%d, %d, %d \n", node.GetNodeId(), node.GetBehindNode().GetNodeId(), nodeID);
        if node.GetNodeId() == node.GetBehindNode().GetNodeId() {
            return node
        }

        // last one
        if node.GetNodeId() > node.GetFrontNode().GetNodeId() && node.GetNodeId() < nodeID {
            return node.GetFrontNode()
        }

        // first one
        if node.GetNodeId() < node.GetBehindNode().GetNodeId() && node.GetNodeId() > nodeID && node.GetBehindNode().GetNodeId() > nodeID {
            return node
        }

        if node.GetNodeId() > node.GetBehindNode().GetNodeId() && node.GetNodeId() > nodeID && node.GetBehindNode().GetNodeId() < nodeID {
            return node
        }

        node = node.GetFrontNode()
        if node == c.HeadNode {
            return nil
        }
    }
}

func (c *CircleModuler) FindBehindNode(nodeID int) Node {
    if c.Count == 0 {
        return nil
    }

    node := c.HeadNode
    for {
        fmt.Printf("FindBehindNode: %d, %d, %d \n", node.GetNodeId(), node.GetFrontNode().GetNodeId(), nodeID);
        if node.GetNodeId() == node.GetFrontNode().GetNodeId() {
            return node
        }

        // last one
        if node.GetNodeId() > node.GetFrontNode().GetNodeId() && node.GetNodeId() < nodeID && node.GetFrontNode().GetNodeId() < nodeID {
            return node
        }

        // first one
        if node.GetNodeId() < node.GetBehindNode().GetNodeId() && node.GetNodeId() > nodeID {
            return node.GetBehindNode()
        }

        // normal one
        if node.GetNodeId() < node.GetFrontNode().GetNodeId() && node.GetNodeId() < nodeID && node.GetFrontNode().GetNodeId() > nodeID{
            return node
        }

        node = node.GetFrontNode()
        if node == c.HeadNode {
            return nil
        }
    }
}

func (c *CircleModuler) Show() {
    if c.Count == 0 {
        fmt.Printf("this circle is empty \n")
        return
    }

    node := c.HeadNode
    for {
        fmt.Printf("%d\n", node.GetNodeId())

        node = node.GetFrontNode()
        if node == c.HeadNode {
            return 
        }
    }
}


/*func main() {
}*/
