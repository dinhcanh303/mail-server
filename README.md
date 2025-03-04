# mail-server

An event-driven microservices social application has been written in Golang 
## Technical stack
- Infrastructure
    - PostgreSQL
    - RabbitMQ
    - Docker and Docker-compose
    - Redis
## Design
![mail-server](docs/mail-server.svg)
## Services
No. | Service | URL
--- | ---- | -----
1 | gRPC Gateway | [http://localhost:5000](http://localhost:5000)
2 | Mail Service | [http://localhost:5001](http://localhost:5001)
3 | Web | loading...

## Development
### Install tools
[Docker Desktop for Mac](https://www.docker.com/products/docker-desktop/) <br>
[TablePlus](https://tableplus.com/) or [pgAdmin4](https://www.pgadmin.org/) <br>
[Golang](https://go.dev/) <br>
[Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) <br>
- Mac OS
```bash
brew install golang-migrate
```
- Linux
```bash
curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
apt-get update
apt-get install -y migrate
```
[DB Docs](https://dbdocs.io/docs)
```bash
npm install -g dbdocs
dbdocs login
```
[DBML CLI](https://dbml.dbdiagram.io/cli/#installation)
```bash
npm install -g @dbml/cli
dbml2sql --version
# or if you're using yarn
yarn global add @dbml/cli
dbml2sql --version
```
[SQLC](https://docs.sqlc.dev/en/stable/index.html)
- Mac OS
```bash
brew install sqlc
```
- Linux
```bash
sudo snap install sqlc
```
[Go mock](https://github.com/golang/mock)
```bash
go install github.com/golang/mock/mockgen@v1.6.0
```
### How to generate code
[Generate dependency injection instances with wire](https://github.com/google/wire)
```bash
make wire
```
[Generate code with sqlc](https://docs.sqlc.dev/en/stable/index.html)
```bash
make sqlc
```
[Generate proto using protobuf ](https://github.com/golang/protobuf)
```bash
make proto
```
### Documentation
Generate DB documentation
```bash
make db_docs
```
Access the DB Documentation at db [go-microservice](https://dbdocs.io/dinhcanhng303/go_microservices).Password: 123456789
### How to run
#### Run using Docker
Start docker core include (postgres , redis, rabbitmq, etc)
```bash
make docker-core
```
#### Run 
```bash
make run
```
#### Run test
```bash
make test
```



