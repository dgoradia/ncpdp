# NCPDP Script 2017071

#### From base64 encoded string
```go
script := `PE1lc3NhZ2UgRGF0YXR5cGVzVmVyc2lvbj0iMjAxNzA3MTUiIFRyYW5zcG9ydFZlcnNpb249IjIwMTcwNzE1IiBUcmFuc2FjdGlvbkRvbWFpbj0iU0NSSVBUIiBUcmFuc2FjdGlvblZlcnNpb249IjIwMTcwNzE1IiBTdHJ1Y3R1cmVzVmVyc2lvbj0iMjAxNzA3MTUiIEVDTFZlcnNpb249IjIwMTcwNzE1Ij48L01lc3NhZ2U+`

message, err := ncpdp.NewMessage(script)
```
