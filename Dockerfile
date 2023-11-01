FROM golang:1.21

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY . .

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

# Run
CMD ["/app/server"]