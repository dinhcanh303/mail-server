module github.com/dinhcanh303/mail-server

go 1.20

require (
	github.com/confluentinc/confluent-kafka-go/v2 v2.3.0
	github.com/go-ldap/ldap/v3 v3.4.6
	github.com/golang-migrate/migrate/v4 v4.16.2
	github.com/golang/glog v1.2.0
	github.com/google/uuid v1.6.0
	github.com/google/wire v0.5.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.18.0
	github.com/ilyakaznacheev/cleanenv v1.5.0
	github.com/kyleconroy/sqlc v1.16.0
	github.com/labstack/echo/v4 v4.11.1
	github.com/meilisearch/meilisearch-go v0.26.0
	github.com/minio/minio-go/v7 v7.0.63
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.18.0
	github.com/rabbitmq/amqp091-go v1.9.0
	github.com/redis/go-redis/v9 v9.3.0
	github.com/rs/cors v1.10.1
	github.com/samber/lo v1.38.1
	github.com/sirupsen/logrus v1.9.3
	github.com/sqlc-dev/pqtype v0.3.0
	go.mongodb.org/mongo-driver v1.13.1
	go.uber.org/automaxprocs v1.5.3
	go.uber.org/zap v1.19.1
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9
	google.golang.org/appengine v1.6.8
	google.golang.org/genproto/googleapis/api v0.0.0-20240304212257-790db918fca8
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.3.0
	google.golang.org/protobuf v1.33.0
)

require (
	cloud.google.com/go/compute v1.23.4 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/Azure/go-ntlmssp v0.0.0-20221128193559-754e69321358 // indirect
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/elastic/elastic-transport-go/v8 v8.3.0 // indirect
	github.com/go-asn1-ber/asn1-ber v1.5.5 // indirect
	github.com/go-jose/go-jose/v3 v3.0.0 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20231201235250-de7065d80cb9 // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/nats-io/nkeys v0.4.7 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.45.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/valyala/fasthttp v1.37.1-0.20220607072126-8a320890c08d // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.uber.org/multierr v1.7.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
	gorm.io/driver/postgres v1.5.7 // indirect
)

require (
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	google.golang.org/grpc v1.62.1
)

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/coreos/go-oidc/v3 v3.7.0
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/elastic/go-elasticsearch/v8 v8.11.1
	github.com/go-redis/redis/v8 v8.11.5
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.1.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hibiken/asynq v0.24.1
	github.com/joho/godotenv v1.5.1
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.2 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/knadh/smtppool v1.1.0
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/lib/pq v1.10.9
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/minio/sha256-simd v1.0.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nats-io/nats.go v1.34.0
	github.com/rs/xid v1.5.0 // indirect
	github.com/stretchr/testify v1.8.4
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/crypto v0.21.0
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/oauth2 v0.18.0
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.4.0
	golang.org/x/tools v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240228224816-df926f6c8641
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/gorm v1.25.7
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)
