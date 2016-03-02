# Include common Makefile code.
BASE_IMAGE_NAME = go
VERSIONS = 1.4 1.5 1.6
OPENSHIFT_NAMESPACES = 1.4 1.5 1.6

# Include common Makefile code.
include hack/common.mk
