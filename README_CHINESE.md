# crud

[中文](https://github.com/kelvinkuo/crud/blob/main/README_CHINESE.md)
[ENGLISH](https://github.com/kelvinkuo/crud)

## 介绍
crud 是一个从数据源生成CRUD接口定义文件的工具。

受到[sql2pb](https://github.com/Mikaelemmmm/sql2pb)的启发，但是代码完全重新写的，更加模块化和易于扩展。

**数据源支持:**
1. mysql

**接口格式支持:**
1. proto3
2. zero api

## 安装
```bash
go install github.com/kelvinkuo/crud@latest
```

## 使用方法

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

## 示例

一般用法
```bash
crud -f zeroapi --source "root:123456@tcp(127.0.0.1:3306)/shop" -s shop > shop.api
```

完全形式
```bash
crud -f proto3 --source "root:123456@tcp(127.0.0.1:3306)/shop" -m "add,delete,update,info,list,search" \
-c "created_at,updated_at,deleted_at" -t * -s shop -go_package shop -package "./shop" --style go_crud > shop_model.proto
```
