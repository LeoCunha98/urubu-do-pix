# syntax=docker/dockerfile:1

FROM golang:1.21.0 as builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
COPY ../cmd/*.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o server

FROM scratch 
COPY --from=builder /app/server /server

# Run
ENTRYPOINT [ "/server" ]