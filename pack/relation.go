package pack

import "SimpleTikTok/commom"

func RelationsPtrs(relationsPtrs []*commom.Userinfo) []commom.Userinfo {
	if relationsPtrs != nil {
		var relation = make([]commom.Userinfo, len(relationsPtrs))
		for i, ptr := range relationsPtrs {
			relation[i] = *ptr
		}
		return relation
	}
	return []commom.Userinfo{}
}
