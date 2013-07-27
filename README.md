serve
=====

Very minimal HTTP file server in the spirit of python's `SimpleHTTPServer`.

## Install

Assumptions:

  * `GOPATH` is correctly set
  * `GOPATH` is available on your path (optional but might be a good idea)

Go get it.

```
  $ go get github.com/fcarriedo/serve
```

Install it.

```
  $ go install github.com/fcarriedo/serve
```

That's it!

## Usage

If `GOPATH` is set on your path, your should be able to run it directly:

```
  serve
```

The previous should start the server on your current directory `.`, listening
on port `8080`.

Defaults:

  * `port` := `8080`
  * `dir` := `.`

If you want a different configuration, just specify the desired values:

```
  $ serve -dir ~/my/files/dir -p 80
```

You should be able to browser your `~/my/files/dir` files going to your browser
and pointing it to `localhost` on port `80`.

## License

MIT
