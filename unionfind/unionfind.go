package unionfind

type UnionFind interface {
	Add(interface{}) (bool, error)
	Union(a, b interface{}) (bool, error)
	IsConnected(a, b interface{}) (bool, error)
}
