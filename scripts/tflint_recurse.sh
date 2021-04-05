#! /bin/sh

if ! command -v tflint > /dev/null 2>&1
then
    echo "Error: tflint is not in PATH!"
    exit 1
fi

if [ -z "$1" ]; then
  cat << EOF
Usage: $0 <path> 
For recursive tflint on <path>.
EOF
else
  LINT=$(find "$1" -type d -not -path '*/.*' -exec tflint {} \; 2>&1)
  case "$LINT" in 
  (*[![:space:]]*)
    echo "Issues have been found:"
    echo "$LINT"
    exit 1
    ;;
  esac
fi
