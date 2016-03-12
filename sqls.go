package main

var (
	GlobalSqlRegistry = SqlRegistry{
		Map: map[string]NodeInfo{},
	}
)

func Register(id string, info NodeInfo) {
	GlobalSqlRegistry.Mutext.Lock()
	defer GlobalSqlRegistry.Mutext.Unlock()
	GlobalSqlRegistry.Map[id] = info
}
