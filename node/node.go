package node

import "orchestrator-from-scratch/stats"

type Node struct {
	Name            string
	Ip              string
	Cores           int
	Memory          int
	MemoryAllocated int
	Disk            int
	DiskAllocated   int
	Stats           stats.Stats
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
