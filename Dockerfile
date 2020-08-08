FROM node:12.16 AS ANGULAR_BUILD
RUN npm install -g @angular/cli@latest
COPY frontend /frontend
WORKDIR frontend
RUN npm install && ng build --prod

FROM golang:1.14.6-alpine AS GO_BUILD
COPY . /server
WORKDIR /server/server
RUN apk update && apk add build-base
RUN go build -o /go/bin/server/server

FROM alpine:3.10
WORKDIR app
COPY --from=ANGULAR_BUILD /frontend/dist/fruitshop-ui/* ./frontend/dist/fruitshop-ui/
COPY --from=GO_BUILD /go/bin/server/ ./
RUN ls
EXPOSE 8080
CMD ./server    