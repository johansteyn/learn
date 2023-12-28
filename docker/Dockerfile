# ------------ Builder stage ------------
FROM golang:1.21.5 AS builder

# Set the working directory
WORKDIR /app

# Copy the source code
COPY . .

# Download dependencies
# Not needed for a simple "helloworld" app
#RUN go mod download

# Build the application
# Note that CGO must be disabled otherwise we get "File not found" when running the container
# https://stackoverflow.com/questions/36279253/go-compiled-binary-wont-run-in-an-alpine-docker-container-on-ubuntu-host
RUN CGO_ENABLED=0 go build helloworld.go


# ------------ Final stage ------------
FROM alpine:latest  

WORKDIR /opt

COPY --from=builder /app/helloworld .

# Expose port
EXPOSE 8080

# Output is shown when building with: --progress=plain --no-cache
RUN ls -la

# Run the application when the container starts
CMD ["./helloworld"]
# These also work:
#CMD ./helloworld
#CMD /opt/helloworld

