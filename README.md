# Goodhosts

Simple [hosts file](http://en.wikipedia.org/wiki/Hosts_%28file%29) management in Go (golang).

<img src="http://www.hangthebankers.com/wp-content/uploads/2013/09/Masks-Rothschild-party1.jpg" width=400><br>
[A Surrealist Parisian Dinner Party chez Madame Rothschild, 1972](http://www.messynessychic.com/2013/08/27/a-surrealist-parisian-dinner-party-chez-madame-rothschild-1972/)

## Command-Line Usage

### List entries.

```bash
$ goodhosts list
127.0.0.1 localhost
10.0.0.5 my-home-server xbmc-server
10.0.0.6 my-desktop
```

### Check for an entry.

```bash
$ goodhosts check 10.0.0.5 my-home-server
```

### Add an entry.

```bash
$ goodhosts add 10.0.0.5 music-server
```

### Remove an entry.

```bash
$ goodhosts remove 10.0.0.5 music-server
```

## API Usage

### [Docs](API.md)

### Examples

Send Facebook to localhost in hosts file.


```go
func main() {
    h := hosts.NewHosts()

    // Note that nothing will be added to the hosts file until ``Flush`` is called.
    h.AddEntry("127.0.0.1", "facebook.com")

    if err := h.Flush(); err != nil {
        panic(err)
    }
}
```

Remove Facebook from the hosts file

```go
func main() {
    h := hosts.NewHosts()

    // Same deal, yo: call h.Flush() to make permanent.
    h.RemoveEntry("127.0.0.1", "facebook")

    if err := h.Flush(); err != nil {
        panic(err)
    }
}
```
## License

[MIT](LICENSE)
