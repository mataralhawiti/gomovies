FROM golang:1.18 as build

WORKDIR /gomovies
COPY . .

RUN go mod download
RUN go vet

RUN CGO_ENABLED=0 go build -o /go/bin/gomovies

FROM gcr.io/distroless/static-debian11

COPY --from=build /go/bin/gomovies /
CMD ["/gomovies"]
