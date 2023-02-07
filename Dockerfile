FROM golang:1.20 AS builder

WORKDIR $GOPATH/src/github.com/bcaldwell/ci-scripts

COPY . ./
RUN go build -o /ci-scripts ./cmd/ciscripts/ci-scripts.go


# Alpine linux with docker installed
FROM docker:20
COPY --from=docker/buildx-bin /buildx /usr/libexec/docker/cli-plugins/docker-buildx

ENV HELM_VERSION=3.7.0
ENV KUBECTL_VERSION=1.22.2

# install git, helm and kubectl
RUN apk add --update --no-cache curl ca-certificates git bash && \
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

ENTRYPOINT [ "/bin/bash" ]
