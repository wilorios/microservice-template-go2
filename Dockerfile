FROM golang:latest
LABEL maintainer="Wilson Rios wilorios@gmail.com"
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN cd /app/cmd/service && go build
WORKDIR /app/cmd/service
CMD [ "./service" ]
