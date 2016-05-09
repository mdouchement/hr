# Human Readable

Simple binary aimed to humanize the sizes

## Usage

```sh
$ echo 5436537 | hr
5.4M
```

## Installation with Golang

```bash
go get github.com/mdouchement/hr
```

## Compilation (cross-compilation)

- On your current OS

```sh
go build hr.go
```

> According to [Golang system list](https://github.com/golang/go/blob/master/src/go/build/syslist.go)

- GNU/Linux

```sh
GOOS=linux GOARCH=amd64 go build -o hr-Linux-x86_64 hr.go
```

- Mac OSX (darwin)

```sh
GOOS=darwin GOARCH=amd64 go build -o hr-Darwin-x86_64 hr.go
```

- MS Windows

```sh
GOOS=windows GOARCH=amd64 go build -o hr-Windows-x86_64.exe hr.go
```


## License

**MIT**

## Contributing

1. Fork it ( https://github.com/mdouchement/hr/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
