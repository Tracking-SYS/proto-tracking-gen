## INSTALLATION

### HOW TO RELEASE MODULE

1. Clone proto repository which is located at same directory level with this repository
```
git clone git@github.com:Tracking-SYS/proto-tracking-gen.git
```
2. Generate go files in **proto-tracking** repository with command
```
./build.sh
```

3. Tag version and publish
```
./release.sh
```

### HOW TO USE MODULE IN PROJECT
1. Setup **GOPRIVATE** to bypass Go Mod proxy site
```
go env -w GOPRIVATE=github.com/Tracking-SYS
```

2. Execute command to download latest version
```
go get -u github.com/Tracking-SYS/proto-tracking-gen
```
