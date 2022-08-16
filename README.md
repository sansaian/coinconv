# COINCONV
//todo
- [coincov](#coinconv)
    - [Description](#description)
    - [Requirements](#requirements)
    - [Instructions](#instructions)
- [Environment variables](#environment-variables)
- [Resources](#resources)
## Description
Command line utility that converts one currency to another using
[coinmarketcap](https://coinmarketcap.com/api/v1/#section/Introduction) as a data source.

## Requirements
- Go 1.18

## TODOS
- more test keys
- use gomock for tests
## Instructions
### Build
To run build make sure you have docker installed.

```shell
make build
```

### Running Tests
After that use the command:
```shell
make test
```

### Start execution
example execution comand.
```shell
./coinconv 123.45 USD BTC
```

# Environment variables

| env variable name | description                         | default |
|-------------------|-------------------------------------|---------|
| LOGGER_LEVEL      | Logging level for logrus framework. | info    |
| TIMEOUT           | request timeout.                    | 5s      |
| URL               | url for request data.               | none    |
| TOKEN             | token for coinmarketcap.            | none    |

# Resources
| URL                                                                                       | description   |
|-------------------------------------------------------------------------------------------|---------------|
| https://coinmarketcap.com/api/v1/#section/Introduction                                    | coinmarketcap |
| https://github.com/evrone/go-clean-template/tree/8fb159d185a8f00c16eed753f15f9bfeb8d67fb9 | template      |



