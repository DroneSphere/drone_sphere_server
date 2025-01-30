package product_entity

type ProductDomain string

const (
	ProductDomainDrone  ProductDomain = "0" // 无人机
	ProductDomainGimbal ProductDomain = "1" // 云台
	ProductDomainRC     ProductDomain = "2" // 遥控器
	ProductDomainDock   ProductDomain = "3" // 机场
)

type Type int
type SubType int

const (
	TypeMavic3 Type = 0
)

const (
	SubTypeDefault  SubType = 0 // 默认
	SubTypeAddition SubType = 1 // 附加
)

// DroneBidirectionalMap 用于存储无人机的名称、类型和子类型之间的映射关系
type DroneBidirectionalMap struct {
	nameToType        map[string]Type             // Name -> Type
	nameToSubType     map[string]SubType          // Name -> SubType
	typeToDefaultName map[Type]string             // Type -> Name (SubType 为 0 的默认值)
	typeSubTypeToName map[Type]map[SubType]string // Type + SubType -> Name (SubType 不为 0 的情况)
}

// Add 添加无人机的名称、类型和子类型之间的映射关系
func (bm *DroneBidirectionalMap) Add(name string, t Type, st SubType) {
	// Name -> Type 和 Name -> SubType 的映射
	bm.nameToType[name] = t
	bm.nameToSubType[name] = st

	// Type + SubType -> Name 的映射
	if st == 0 {
		// SubType 为 0 的情况，存储到 typeToDefaultName
		bm.typeToDefaultName[t] = name
	} else {
		// SubType 不为 0 的情况，存储到 typeSubTypeToName
		if _, ok := bm.typeSubTypeToName[t]; !ok {
			bm.typeSubTypeToName[t] = make(map[SubType]string)
		}
		bm.typeSubTypeToName[t][st] = name
	}
}

// GetTypeSubTypeByName 根据名称获取无人机的类型和子类型
func (bm *DroneBidirectionalMap) GetTypeSubTypeByName(name string) (Type, SubType, bool) {
	t, ok1 := bm.nameToType[name]
	st, ok2 := bm.nameToSubType[name]
	return t, st, ok1 && ok2
}

// GetNameByTypeSubType 根据类型和子类型获取无人机的名称
func (bm *DroneBidirectionalMap) GetNameByTypeSubType(t Type, st SubType) (string, bool) {
	if st == 0 {
		// SubType 为 0 的情况，优先从 typeToDefaultName 中查找
		if name, ok := bm.typeToDefaultName[t]; ok {
			return name, true
		}
		return "", false
	} else {
		// SubType 不为 0 的情况，从 typeSubTypeToName 中查找
		if subTypeMap, ok := bm.typeSubTypeToName[t]; ok {
			name, exists := subTypeMap[st]
			return name, exists
		}
		return "", false
	}
}

// ListNames 获取所有无人机的名称
func (bm *DroneBidirectionalMap) ListNames() []string {
	names := make([]string, 0, len(bm.nameToType))
	for name := range bm.nameToType {
		names = append(names, name)
	}
	return names
}

var droneMap DroneBidirectionalMap

func init() {
	droneMap = DroneBidirectionalMap{
		nameToType:        make(map[string]Type),
		nameToSubType:     make(map[string]SubType),
		typeToDefaultName: make(map[Type]string),
		typeSubTypeToName: make(map[Type]map[SubType]string),
	}

	droneMap.Add("Mavic 3E", TypeMavic3, SubTypeDefault)
	droneMap.Add("Mavic 3T", TypeMavic3, SubTypeAddition)
}
