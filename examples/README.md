# Examples

This directory contains an example project which has multiple enums defined in the same configuration file. For verbosity, the configuration file is written in both JSON and YAML formats:

- [enum_config.json](./enum_config.json)
- [enum_config.yaml](./enum_config.yaml)

These files demonstrate how to perform the same configuration in both JSON and YAML formats.

The [enums directory](./enums) contains:
- `role.go` - The generated code for the `Role` enum
- `role_ext.go` - Some manually written code to extend the `Role` enum with domain specific methods
- `account_status.go` - The generated code for the `AccountStatus` enum
- `account_status_ext.go` - Some manually written code to extend the `AccountStatus` enum with domain specific methods

The `role.go` and `account_status.go` files can be regenerated using any of the following commands (from the [projects root directory](..)):

```sh
# Using the JSON configuration file
go run ./cmd/goenums/ ./examples/config/enums.json ./examples/enums/

# Using the YAML configuration file
go run ./cmd/goenums/ ./examples/config/enums.yaml ./examples/enums/

# Using the output_path defined in the YAML configuration file
go run ./cmd/goenums/ ./examples/config/enums.yaml

# Using the go:generate directive in the `multiple-enums/main.go` file
go generate ./examples/...
```

The `main.go` file contains the go:generate directive to regenerate the `role.go` and `account_status.go` files. The `role_ext.go` and `account_status_ext.go` files are manually written and are not regenerated.

The `main.go` file demonstrates how to use the `Role` and `AccountStatus` enums outside the `auth` package.

The extension files (`role_ext.go` and `account_status_ext.go`) demonstrate how to extend the enums from within the `auth` package.