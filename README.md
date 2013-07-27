serve
=====

Very minimal HTTP file server in the spirit of python's `SimpleHTTPServer`.

## Install

Assumming you have [Go installed](http://golang.org/doc/install), just Go get
it.

```
  $ go get github.com/fcarriedo/serve
```

Yep... **That's it!** You have a production grade file server ready to run.

## Usage

If `GOPATH` is [set on your path](http://golang.org/doc/code.html#GOPATH)
(which might be a good idea), your should be able to run it directly:

```
  serve
```

The previous should start the server on your current directory `.`, listening
on port `8080`.

Defaults:

  * `port` := `8080`
  * `dir` := `.`

If you want a different configuration, just specify (any of) the desired
values:

```
  $ serve -dir ~/my/files/dir -p 80
```

You should be able to browser your `~/my/files/dir` files going to your browser
and pointing it to `localhost` on port `80`.

## Revelations

Yes, the README is way, way longer than the source code (which is awesome). It
is just a very thin wrapper around Go's FileServer.

## License

MIT
