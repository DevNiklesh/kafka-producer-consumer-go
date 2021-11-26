FROM public.ecr.aws/bitnami/golang:1.17 AS consumer
WORKDIR /go/src/app
COPY go.* ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-s' -o consumer .

FROM scratch
COPY --from=consumer /go/src/app/consumer ./
CMD ["/consumer"]