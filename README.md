# Terraform-Docker
  A docker image for pipeline validation and application with terraform projects

## TFD
  tfd is an in-house golang project that wraps Terraform validation projects (listed below) to add some helpful features, such as recursive validation with specified exceptions as well as wrapping the Terraform CLI to allow programmatic plans and application of infrastructure resources.

    A Fast and Flexible tool for pipeline validation and application of Terraform.
    
    Usage:
      tfd [command]
    
    Available Commands:
      deploy      deploys a terraform directory
      help        Help about any command
      validate    validates a terraform directory recursively
    
    Flags:
          --config string      config file (default is $HOME/.tfd.yaml)
      -h, --help               help for tfd
      -v, --verbosity string   Verbosity in logging level. E.g. Info, Warn, Debug. (default "Info")
    
    Use "tfd [command] --help" for more information about a command.

### Commands
<details>
  <summary>validate</summary>

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

</details>

<details>
  <summary>deploy</summary>

    This command executes Terraform commands to deploy infrastructure.
    
    Usage:
      tfd deploy [flags]
    
    Flags:
      -a, --action string      Action you wish to execute in the path. (default "plan")
          --auto-apply         Whether running in pipeline or not.
      -h, --help               help for deploy
      -p, --path string        Path to the directory you wish to deploy.
      -w, --workspace string   Workspace/Environment you wish to deploy.
    
    Global Flags:
          --config string      config file (default is $HOME/.tfd.yaml)
      -v, --verbosity string   Verbosity in logging level. E.g. Info, Warn, Debug. (default "Info")

</details>
<br>

### Configuration
tfd uses [viper](https://github.com/spf13/viper) to allow configuration by file (default is $HOME/.tfd.yaml), environmental variables, or Global Flags (cobra and viper bindings).

>All environmental variable **keys** are case sensitive and prefixed by `TFD_`

| Name | Type | Supported Values | Default | Flags | Value Case Sensitive |
|------|:----:|:----------------:|:-------:|:-----:|:--------------------:|
| LOGLEVEL | `string` | trace, debug, warn, info | info | -v --verbosity | false |
| AUTOAPPLY | `bool` | true, false | false | --auto-apply | false |
| ACTION | `string` | init, plan, apply | plan | -a, --action | false |
| PATH | `string` | any valid filepath | "" | -p, --path | false |
| WORKSPACE | `string` | any valid workspace | "" | -w, --workspace | false |
<br>

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
