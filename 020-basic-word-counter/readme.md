### Word Counter


Count the words in environment variable "input" -

```go
// setting env var
export input="what is my name"

go run counter.go   // 4

// reset the env var
export input=""
```

if env var "input" is not set, it will count the words in hardcoded string

```go
go run counter.go   // 26
```