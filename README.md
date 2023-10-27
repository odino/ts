# ts

A small utility to generate unix timestamps from your command line.

```cli
$ ts 
1698418446

$ ts +30s
1698418476

$ ts -h
Usage of ts:
 
  ts         # get current timestamp
  ts +2h     # get timestamp 2h in the future
  ts -30s    # get timestamp 30s in the past
 
 Supported formats https://pkg.go.dev/time#ParseDuration

```

Grab a [release](https://github.com/odino/ts/releases) and have fun!