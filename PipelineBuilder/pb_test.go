package PipelineBuilder_test

import (
	pb "Mongorm/PipelineBuilder"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestPB(t *testing.T) {
	x := pb.Begin().Sort(bson.M{"Name": -1}).Unwind("arr", true).End()
	t.Log(x)
}
