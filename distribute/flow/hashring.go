package flow

import (
	"fmt"
	"strconv"
	"strings"
)

var GlobalHashRing = HashRing{make(map[uint32]Node), make([]uint32, 0, 10)}

type Node struct {
	ip       string
	port     string
	vid      int32
	hashcode uint32
}

type HashRing struct {
	Nodes     map[uint32]Node
	HashSlice []uint32
}

func (n Node) GetIp() string {
	return n.ip
}

func (n Node) GetPort() string {
	return n.port
}

func (n Node) GetVid() int32 {
	return n.vid
}

func (n Node) GetHashcode() uint32 {
	return n.hashcode
}

// AddNodesOfLocalHost 每个程序最开始调用，把自己的节点加入本地的哈希环中，在每次增加节点的时候，把新增节点信息发送给etcd
func (h *HashRing) AddNodesOfLocalHost(ip string, port string, count int) {
	for i := 0; i < count; i++ {
		h.AddNodeOfLocalHost(ip, port, int32(i))
	}
}

// AddNodeOfLocalHost 在本地哈希环中增加一个新的节点，该节点是本地的虚拟节点
func (h *HashRing) AddNodeOfLocalHost(ip string, port string, vid int32) uint32 {
	hashcode := hashKey(ip, port, vid)
	node := Node{ip, port, vid, hashcode}
	h.Nodes[hashcode] = node
	h.HashSlice = insert(h.HashSlice, hashcode)
	return hashcode
}

// AddNodeOfEtcd 将从etcd监听到的节点增加到本地的哈希环中
func (h *HashRing) AddNodeOfEtcd(key string, value string) {
	ip, port, vid := transformToNode(key)
	hashCode, _ := strconv.Atoi(value)
	if _, isOk := h.Nodes[uint32(hashCode)]; !isOk {
		h.AddNodeWithHashcode(ip, port, vid, uint32(hashCode))
	} else {
		//fmt.Println("exist")
	}
}

// AddNodeWithHashcode 在哈希环中增加一个新的节点，参数是从etcd解析出来的
func (h *HashRing) AddNodeWithHashcode(ip string, port string, vid int32, hashcode uint32) {
	node := Node{ip, port, vid, hashcode}
	h.Nodes[hashcode] = node
	h.HashSlice = insert(h.HashSlice, hashcode)
}

// GetNodeByHashcode 输入一个hash值，返回在hash环中的下一个节点
func (h *HashRing) GetNodeByHashcode(hashcode uint32) Node {
	for i := 0; i < len(h.HashSlice); i++ {
		if h.HashSlice[i] >= hashcode {
			return h.Nodes[h.HashSlice[i]]
		}
	}
	return h.Nodes[h.HashSlice[0]]
}

// GetNextNodeByHashcode 输入一个hash值，返回在hash环该点中的下一个节点
func (h *HashRing) GetNextNodeByHashcode(hashcode uint32) Node {
	for i := 0; i < len(h.HashSlice); i++ {
		if h.HashSlice[i] > hashcode {
			return h.Nodes[h.HashSlice[i]]
		}
	}
	return h.Nodes[h.HashSlice[0]]
}

// Print 输出哈希环的一些节点信息，包括ip和端口以及vid
func (h *HashRing) Print() {
	fmt.Println("the count of Nodes is ", len(h.HashSlice))
	for i := 0; i < len(h.HashSlice); i++ {
		fmt.Println("<", h.HashSlice[i], " ", h.Nodes[h.HashSlice[i]].ip, h.Nodes[h.HashSlice[i]].vid, ">")
	}
}

// GetLocalVidList 获取本地的vid list
func (h *HashRing) GetLocalVidList(ip string) []int32 {
	var res []int32
	for _, node := range h.Nodes {
		if node.ip == ip {
			res = append(res, node.vid)
		}
	}
	return res
}

//将ip,port,vid整合一起取hash值
func hashKey(ip string, port string, vid int32) uint32 {
	return CalcHash(ip + port + strconv.Itoa(int(vid)))
}

//将一个元素添加到一个升序的uint32切片中，同时保证添加后的切片还是升序的。
func insert(slice []uint32, num uint32) []uint32 {
	slice = append(slice, num)
	if len(slice) == 1 {
		return slice
	}
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > num {
			copy(slice[i+1:], slice[i:])
			slice[i] = num
			break
		}
	}
	return slice
}

//将etcd的获得的string解析为节点的具体数据
func transformToNode(str string) (string, string, int32) {
	res := strings.Split(str, "/")
	vid, _ := strconv.Atoi(res[3])
	return res[1], res[2], int32(vid)
}
