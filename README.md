# squarepic

![](https://raw.githubusercontent.com/mc0239/squarepic/master/banner.png?token=AEFGFKK6NMYBW3URQBXCKXC626LAO)

**squarepic** is a simple server that generates a unique picture of squares for
any GET request. It features:

- per-request image size and square count parameters
- configurable defaults

## Requirements

Some version (preferably latest) of Go installed.

## Running
 
Acquire dependency & run:
 
```bash
$ go get github.com/mc0239/logm
```

```bash
$ go run .
```

On first run, program will generate a config file and exit. 
Edit configuration file (if needed) and re-run.

Making a request with non-empty `help` GET parameter displays a help page with possible options:
```
http://localhost:9001/?help=1
```

Example requests:
```
http://localhost:9001/literally-anything!!!
http://localhost:9001/aaaaa?size=100
http://localhost:9001/something/very/cool?size=200&squares=10
http://localhost:9001/bzk?size=10&mirror=1
```

## Building

Instead of `run`, do `build`:

```bash
$ go build .
```

## License

MIT
