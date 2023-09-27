FROM golang AS builder
ARG GOOGLE_VERSION=v1.2.3
RUN mkdir /nonexistent && chmod 777 /nonexistent
USER nobody:nogroup
RUN git clone -b $GOOGLE_VERSION https://github.com/1268/google /go/google && \
    cd /go/google/cmd/play && \
    CGO_ENABLED=0 go build

FROM alpine
ARG UID=1000
ARG GID=1000
COPY --chown=0:0 --chmod=755 --from=builder /go/google/cmd/play/play /usr/local/bin
RUN mkdir /google && \
    chown "$UID:$GID" /google
USER $UID:$GID
WORKDIR /google
ENTRYPOINT ["/usr/local/bin/play"]