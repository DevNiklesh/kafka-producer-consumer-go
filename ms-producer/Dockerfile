FROM public.ecr.aws/bitnami/golang:1.17 AS producer
WORKDIR /go/src/app
COPY go.* ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-s' -o producer .

FROM scratch
COPY --from=producer /go/src/app/producer ./
CMD ["/producer"]