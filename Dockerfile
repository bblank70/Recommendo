FROM golang:1.22-bullseye
 
# Creates an app directory to hold your app’s source code
WORKDIR /App
 
# Copies everything from your root directory into /app
COPY . .

# Installs Go dependencies
RUN go mod download
 
WORKDIR /App/cmd/web
# Builds your app with optional configuration
RUN go build -o main
 
# Tells Docker which network port your container listens on
EXPOSE 6060
 
# Specifies the executable command that runs when the container starts
CMD ["./"]

