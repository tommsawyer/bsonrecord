# BSONRecord
[![GoDoc](https://godoc.org/github.com/tommsawyer/bsonrecord?status.svg)](https://godoc.org/github.com/tommsawyer/bsonrecord)

Helps you to track changes on mongodb document in golang.  
Example usage:  

```go
type PersonDocument struct {
    *bsonrecord.Record `bson:"-"`
    
    Name string `bson:"name"`
    Age  int    `bson:"age"`
}

func NewPerson(name string, age int) *PersonDocument {
    doc := &PersonDocument{
        Name: name,
        Age: age,
    }
    
    doc.Record = bsonrecord.New(doc)
    return doc
}

func main() {
    john := NewPerson("John", 25)
    john.Age = 35
    diff := john.Diff()
    fmt.Println(diff)
}
```
