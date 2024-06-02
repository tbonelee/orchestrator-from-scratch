package node

type Node struct {
	Name            string
	Ip              string
	Cores           int
	Memory          int
	MemoryAllocated int
	Disk            int
	DiskAllocated   int
	Role            string
	TaskCount       int
}

func NewNode(name string, ip string, role string) *Node {
	return &Node{
		Name: name,
		Ip:   ip,
		Role: role,
	}
}
