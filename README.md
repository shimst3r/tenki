# tenki 天気

> 天気 (tenki), Japanese for _weather conditions_

tenki is a command-line client for [wttr.in](https://wttr.in/) written in Go, using the [Cobra](https://github.com/spf13/cobra) library.

## Usage Examples

```shell
$ ./tenki get --help
Query weather information either for your default location or for a specific one.

Usage:
  tenki get [--location location] [flags]

Flags:
  -h, --help              help for get
      --location string   Query location

Global Flags:
      --language string   Language to generate output in (default "en")
```