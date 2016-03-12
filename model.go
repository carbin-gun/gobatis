package main

import (
	"encoding/xml"
	"sync"
)

// SqlNode is the node that can execute query
// which can be query,

type GoBatis struct {
	XMLName xml.Name  `xml:"gobatis"`
	Nodes   []RawNode `xml:",any"`
}

type RawNode struct {
	XMLName    xml.Name
	Id         string `xml:"id,attr"`
	ResultType string `xml:"resultType,attr"`
	ResultMap  string `xml:"resultMap,attr"`
	Data       string `xml:",chardata"`
}

////SqlNode is a marker interface .the following definition struct
//// represents sql operations of `select`,`insert`,`delete` and `update` respectively
//type SqlNodeType interface {
//}
//
//type SelectNode struct{}
//
//func (node SelectNode) String() string {
//	return "select"
//}
//
//type InsertNode struct{}
//
//func (node InsertNode) String() string {
//	return "insert"
//}
//
//type DeleteNode struct{}
//
//func (node DeleteNode) String() string {
//	return "delete"
//}
//
//type UpdateNode struct{}
//
//func (node UpdateNode) String() string {
//	return "update"
//}

var (
	ExecutionNode = map[string]struct{}{
		"select": struct{}{},
		"delete": struct{}{},
		"update": struct{}{},
		"insert": struct{}{},
	}
	SupportSqlType = map[string]struct{}{
		"select": struct{}{},
		"delete": struct{}{},
		"update": struct{}{},
		"insert": struct{}{},
		//如上和ExcutionNode一致
		"typeAlias": struct{}{},
		"resultMap": struct{}{},
	}
)

type NodeInfo struct {
	SqlNodeType string
	ResultType  string
	ResultMap   string
	Sql         string
	Id          string
}

type SqlRegistry struct {
	Map    map[string]NodeInfo
	Mutext sync.RWMutex
}

//type QueryNode struct {
//	Id         string `xml:"id"`
//	ResultType string `xml:"resultType"`
//}
//
//type InsertNode struct {
//	Id         string `xml:"id"`
//	ResultType string `xml:"resultType"`
//}
//type UpdateNode struct {
//	Id         string `xml:"id"`
//	ResultType string `xml:"resultType"`
//}
//
//type DeleteNode struct {
//	Id         string `xml:"id"`
//	ResultType string `xml:"resultType"`
//}
