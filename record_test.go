package bsonrecord

import (
	"fmt"
	"testing"
	"time"
)

type TestDocument struct {
	Field       string `bson:"field"`
	NestedField struct {
		Field string `bson:"field"`
	} `bson:"nested_field"`
	Time time.Time `bson:"time"`
}

func TestNoChangesInCurrentModel_ReturnsEmptyDiff(t *testing.T) {
	str := &TestDocument{}
	trackable := New(str)
	diff := trackable.Diff()
	if len(diff) > 0 {
		t.Errorf("there should be no difference, but got: %v", len(diff))
	}
}

func TestChangeCurrentModel_ReturnsDiff(t *testing.T) {
	str := &TestDocument{}
	trackable := New(str)
	str.Field = "changed field"
	diff := trackable.Diff()
	if len(diff) != 1 {
		t.Fatalf("should be difference, but it was not")
	}

	if diff["field"] != "changed field" {
		t.Fatalf("should be new value 'changed field', but got '%s'", diff["field"])
	}
}

func TestShouldAcceptNestedFields(t *testing.T) {
	str := &TestDocument{}
	trackable := New(str)
	str.NestedField.Field = "changed field"
	diff := trackable.Diff()
	if len(diff) != 1 {
		t.Fatalf("should be difference, but it was not")
	}

	if diff["nested_field.field"] != "changed field" {
		t.Fatalf("should be new value 'changed field', but got '%s'", diff["nested_field.field"])
	}
}

func ExampleRecord_Diff() {
	yourDocument := &struct {
		*Record `bson:"-"`

		StringField string `bson:"string_field"`
		IntField    int    `bson:"int_field"`
	}{}

	yourDocument.Record = New(yourDocument)

	yourDocument.IntField = 10
	yourDocument.StringField = "hello world"

	diff := yourDocument.Diff()
	fmt.Println(diff["string_field"])
	fmt.Println(diff["int_field"])
	// Output: hello world
	// 10
}
