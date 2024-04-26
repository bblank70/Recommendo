FROM golang:1.22-alpine
 
# Creates an app directory to hold your app’s source code
WORKDIR /app
 
# Copies everything from your root directory into /App

COPY . . 

# Installs Go dependencies
RUN go mod download
 


# Builds builds the binary
RUN go build ./cmd/web/

 
# Tells Docker which network port to listen
EXPOSE 6060


# Specifies the executable command that runs when the container starts
CMD ["./web"]

