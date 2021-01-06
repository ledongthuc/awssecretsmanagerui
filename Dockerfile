FROM golang:latest as gobuilder
WORKDIR /app
COPY ./server .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./

FROM node:lts-alpine as vuebuilder
WORKDIR /app
COPY ui/package*.json ./
RUN npm install
COPY ./ui/ .
RUN npm run build

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=gobuilder /app/server .
COPY --from=vuebuilder /app/dist ./static
EXPOSE 8081
CMD ["./server"] 
