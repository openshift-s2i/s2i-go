# This image provides an environment for building and running Go applications.

FROM openshift/base-centos7

MAINTAINER SoftwareCollections.org <sclorg@redhat.com>

EXPOSE 8080

ENV GO_VERSION=1.4 \
    GOPATH=$HOME/go \
    GOBIN=$HOME/go/bin \
    SOURCE=$HOME/go/src/main \
    PATH=$PATH:$HOME/go/bin:/usr/local/go/bin

LABEL io.k8s.description="Platform for building and running Go applications" \
      io.k8s.display-name="Go 1.4" \
      io.openshift.expose-services="8080:http" \
      io.openshift.tags="builder,go,go14"

# Install Go toolchain
RUN yum install -y centos-release-scl && \
    INSTALL_PKGS="mercurial" && \
    yum install -y --setopt=tsflags=nodocs --enablerepo=centosplus $INSTALL_PKGS && \
    rpm -V $INSTALL_PKGS && \
    yum clean all -y && \
    rm -rf /opt/app-root/src/.pki && \
    (curl https://storage.googleapis.com/golang/go1.4.3.linux-amd64.tar.gz | \
        tar -zx -C /usr/local)

# Copy the S2I scripts from the specific language image to $STI_SCRIPTS_PATH
COPY ./s2i/ $STI_SCRIPTS_PATH

USER 1001

# Set the default CMD to print the usage of the language image
CMD $STI_SCRIPTS_PATH/usage
