# goodhosts
--
    import "github.com/lextoumbourou/goodhosts"


## Usage

#### type Hosts

```go
type Hosts struct {
	Path  string
	Lines []HostsLine
}
```


#### func  NewHosts

```go
func NewHosts() Hosts
```

#### func (*Hosts) AddEntry

```go
func (h *Hosts) AddEntry(ip string, host string)
```
Add an entry to the hosts file.

#### func (Hosts) Flush

```go
func (h Hosts) Flush() error
```
Flush any changes made to hosts file.

#### func (Hosts) HasEntry

```go
func (h Hosts) HasEntry(ip string, host string) (bool, error)
```
Return a bool if ip/host combo in hosts file.

#### func (*Hosts) Load

```go
func (h *Hosts) Load() error
```
Load the hosts file into ``l.Lines``.

#### func (*Hosts) RemoveEntry

```go
func (h *Hosts) RemoveEntry(ip string, host string) error
```
Remove an entry from the hosts file.

#### type HostsLine

```go
type HostsLine struct {
	Ip    string
	Hosts []string
	Raw   string
}
```


#### func  NewHostsLine

```go
func NewHostsLine(raw string) HostsLine
```
Create a new instance of ```HostsLine```.

#### func (HostsLine) IsComment

```go
func (l HostsLine) IsComment() bool
```
Return ```true``` if the line is a comment.
