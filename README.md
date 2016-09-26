opus-go ![opus](https://xiph.org/images/logos/fish_xiph_org.png)
=======

Package opus provides Go bindings for Opus encoder/decoder reference implementation from [Xiph.org](https://www.xiph.org).<br />
All the binding code has automatically been generated with rules defined in [opus.yml](/opus.yml).

### Usage

```
$ brew install opus
$ go get github.com/xlab/opus-go/opus

$ go get github.com/xlab/opus-go/decoder
(optionally if you need a quickstart in decoding)
```

Examples of usage: an implemented [Opus decoder](/decoder) for Go programming language.

### Demo

```
// webm
```

### Rebuilding the package

You will need to get the [cgogen](https://git.io/cgogen) tool installed first.

```
$ git clone https://github.com/xlab/opus-go && cd opus-go
$ make clean
$ make
```
