# Extract email and domains from one file.

Simple and super fast

## run

```go
go run . -file [PATH_TO_FILE]
```

## set output path

```go
go run . -file [PATH_TO_FILE] -o [PATH]
```

## set buffer size

Default 200 MB

```go
go run . -file [PATH_TO_FILE] -buffer [SIZE_IN_MB]
```

## Delete duplicates

```go
go run . -clean [PATH_TO_FILE]
```

