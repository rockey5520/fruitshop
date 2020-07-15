FROM golang:latest

ENV WORK_DIR "$GOPATH/src/github.com/github.com/rockey5520/fruitshop"

RUN mkdir -p $WORK_DIR
ADD . $WORK_DIR
WORKDIR $WORK_DIR

#COPY swagger/swagger.json /opt/goa/swagger/
#COPY swagger/swagger.yaml /opt/goa/swagger/

RUN git clone --branch v1 https://github.com/goadesign/goa.git $GOPATH/src/github.com/goadesign/goa

RUN go get
RUN go build -o fruitshop .
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o fruitshop .

EXPOSE 9090
RUN chmod +x ./fruitshop
CMD ["./fruitshop"]