FROM node:lts-alpine as vuebuilder
WORKDIR /app
COPY ui/package*.json ./
RUN npm install
COPY ./ui/ .
ENV SERVER=/
RUN npm run export

FROM golang:1.16.2 as gobuilder
WORKDIR /app
COPY ./server .
COPY --from=vuebuilder /app/public ./static
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./aws-secrets-manager-ui .;

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=gobuilder /app/aws-secrets-manager-ui .
EXPOSE 3000
ENV HOST 0.0.0.0
CMD ["./aws-secrets-manager-ui"] 
