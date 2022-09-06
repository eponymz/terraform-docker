#! /bin/sh
set -e

if [ -n "${GO_COMMON_USER}" ]; then
  echo "Configuring GO_COMMON_USER: $GO_COMMON_USER"
  touch $HOME/.gitconfig
  git config --global url."https://${GO_COMMON_USER}:${GO_COMMON_PASS}@gitlab.com".insteadOf "https://gitlab.com"
else
  echo "WARNING: GO_COMMON_USER credentials are not set"
fi

go install . 
rm $HOME/.gitconfig
