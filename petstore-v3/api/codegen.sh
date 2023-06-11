#!/bin/bash

oapi-codegen --config=types.cfg.yaml ../petstore.yaml
oapi-codegen --config=server.cfg.yaml ../petstore.yaml

