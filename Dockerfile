FROM golang:alpine3.13
ARG TERRAFORM_VERSION=0.13.6
ARG HADOLINT_VERSION=2.1.0
ARG SHELLCHECK_VERSION=0.7.1
WORKDIR /
# External tools
RUN wget -q https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
  unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip && rm terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
  mv terraform /usr/bin/terraform
RUN wget -qO hadolint https://github.com/hadolint/hadolint/releases/download/v${HADOLINT_VERSION}/hadolint-Linux-x86_64 && \
  chmod +x hadolint && mv hadolint /usr/bin/hadolint
RUN go get github.com/terraform-docs/terraform-docs@v0.12.0
RUN go get github.com/terraform-linters/tflint@v0.29.1
RUN go get github.com/tfsec/tfsec/cmd/tfsec@v0.40.3
RUN apk add build-base --no-cache
# Frequent cache invalidators
COPY Dockerfile .
RUN hadolint --ignore DL3006 --ignore DL3018 Dockerfile
COPY tfd /tfd
WORKDIR /tfd
RUN export TFD_LOGLEVEL=trace && go test ./test -v -coverpkg=./...
RUN go build . && mv tfd /go/bin/
CMD ["tfd"]
