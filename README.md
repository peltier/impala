Impala
======

A fun attempt at making a bittorrent tracker in Go. This project will move slowly.

### Usage

```
$ ./bin/impala --help
Usage of ./bin/impala:
  -blacklisted-ips="": Blacklisted IPs.
  -max-connections=20: Max number of TCP connections.
  -whitelisted-clients="": Whitelist of client types.
```

### Building

`make`

### Running

`./bin/impala`

### Testing

`go test`