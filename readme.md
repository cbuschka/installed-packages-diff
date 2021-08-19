# pkgdiff -compare packages and versions on server pairs

## Prerequisites

* GNU make
* python >= 3.9
* pipenv

## Usage

### Setup

```bash
make install_deps
```

### Run tests

```bash
make tests
```

### Create a config config.yml

```yaml
groups:
  web:
    servers:
      - username: root
        hostname: web-dev
        excludes:
          - "missing"
      - username: root
        hostname: web-live
```

### Run pkgdiff

```bash
make run
```

## License

Copyright (c) 2021 by [Cornelius Buschka](https://github.com/cbuschka).

[MIT](./license.txt)