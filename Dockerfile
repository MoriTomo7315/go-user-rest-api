FROM golang:1.17 as build

# createe app dir on container
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY ./ ./
RUN go mod download

ARG _STAGE
ENV STAGE=${_STAGE}

# Build the binary.
RUN GOOS=linux GOARCH=amd64 go build -mod=readonly -v -o server

EXPOSE 5002
CMD GO_ENV=${STAGE} /app/server