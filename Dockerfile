FROM golang:1.12

ENV GO111MODULE=on

WORKDIR /Fruitshop

COPY . .

#RUN ["go", "mod", "vendor"]
#RUN ["go", "get", "-u", "goa.design/goa/v3/...@v3"]
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
EXPOSE 8080
ENTRYPOINT CompileDaemon -log-prefix=false -build="go build /Fruitshop/cmd/fruitshop" -command="./fruitshop"