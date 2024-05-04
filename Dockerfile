# Build stage
FROM golang:1.20 AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN GOOS=linux go build -o l2scan-indexer .

# Final stage
FROM ubuntu
RUN apt update && apt install ca-certificates -y
WORKDIR /app
COPY --from=build /app/l2scan-indexer .
CMD ["./l2scan-indexer"]
