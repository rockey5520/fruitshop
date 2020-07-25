FROM node:13.13.0 AS ANGULAR_BUILD
RUN npm install -g @angular/cli@latest
COPY frontend /webapp
WORKDIR webapp
RUN npm install && ng build --prod

FROM golang:1.13.1-alpine AS GO_BUILD
COPY server /server
WORKDIR /server
RUN go build -o /go/bin/server

FROM alpine:3.10
WORKDIR app
COPY --from=ANGULAR_BUILD /webapp/dist/fruitshop-ui/* ./webapp/dist/webapp/
COPY --from=GO_BUILD /go/bin/server ./
RUN chmod +x /app/webapp/dist/webapp/photos
RUN ls
CMD ./server