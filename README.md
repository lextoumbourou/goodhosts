# goodhost

Simple host file management in Go.

<img src="http://www.hangthebankers.com/wp-content/uploads/2013/09/Masks-Rothschild-party1.jpg" width=400><br>
[A Surrealist Parisian Dinner Party chez Madame Rothschild, 1972](http://www.messynessychic.com/2013/08/27/a-surrealist-parisian-dinner-party-chez-madame-rothschild-1972/)

## command-line usage

### list all entries

```
$ goodhost list
```

### check an entry

```
$ goodhost check 127.0.0.1 google.com 
```

### add an entry

```
$ goodhost add 127.0.0.1 google.com
```

### remove a line (coming soon)

```
$ goodhost remove 127.0.0.1 google.com
```
