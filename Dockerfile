FROM golang:alpine3.13 as builder
ARG TERRAFORM_VERSION=0.13.6
ARG HADOLINT_VERSION=2.1.0
ARG SHELLCHECK_VERSION=0.7.1
WORKDIR /
# External tools
RUN mkdir /root/bin
RUN wget -q https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
  unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip && rm terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
  mv terraform /root/bin/terraform
RUN wget -qO hadolint https://github.com/hadolint/hadolint/releases/download/v${HADOLINT_VERSION}/hadolint-Linux-x86_64 && \
  chmod +x hadolint && mv hadolint /usr/bin/hadolint
RUN wget -q https://github.com/koalaman/shellcheck/releases/download/v${SHELLCHECK_VERSION}/shellcheck-v${SHELLCHECK_VERSION}.linux.x86_64.tar.xz && \
  tar -xvf shellcheck-v${SHELLCHECK_VERSION}.linux.x86_64.tar.xz && mv shellcheck-v${SHELLCHECK_VERSION}/shellcheck /usr/bin/shellcheck
RUN go get github.com/terraform-docs/terraform-docs
RUN go get github.com/terraform-linters/tflint
RUN go get github.com/tfsec/tfsec/cmd/tfsec
RUN apk add build-base --no-cache
# Frequent cache invalidators
COPY Dockerfile .
RUN hadolint --ignore DL3006 --ignore DL3018 Dockerfile
COPY tfd /tfd
WORKDIR /tfd
RUN export PATH=$PATH:/root/bin && export TFD_LOGLEVEL=trace && go test ./test -v -coverpkg=./...
RUN go build . && mv tfd /go/bin/
CMD ["tfd"]
