# iron/go:dev is the alpine image with the go tools added
FROM iron/go:dev
WORKDIR /app

ENV SRC_DIR=/go/src/github.com/egig/tax_calculator/

ADD . $SRC_DIR

# Build
RUN cd $SRC_DIR; go get; go build -o tax_calculator .; cp tax_calculator /app/

ENTRYPOINT ["./tax_calculator"]