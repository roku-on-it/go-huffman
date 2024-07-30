package huffman

type Node struct {
	Left  *Node  `json:"Left,omitempty"`
	Right *Node  `json:"Right,omitempty"`
	Char  string `json:"Char,omitempty"`
	// TODO: Try to remove this field later
	Key   string `json:"Key,omitempty"`
	Value uint64 `json:"Value,omitempty"`
	Bit   string `json:"Bit,omitempty"`
}

func BuildHuffmanTree(freqMap map[string]uint64) Node {

	nodeMap := make(map[string]Node)

	for k, v := range freqMap {
		nodeMap[k] = Node{Char: k, Value: v, Key: k}
	}

	for range len(nodeMap) - 1 {
		s1 := findSmallest(&nodeMap)
		s2 := findSmallest(&nodeMap)

		parent := Node{Left: &s1, Right: &s2, Value: s1.Value + s2.Value, Key: s1.Key + s2.Key}
		nodeMap[parent.Key] = parent
	}

	root := getFirstNode(&nodeMap)

	return root
}

func BuildHuffmanMap(root Node) map[string]string {
	m := make(map[string]string)
	buildHuffmanMap(&m, root, "")
	return m
}

func buildHuffmanMap(m *map[string]string, node Node, bit string) {
	if node.Left == nil && node.Right == nil {
		node.Bit = bit
		(*m)[node.Char] = bit
		return
	}

	buildHuffmanMap(m, *node.Left, bit+"0")
	buildHuffmanMap(m, *node.Right, bit+"1")
}

func findSmallest(m *map[string]Node) Node {
	firstSmallest := getFirstNode(m)

	for _, node := range *m {
		if node.Value == 1 {
			firstSmallest = node
			break
		}

		if node.Value < firstSmallest.Value {
			firstSmallest = node
		}
	}

	delete(*m, firstSmallest.Key)

	return firstSmallest
}

func getFirstNode(nodeMap *map[string]Node) Node {
	var n Node
	for _, node := range *nodeMap {
		n = node
		break
	}

	return n
}
