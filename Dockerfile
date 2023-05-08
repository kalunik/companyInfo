FROM golang:1.20.3-alpine AS binary
LABEL authors="kalunik"
WORKDIR /go/src/github.com/kalunik/companyInfo/
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o apiCompany cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=binary /go/src/github.com/kalunik/companyInfo/apiCompany ./
CMD ["./apiCompany"]