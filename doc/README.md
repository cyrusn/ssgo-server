# Documentation

This documentation are used to show the RESTful api of ssgo, swagger parse `./ssgo/swagger.yaml`

## Installation of Swagger
Swagger is chosen as the first spec to be used.  
See `http-api-specs.md` for details on various specs and tools.

## Prerequisites

```sh
npm i -g bootprint bootprint-openapi
```

## Generation

```sh
bootprint openapi swagger.yaml api/
```
