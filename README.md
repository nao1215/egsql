[![Build](https://github.com/nao1215/egsql/actions/workflows/build.yml/badge.svg)](https://github.com/nao1215/egsql/actions/workflows/build.yml)
[![UnitTest](https://github.com/nao1215/egsql/actions/workflows/unit_test.yml/badge.svg)](https://github.com/nao1215/egsql/actions/workflows/unit_test.yml)
[![reviewdog](https://github.com/nao1215/egsql/actions/workflows/reviewdog.yml/badge.svg)](https://github.com/nao1215/egsql/actions/workflows/reviewdog.yml)
[![codecov](https://codecov.io/gh/nao1215/egsql/branch/main/graph/badge.svg?token=eXJe7aqoCr)](https://codecov.io/gh/nao1215/egsql)
[![Go Report Card](https://goreportcard.com/badge/github.com/nao1215/egsql)](https://goreportcard.com/report/github.com/nao1215/egsql)
![GitHub](https://img.shields.io/github/license/nao1215/egsql)  
# [WIP] egsql - Database management system to be embedded in the application
## What is egsql
This repository will provide two libraries and one application. They will be developed to satisfy my interest in DB.

- **egsql DBMS**: It is a DBMS similar to sqlite. In other words, it is not a server/client model. I plan to implement eqsql DBMS in pure golang. egsql is based on [bogoDB (Copyright (c) 2019 Shogo Arakawa, MIT LICENSE)](https://github.com/ad-sho-loko/bogoDB).
- **egsql driver**: In golang, several interfaces are defined in "database/sql/driver" to manipulate the DBMS. egsql driver is a driver that satisfies them all.
- **egsql client**: This client is a CLI command that provides the ability to check/modify the DB schema using egsql driver. This CLI command will be used for debugging purposes!


## Origin of the "eg" name
- **e**mbed in **g**olang: It's a DBMS that's embedded in an application.
- **e**ver**g**reen: I hope egsql stands the test of time.
- easy: To Japanese, "eg" and "easy" sound the same. An easy-to-use library is a good library.

