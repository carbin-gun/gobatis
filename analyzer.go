package main

import (
	"fmt"
	"strings"
)

func AnalyzeSource(root GoBatis) {
	for _, node := range root.Nodes {
		rawNodeType := node.XMLName.Local
		_, ok := SupportSqlType[rawNodeType]
		if !ok {
			msg := fmt.Sprintf("not supported node:%s", rawNodeType)
			panic(msg)
		}
		id := strings.TrimSpace(node.Id)
		//fmt.Println("resultType:", node.ResultType)
		sql := strings.TrimSpace(node.Data)
		info := NodeInfo{
			SqlNodeType: rawNodeType,
			Sql:         sql,
			Id:          id,
			ResultType:  node.ResultType,
			ResultMap:   node.ResultMap,
		}

		Register(id, info)
	}
	fmt.Println("\n######\n analyze over\n#######\n")
	for k, v := range GlobalSqlRegistry.Map {
		if _, ok := ExecutionNode[v.SqlNodeType]; !ok {
			continue
		}
		fmt.Printf("%s:%s,resultType:%s,resultMap:%s\n", k, v.Sql, v.ResultType, v.ResultMap)
	}
}
