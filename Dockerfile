FROM --platform=$BUILDPLATFORM golang:1.16-alpine3.14 as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM

ARG TARGETARCH
ARG TARGETOS

WORKDIR /app

COPY ./ ./

RUN echo "Here is the build platform ${BUILDPLATFORM}, here is the target platform ${TARGETPLATFORM}"
RUN echo "Here is the target OS ${TARGETOS}, here is the target arch ${TARGETARCH}"

ENV GIN_MODE=release
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o _output/ping .


FROM alpine:3.14.0

WORKDIR /app

COPY --from=builder /app/_output/ping ./

ENTRYPOINT [ "/app/ping" ]