# Terraform-Docker
  A docker image for pipeline validation (and eventually application) with terraform projects

## TFLint
  [TFLint](https://github.com/terraform-linters/tflint) can be used for static code analysis to catch syntactical and some provider related errors.

## TFSec
  [TFSec](https://tfsec.dev/) can be used for static code analysis to assist in meeting security standards. It can also be extended with custom rules.

## Formatting
  [Terraform fmt](https://www.terraform.io/docs/cli/commands/fmt.html) can be leveraged to identify terraform files that need to be formatted. Currently, `tffmt_diff.sh` will only report which files need to be formatted.

## Terraform-docs
  [Terraform-docs](https://github.com/terraform-docs/terraform-docs) can generate terraform documentation automatically. We are able to detect documentation drift by comparing existing docs with freshly generated ones.

## Contributing
  This project gets automatically checked by [Hadolint](https://github.com/hadolint/hadolint) for dockerfile standards and [Shellcheck](https://github.com/koalaman/shellcheck) for scripting standards on each docker build.
