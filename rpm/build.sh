#!/bin/bash
echo "Recreating rpmbuild directory"
rm -rvf /root/rpmbuild/
rpmdev-setuptree
echo "Building SRPM"
rpmbuild --undefine=_disable_source_fetch -bs /project/kubesw/rpm/kubesw.spec
mkdir -p ~/.config
mv /project/kubesw/copr ~/.config/copr
copr-cli build kubesw /root/rpmbuild/SRPMS/*.src.rpm
