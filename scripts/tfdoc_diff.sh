#! /bin/sh

if ! command -v terraform-docs > /dev/null 2>&1
then
    echo "Error: terraform-docs is not in PATH!"
    exit 1
fi

if [ -z "$1" ]; then
  cat << EOF
Usage: $0 <path> 
For recursive terraform-docs diff on <path>.
EOF
else
  DOCDIFF=$(find "$1" -type f -iname README.md -not -path '*/.*' \
  -exec sh -c 'i="$1"; terraform-docs markdown --sort-by-required=true ${i%/README.md} | diff -q $i -' _ {} \;)
  if [ "$(echo "$DOCDIFF" | grep -c differ)" -gt 0 ]; then
    echo "Differences have been found:"
    echo "$DOCDIFF"
    exit 1
  fi
fi
