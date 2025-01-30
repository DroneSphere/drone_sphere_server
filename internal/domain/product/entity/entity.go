package product_entity

import (
	"sort"
	"sync"
)

type Product struct {
	ID          int64         `json:"id"`
	Name        string        `json:"name"`
	Domain      ProductDomain `json:"domain"`
	Type        int           `json:"type"`
	SubType     int           `json:"sub_type"`
	GimbalIndex int           `json:"gimbal_index"`
	SN          string        `json:"sn"`
	Remark      string        `json:"remark"`
}

type Topo struct {
	mu    sync.RWMutex        // 并发锁
	Nodes map[string]*Product `json:"nodes"` // Key=SN
	Edges map[string][]string `json:"edges"` // 邻接表，Key=父设备SN
}

// ApplyUpdate 更新拓扑
func (t *Topo) ApplyUpdate(device *Product, subDevices []*Product) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	// 更新主设备节点（存在则合并字段）
	if _, exists := t.Nodes[device.SN]; !exists {
		t.Nodes[device.SN] = device
	}

	// 按子设备SN排序
	sort.Slice(subDevices, func(i, j int) bool {
		return subDevices[i].SN < subDevices[j].SN
	})

	// 更新子设备节点及邻接表
	children := make([]string, 0, len(subDevices))
	for _, sub := range subDevices {
		// 添加或更新子节点
		if _, exists := t.Nodes[sub.SN]; !exists {
			t.Nodes[sub.SN] = sub
		}
		children = append(children, sub.SN)
	}

	// 更新邻接表（覆盖旧关系）
	t.Edges[device.SN] = children
	return nil
}

// Contains 判断设备是否存在拓扑中
func (t *Topo) Contains(sn string) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	_, exists := t.Nodes[sn]
	return exists
}

// GetChildren 获取直接子设备列表（按Index顺序）
func (t *Topo) GetChildren(sn string) []string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.Edges[sn]
}

// GetAllDescendants 递归获取所有子孙设备（用于嵌套拓扑）
func (t *Topo) GetAllDescendants(sn string) []string {
	t.mu.RLock()
	defer t.mu.RUnlock()

	var descendants []string
	queue := t.Edges[sn]
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		descendants = append(descendants, current)
		queue = append(queue, t.Edges[current]...)
	}
	return descendants
}

// RemoveDevice 删除设备及其关联拓扑
func (t *Topo) RemoveDevice(sn string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	// 删除节点
	delete(t.Nodes, sn)

	// 删除所有出边（该设备作为父节点）
	delete(t.Edges, sn)

	// 删除所有入边（该设备作为子节点）
	for parent, children := range t.Edges {
		filtered := make([]string, 0, len(children))
		for _, child := range children {
			if child != sn {
				filtered = append(filtered, child)
			}
		}
		t.Edges[parent] = filtered
	}
}
