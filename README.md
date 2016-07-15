# gophersay

```
$ go get github.com/adamryman/gophersay
$ gophersay
 ------------------------
 Concurrency is not parallelism
 ------------------------
   \
    \
     \   ,_---~~~~~----._
  _,,_,*^____      _____``*g*\"*,
 / __/ /'     ^.  /      \ ^@q   f
[  @f | @))    |  | @))   l  0 _/
 \`/   \~____ / __ \_____/    \
  |           _l__l_           I
  }          [______]           I
  ]            | | |            |
  ]             ~ ~             |
  |                            |
   |                           |
```

`gophersay` is `cowsay` with Go Proverbs

Included is `gophersay.sh` which is the bash equivalent of this go program as a bash script.

I recommend you add this to the end of your `.bash_profile`. \ʕ◔ϖ◔ʔ/

## Uses

### Command line args

Say whatever you want!
```
$ gophersay gosay is a better name, wish I would have thought of that
 ------------------------
gosay is a better name, wish I would have thought of that
 ------------------------
   \
    \
     \   ,_---~~~~~----._
  _,,_,*^____      _____``*g*\"*,
 / __/ /'     ^.  /      \ ^@q   f
[  @f | @))    |  | @))   l  0 _/
 \`/   \~____ / __ \_____/    \
  |           _l__l_           I
  }          [______]           I
  ]            | | |            |
  ]             ~ ~             |
  |                            |
   |                           |

```

### Stdin

```
$ tail main.go | gophersay
 ------------------------
		// Seed rand for psudo random numbers
		rand.Seed(time.Now().UnixNano())

		saying = sayings[rand.Intn(len(sayings))]
	}

	fmt.Println(" ------------------------")
	fmt.Println(saying)
	fmt.Print(gopherArt)
}
 ------------------------
   \
    \
     \   ,_---~~~~~----._
  _,,_,*^____      _____``*g*\"*,
 / __/ /'     ^.  /      \ ^@q   f
[  @f | @))    |  | @))   l  0 _/
 \`/   \~____ / __ \_____/    \
  |           _l__l_           I
  }          [______]           I
  ]            | | |            |
  ]             ~ ~             |
  |                            |
   |                           |

```

## Generate art

```
# From the gophersay directory
$ go generate
```
## Credits

Go Proverbs http://go-proverbs.github.io/

Ascii Art  https://gist.github.com/belbomemo/b5e7dad10fa567a5fe8a
