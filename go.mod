module GoWork

go 1.16

require (
	github.com/apache/thrift v0.15.0
	github.com/chrislusf/glow v0.0.0-20181102060906-4c40a2717eee
	github.com/coreos/bbolt v1.3.6 // indirect
	github.com/coreos/etcd v3.3.26+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/psilva261/timsort v1.0.0 // indirect
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20201229170055-e5319fda7802 // indirect
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	go.etcd.io/etcd v3.3.26+incompatible
	go.uber.org/zap v1.19.1 // indirect
	golang.org/x/net v0.0.0-20211007125505-59d4e928ea9d // indirect
	golang.org/x/sys v0.0.0-20211007075335-d3039528d8ac // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	google.golang.org/genproto v0.0.0-20211007155348-82e027067bd4 // indirect
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace github.com/coreos/bbolt v1.3.6 => go.etcd.io/bbolt v1.3.6

replace google.golang.org/grpc v1.41.0 => google.golang.org/grpc v1.26.0
