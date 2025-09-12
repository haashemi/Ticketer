FROM golang:1.25 AS build-stage

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code.
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /mock ./cmd/mock
RUN CGO_ENABLED=0 GOOS=linux go build -o /ticketer ./cmd/ticketer

# Deploy
FROM alpine:latest

WORKDIR /

COPY --from=build-stage /mock /mock
COPY --from=build-stage /ticketer /ticketer

# Run
CMD ["/ticketer"]
