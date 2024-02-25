generate:
	buf generate
	make generate_go

generate_go:
	protoc-go-inject-tag -input ./go/saturn_sync_protocol/v1/saturnsync.pb.go -remove_tag_comment
	gofmt -w ./go/saturn_sync_protocol/v1/saturnsync.pb.go
