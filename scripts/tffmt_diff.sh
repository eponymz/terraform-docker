#! /bin/sh

if ! command -v terraform > /dev/null 2>&1
then
    echo "Error: terraform is not in PATH!"
    exit 1
fi

if [ -z "$1" ]; then
  cat << EOF
Usage: $0 <path> 
For recursive terraform fmt diff on <path>.
EOF
else
  FMTDIFF=$(terraform fmt -recursive "$1")
  case "$FMTDIFF" in 
  (*[![:space:]]*)
    echo "Differences have been found:"
    echo "$FMTDIFF"
    exit 1
    ;;
  esac
fi
