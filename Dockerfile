FROM golang:alpine3.13 as builder
ARG TERRAFORM_VERSION=0.13.6
ARG HADOLINT_VERSION=2.1.0
ARG SHELLCHECK_VERSION=0.7.1
WORKDIR /
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
COPY Dockerfile .
COPY scripts/*.sh .
RUN cp -r /go/bin/* /root/bin
RUN hadolint Dockerfile && shellcheck -- *.sh

FROM alpine:3.13
WORKDIR /
COPY --from=builder /root/bin/* /usr/bin/
COPY scripts/*.sh /usr/bin/
