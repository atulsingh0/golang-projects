## IPv4 Validator

Program accepts one or more IPv4 (sperated by comma ",") and validate them.

```
$ go run 03-ipv4-validator/main.go 127.0.0.1
       128.0.0.1 : true
```

- Or

```
$ go run 03-ipv4-validator/main.go "28.0.0.1,192.189.0., 128.0.0.2,111.111.222.222"
        28.0.0.1 : true
      192.189.0. : false
       128.0.0.2 : true
 111.111.222.222 : true
```