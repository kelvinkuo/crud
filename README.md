# crud

[中文](https://github.com/kelvinkuo/crud/blob/main/README_CHINESE.md)
[ENGLISH](https://github.com/kelvinkuo/crud)

## intro
crud is a tool for generating api defined files from datasource.

inspired by [sql2pb](https://github.com/Mikaelemmmm/sql2pb) but code is completly rewritten.

**datasource supported:**
* mysql

**api files supported:**
* proto3
* zero api

## install
```bash
go install github.com/kelvinkuo/crud@latest
```

## usage

```
Usage:
  crud [flags]

Flags:
  -f, --format string        output format support proto3 or zeroapi (default "proto3")
      --source string        datasource example: root:123456@tcp(127.0.0.1:3306)/shop
  -m, --method string        methods separated by "," (default "add,delete,update,info,list,search")
  -c, --ignore_cols string   columns ignored separated by "," (default "create_at,create_time,created_at,update_at,update_time,updated_at")
  -t, --table string         tables separated by "," or "*" for all tables (default "*")
  -s, --service string       service name
      --go_package string    go_package used in proto3 (default the same as service)
  -p, --package string       package used in proto3 (default the same as service)
      --style string         style of field name in proto3 message, goCrud or go_crud (default "go_crud")
  -h, --help                 help for crud
```

## examples

simple
```bash
crud -f zeroapi --source "root:123456@tcp(127.0.0.1:3306)/shop" -s shop > shop.api
```

full
```bash
crud -f proto3 --source "root:123456@tcp(127.0.0.1:3306)/shop" -m "add,delete,update,info,list,search" \
-c "created_at,updated_at,deleted_at" -t * -s shop -go_package shop -package "./shop" --style go_crud > shop_model.proto
```

## todo
- [x] sort messages
- [x] data type convert from db to api file
- [x] db enum type
- [x] clean code of convert module
- [x] filter out time-related fields in tables
- [X] args parser
- [X] project layout
- [X] api file support
- [X] custom proto3 message field style
- [X] fully test
- [X] document
  - [X] comment
  - [X] README

## contact
please feel free to contact me for any question or suggestion of crud

- email: kelvinkuo224@gmail.com
- github: https://github.com/kelvinkuo
