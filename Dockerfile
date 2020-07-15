FROM golang:1.14.4

ENV GO111MODULE=on

WORKDIR /Fruitshop

COPY . .

#RUN ["go", "mod", "vendor"]
RUN ["go", "get", "-u", "goa.design/goa/v3/...@v3"]
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
RUN ["go", "build", "/Fruitshop/cmd/fruitshop"]
EXPOSE 8080
ENTRYPOINT ["./fruitshop"]