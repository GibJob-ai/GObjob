FROM golang:1.16.3-alpine AS go_builder

# Prepare and enter the src folder
WORKDIR /backend

# Download and precompile all third party libraries, ignoring errors (some have broken tests or whatever)
ADD backend/go.mod .
ADD backend/go.sum .
RUN go mod download -x

# Add the sources
ADD ./backend/. .

# Compile! Should only compile our sources since everything else is precompiled
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o main -ldflags "-w" .

FROM node:12.22 AS node_builder
# ADD ./frontend/yarn.lock .
ADD ./frontend/ .
RUN yarn
RUN yarn build

# Put everything in a SMOL container (as low as a few MBs for the go part at least)
# FROM scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=go_builder /backend/main /backend/main
COPY --from=go_builder /backend/schema /backend/schema
COPY --from=node_builder /public /frontend/public
ADD /boot.sh .

EXPOSE 80
EXPOSE 443
CMD ["./boot.sh"]
