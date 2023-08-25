# KSD Version Service

KSD Version service provides an API to access the supported versions of KSD 
and its dependencies.

For each version of the operator, there is a file in the `data` directory that 
contains a list of supported components for that version.

## Running locally

```bash
make run
```

To run with tls, make sure that [`mkcert`](https://github.com/FiloSottile/mkcert) is installed, then run

```bash
make run-tls
```
