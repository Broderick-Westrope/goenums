# goenums

goenums is a powerful command-line tool designed to generate type-safe Go enums with rich features such as string representations and JSON (de)serialization. Unlike traditional iota-based enums, goenums provides a more robust and flexible approach to using enums in Go, enhancing code safety and developer convenience.

This repository was forked from [zarldev/goenums](https://github.com/zarldev/goenums).

## Features

- Type Safety: Generates enums as concrete types, preventing misuse and enhancing code safety.
- String Representation: Automatically generates string methods for enums, facilitating debugging and logging.
- JSON Support: Includes methods for JSON marshaling and unmarshaling, ideal for RESTful APIs.
- Configuration Flexibility: Supports both JSON and YAML configurations, automatically detected based on file extension or manually specified with the --format flag.
- Extendability: Allows for easy extension of enum functionality without modifying generated code, supporting a clean separation of generated and custom logic.

## Installation

To install goenums, you can use `go install`:

```sh
go install github.com/broderick-westrope/goenums@latest
```

### Stringer

This tool provides an option to leverage the official [`stringer`](https://pkg.go.dev/golang.org/x/tools/cmd/stringer) tool to automatically generate `String()` methods. This behaviour is enabled by default, however, if you want to use `goenums` without `stringer` you can do so with the `--no-stringer` flag.

If you want to use `stringer` (I recommend you do), then please make sure you have it installed: 

```sh
go install golang.org/x/tools/cmd/stringer
```

## Usage

```sh
goenums [flags] config output
```

- `config`: the path for the configuration file.
  - When the `--format` flag is omitted, the format of the file contents are assumed to match the file extension (eg. `input.json` is assumed to be a JSON file).
- `output`: the path to place the generated files (hence, it should end in a directory).
  - When the output path is specified in the configuration file, this argument may be omitted.
  - When provided, this argument will always be used, even when the configuration file specifies an output path. 

### Flags

- `-h`, `--help`: Show help/usage information.
- `-f`, `--format`: Manually specify the configuration format (json or yaml). When this flag is omitted, the configuration format is derived from the file extension (eg. a `input.json` config file is assumed to be of the JSON format).

### Configuration

Define your enums in a JSON or YAML file. `goenums` will automatically detect the format based on the file extension or use the format specified with the `--format` flag.

JSON Example
```json
{
  "output_path": "./some_output_dir",
  "enums": [
    {
      "package": "validation",
      "type": "Status",
      "values": ["Failed", "Passed", "Skipped", "Scheduled", "Running"]
    }
  ]
}
```

YAML Example
```yaml
output_path: some_output_dir
enums:
- package: validation
  type: Status
  values:
    - Failed
    - Passed
    - Skipped
    - Scheduled
    - Running
```

#### Naming

All configuration items (excluding the `output_path`) are parsed through the [iancoleman/strcase](https://github.com/iancoleman/strcase) Go package. This means that, at the time of writing this, you can indicate new words using capitalisation (eg. camelCase), or by separating words with underscores (_), hyphens (-), dots (.) or spaces. The package will automatically convert these to the appropriate Go naming convention. Please check the package documentation for the most up-to-date information.

### Generating Enums

Run goenums with the path to your configuration file and the desired output directory:

```sh
goenums ./config.json ./output
```

Omit the output argument if you want to use the output path defined in the configuration file:

```sh
goenums ./config.json
```

Use the `--format` flag if your configuration file does not have an appropriate extension (eg. a JSON file not ending in `.json`):

```sh
goenums --format=json ./config.other ./output
```

### Examples

Find more example configurations and generated code in the [examples directory](./examples) of the project repository.

## Contributing

Contributions are welcome! Please feel free to submit pull requests, report bugs, and suggest features.

