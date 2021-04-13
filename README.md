# Terraform-Docker
  A docker image for pipeline validation and application with terraform projects

## TFD
  tfd is an in-house golang project that wraps other validation projects (listed below) and adds some helpful features, such as recursive validation with specified exceptions.

    A Fast and Flexible tool for pipeline validation and application of Terraform.

    Usage:
      tfd [command]

    Available Commands:
      help        Help about any command
      validate    validates a terraform directory recursively

    Flags:
          --config string      config file (default is $HOME/.tfd.yaml)
      -h, --help               help for tfd
      -v, --verbosity string   Verbosity in logging level. E.g. Info, Warn, Debug. (default "Info")

    Use "tfd [command] --help" for more information about a command.

<details><summary>tfd validate commands</summary>

    You must pass a directory to validate command or call a subcommand.
    This command recursively validates a terraform directory
    using terraform-docs, terraform fmt, tflint, and tfsec

    Usage:
      tfd validate [flags]
      tfd validate [command]

    Available Commands:
      tfdoc       validates a terraform directory recursively with terraform-docs
      tffmt       validates a terraform directory recursively with tffmt
      tflint      validates a terraform directory recursively with tflint
      tfsec       validates a terraform directory recursively with tfsec

    Flags:
      -h, --help   help for validate

    Global Flags:
          --config string      config file (default is $HOME/.tfd.yaml)
      -v, --verbosity string   Verbosity in logging level. E.g. Info, Warn, Debug. (default "Info")

    Use "tfd validate [command] --help" for more information about a command.

</details><br>

### Configuration
  tfd uses [viper](https://github.com/spf13/viper) to allow configuration by file (default is $HOME/.tfd.yaml), environmental variables, or Global Flags (cobra and viper bindings).
  <details><summary>environmental variables</summary><br>
    All environmental variable keys are case sensitive and prefixed by TFD_

    TFD_LOGLEVEL - supported values: trace, debug, warn, info (default: info). The values are case insensitive. Matching flags: -v --verbosity
  </details><br>

## TFLint
  [TFLint](https://github.com/terraform-linters/tflint) can be used for static code analysis to catch syntactical and some provider related errors.

## TFSec
  [TFSec](https://tfsec.dev/) can be used for static code analysis to assist in meeting security standards. It can also be extended with custom rules.

## Formatting
  [Terraform fmt](https://www.terraform.io/docs/cli/commands/fmt.html) can be leveraged to identify terraform files that need to be formatted.

## Terraform-docs
  [Terraform-docs](https://github.com/terraform-docs/terraform-docs) can generate terraform documentation automatically. We are able to detect documentation drift by comparing existing docs with freshly generated ones.

## Contributing
  This project gets automatically checked by [Hadolint](https://github.com/hadolint/hadolint) for dockerfile standards and go test on each docker build.
