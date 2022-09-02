FROM hadolint/hadolint:v2.8.0-alpine AS dockerlint
WORKDIR /app
COPY Dockerfile .
RUN hadolint Dockerfile --ignore DL3018

FROM golang:alpine as gobuilder
RUN go install github.com/terraform-docs/terraform-docs@v0.14.1 && \
    go install github.com/terraform-linters/tflint@v0.29.1 && \
    go install github.com/tfsec/tfsec/cmd/tfsec@v0.40.3

FROM golang:alpine3.16
ARG TERRAFORM_VERSION=0.13.6
ARG SHELLCHECK_VERSION=0.7.1
WORKDIR /
# External tools
RUN wget -q https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
  unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip && rm terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
  mv terraform /usr/bin/terraform
RUN apk add build-base --no-cache
# Frequent cache invalidators
COPY --from=gobuilder /go/bin/terraform-docs /go/bin/terraform-docs
COPY --from=gobuilder /go/bin/tflint /go/bin/tflint
COPY --from=gobuilder /go/bin/tfsec /go/bin/tfsec
COPY tfd /tfd
WORKDIR /tfd
# RUN export TFD_LOGLEVEL=trace && go test ./test -v -coverpkg=./...
RUN go build . && mv tfd /go/bin/
CMD ["tfd"]
