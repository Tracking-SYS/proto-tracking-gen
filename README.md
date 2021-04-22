## INSTALLATION

### HOW TO RELEASE MODULE

1. Clone proto repository which is located at same directory level with this repository
```
git clone git@github.com:lk153/proto-tracking.git
```
2. Generate go files in proto repository with command
```
make go
```

3. Tag version and publish
```
./release.sh
```



### HOW TO USE MODULE IN PROJECT
1. Setup **GOPRIVATE** to bypass Go Mod proxy site
```
go env -w GOPRIVATE=github.com/lk153/proto-tracking-gen
```

2. Execute command to download latest version
```
go get -u github.com/lk153/proto-tracking-gen
```
