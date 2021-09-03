go build task_generate.go

SET CGO_ENABLED=0
SET GOOS=darwin3
SET GOARCH=amd64
go build task_generate.go

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build task_generate.go


