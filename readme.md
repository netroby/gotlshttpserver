# gotlshttpserver


Golang TLS HTTP Server demo

To generate self signed TLS certificate

```shell
wget -qO- https://rootsh.xyz/gentlscert.sh | sh -s "*.mydomain.com"
```

Replace your domain into *.mydomain.com

Then you can run

```
go get -u github.com/netroby/gotlshttpserver
```

After finished got gotlshttpserver, then you can run

```
./gotlshttpserver
```
