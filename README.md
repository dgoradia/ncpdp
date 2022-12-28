[![codecov](https://codecov.io/gh/dgoradia/ncpdp/branch/main/graph/badge.svg?token=ZNKIEQNZ55)](https://codecov.io/gh/dgoradia/ncpdp)

# NCPDP Script 2017071

## Usage
See tests for more usage examples.

The decoder accepts an `io.Reader` and can decode xml to a `*Message` or `json`.

Decode NewRx (xml) file:
```go
file, err := os.Open("testdata/sample-newrx.xml")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

script := ncpdp.NewDecoder(file)
message, err := script.Decode()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("%+v\n", message.Body.NewRx)
```

Get JSON representation:
```go

message, err := script.ToJson()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("%+v\n", string(message))
```
