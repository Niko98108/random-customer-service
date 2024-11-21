# Use the official Go1.22 image as a base
FROM golang:1.22-alpine AS builder

# Install dependencies for Go and Git
RUN apk add --update --no-cache  \
    git

# Set destination for COPY
WORKDIR /app

#setup Git
COPY . ./ 

# Download Go modules
RUN go mod download

# RUN go mod download
RUN go build -o /opt/dev/random-customer-service cmd/main.go 
RUN ls /opt/dev/random-customer-service


# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS random-customer-service

#USER nonroot:nonroot
# Set the working directory to the root directory path
WORKDIR /opt/dev

# Copy over the binary built from the previous stage
COPY --from=builder /opt/dev /opt/dev

#RUN chown -R nonroot /opt/cleveri

EXPOSE 3001

ENTRYPOINT ["/opt/dev/random-customer-service"]