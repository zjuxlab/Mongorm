package PipelineBuilder

import . "go.mongodb.org/mongo-driver/bson"

type PipelineBuilder A
type PB = PipelineBuilder

func Begin() PB {
	return PB{}
}

func (pb PipelineBuilder) End() A {
	return A(pb)
}

func (pb PipelineBuilder) Sort(idx M) PB {
	return append(pb, M{"$sort": idx})
}

func (pb PipelineBuilder) Lookup(local_field, foreign_col, foreign_field, as string) PB {
	return append(pb, M{"$lookup": M{"from": foreign_col, "localField": local_field, "foreignField": foreign_field, "as": as}})
}

// field: either $xyz or xyz
func (pb PipelineBuilder) Unwind(field string, keep_empty bool) PB {
	if len(field) > 0 && field[0] != '$' {
		field = "$" + field
	}
	return append(pb, M{"$unwind": M{"path": field, "preserveNullAndEmptyArrays": keep_empty}})
}
