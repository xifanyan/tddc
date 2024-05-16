# tddc

```
git clone https://github.com/xifanyan/tddc.git
```

```
go mod init github.com/xifanyan/tddc
```

```
go get github.com/microsoft/kiota-abstractions-go
go get github.com/microsoft/kiota-http-go
go get github.com/microsoft/kiota-serialization-form-go
go get github.com/microsoft/kiota-serialization-json-go
go get github.com/microsoft/kiota-serialization-text-go
go get github.com/microsoft/kiota-serialization-multipart-go
```

```
 .\kiota.exe generate -n github.com/xifanyan/tddc/pkg -d .\totaldiscovery-api-1-0-3.json -l go -o ..\..\workspace\tddc\pkg
 ```