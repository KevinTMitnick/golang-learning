package LinkStack

type Node struct {
	data  interface{}
	pNext *Node
}

func NewStack() *Node {
	return &Node{}
}

type LinkStack interface {
	IsEmpty() bool
	Push(data interface{})
	Pop() interface{}
	Length() int
}

func (n *Node) IsEmpty() bool {
	if n.pNext == nil {
		return true
	} else {
		return false
	}
}

func (n *Node) Push(data interface{}) {
	newnode := &Node{data, nil}
	newnode.pNext = n.pNext
	n.pNext = newnode //头部插入
}

func (n *Node) Pop() interface{} {
	if n.IsEmpty() == true {
		return nil
	}

	value := n.pNext.data   //要弹出的数据
	n.pNext = n.pNext.pNext //删除
	return value
}

func (n *Node) Length() int {
	pnext := n
	length := 0

	for pnext.pNext != nil {
		pnext = pnext.pNext // 节点的循环跳跃
		length++            // 追加
	}
	return length
}
