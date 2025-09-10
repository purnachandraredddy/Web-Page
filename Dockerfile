#build stage to buikd the application
FROM golang:1.25-alpine AS builder 

#set the working directory in the container 
WORKDIR /app

#copy the go mod and go sum files to the container 
COPY go.mod go.sum ./ 


#download the dependencies and verify the mod files 
RUN go mod download && go mod verify 

#copy the application code to the container 
#copy the rest of the application code to the container 
COPY . .
#build the application 
RUN go build -o main main.go 
# final stage to run the application by reducing the size of the image 
FROM alpine:latest
#set the working directory 
WORKDIR /app
#copy the just build application to the container only build time dependencies are copied 
COPY --from=builder /app/main .

#expose the port 
EXPOSE 8000
#run the application 
CMD ["./main"]