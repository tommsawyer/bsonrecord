// Package bsonrecord helps to track changes in structures to make a query in mongodb.
package bsonrecord

import "github.com/globalsign/mgo/bson"

// Record records all made changes to given document.
type Record struct {
	old     bson.M
	current interface{}
}

// New creates a new Record.
func New(current interface{}) *Record {
	return &Record{
		current: current,
		old:     toBSON(current),
	}
}

// Diff returns diff between old record and changed.
func (d *Record) Diff() bson.M {
	return diff(d.old, toBSON(d.current))
}

func toBSON(in interface{}) bson.M {
	var result bson.M
	bs, _ := bson.Marshal(in)
	bson.Unmarshal(bs, &result)
	return result
}
