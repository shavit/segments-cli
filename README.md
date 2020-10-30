# Segments

> Create segments from metadata

## Quick start

```
cmd -p lib/silence_1.xml -max 6000 -ms 2000 -split 400
```

Write to stdout
```
cmd -p lib/silence_1.xml -max 6000 -ms 2000 -split 400 
```

Write to a path
```
cmd -p lib/silence_1.xml -max 6000 -ms 2000 -split 400 -o /tmp/segments.json -s 
```

## Install

Make sure the binary is in your path

```
export PATH=$PATH:$GOPATH/bin
```

Download and build, or install the program
```
# build
# go get -u github.com/shavit/segments-cli/cmd
# rename cmd 
# go build -o $GOPATH/bin/cmd github.com/shavit/segments-cli/cmd

# Install
go install github.com/shavit/segments-cli/cmd
```

## Test
```
go test github.com/shavit/segments-cli/...
```

## Help
```
cmd -h
```
