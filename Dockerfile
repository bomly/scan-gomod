FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux  go build -o ./scan ./cmd

FROM golang:1.24-alpine AS release

WORKDIR /

COPY --from=build /app/scan .

ENTRYPOINT ["/scan"]
