# goodhost

Simple hosts file management in Go.

<img src="http://www.hangthebankers.com/wp-content/uploads/2013/09/Masks-Rothschild-party1.jpg" width=400><br>
[A Surrealist Parisian Dinner Party chez Madame Rothschild, 1972](http://www.messynessychic.com/2013/08/27/a-surrealist-parisian-dinner-party-chez-madame-rothschild-1972/)

## Command-line usage

### list all entries

```
$ goodhost list
127.0.0.1 localhost
10.0.0.5 my-home-server xbmc-server
10.0.0.6 my-desktop
```

### check an entry

```
$ goodhost check 10.0.0.5 my-home-server
```

### add an entry

```
$ goodhost add 10.0.0.5 music-server
```

### remove an entry

```
$ goodhost remove 10.0.0.5 music-server
```
