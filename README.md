# Go-Micro Boilerplate

Boilerplate code for kickstart micro services project.

## Requirements
- [GoLang](https://golang.org/dl/) version `1.14.0` or higher.
- [Micro server](https://github.com/micro/micro) version `v3.0.0-develop` or higher.
- [Protocol Buffers](https://github.com/protocolbuffers/protobuf/releases) version `3.0.0` or higher.

## Instructions
### Build and run all services
```
micro server
./setups.sh
```
### Rebuild individual service
```
micro kill <service_name>
./<service_name>/service_setup.sh
```
### Generate an service
- New Go Micro service.
```
micro new <service_name>
```
- copy `service_setup.sh` file to service folder