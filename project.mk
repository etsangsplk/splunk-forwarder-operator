# Project specific values
OPERATOR_NAME?=$(shell sed -n 's/.*OperatorName .*"\([^"]*\)".*/\1/p' config/config.go)
OPERATOR_NAMESPACE?=$(shell sed -n 's/.*OperatorNamespace .*"\([^"]*\)".*/\1/p' config/config.go)

IMAGE_REGISTRY?=quay.io
IMAGE_REPOSITORY?=$(USER)
IMAGE_NAME?=$(OPERATOR_NAME)
FORWARDER_NAME=splunk-forwarder

FORWARDER_VERSION=8.0.1
FORWARDER_HASH=6db836e2fb9e

VERSION_MAJOR?=0
VERSION_MINOR?=1
