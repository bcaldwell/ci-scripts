FROM golang:1.13 AS builder

WORKDIR $GOPATH/src/github.com/bcaldwell/ci-scripts

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /ci-scripts .


# Alpine linux with docker installed
FROM docker:19

ENV HELM_VERSION=3.0.3
ENV KUBECTL_VERSION=1.17.0

# install helm and kubectl
RUN apk add --update --no-cache curl ca-certificates && \
    curl -L https://get.helm.sh/helm-v${HELM_VERSION}-linux-amd64.tar.gz | tar xvz && \
    mv linux-amd64/helm /usr/bin/helm && \
    chmod +x /usr/bin/helm && \
    rm -rf linux-amd64 && \
    curl -LO https://storage.googleapis.com/kubernetes-release/release/v${KUBECTL_VERSION}/bin/linux/amd64/kubectl && \
    mv ./kubectl /usr/bin/kubectl && \
    chmod +x /usr/bin/kubectl && \
    apk del curl && \
    rm -f /var/cache/apk/*

COPY --from=builder /ci-scripts /usr/bin/ci-scripts

ENTRYPOINT [ "/bin/sh" ]