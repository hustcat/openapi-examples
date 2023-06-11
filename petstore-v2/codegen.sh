#!/bin/bash

# gen server code
swagger generate server -f petstore.json --main-package=petstore-server
