# goodhost
--
    import "github.com/lextoumbourou/goodhost"


## Usage

#### func  AddEntry

```go
func AddEntry(ip string, host string) error
```

#### func  GetLines

```go
func GetLines(includeComments bool) ([]string, error)
```

#### func  HasEntry

```go
func HasEntry(ip string, host string) (bool, error)
```

#### func  RemoveEntry

```go
func RemoveEntry(ip string, host string) error
```
