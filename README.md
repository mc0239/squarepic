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

## Building

Instead of `run`, do `build`:

```bash
$ go build .
```

## License

MIT
