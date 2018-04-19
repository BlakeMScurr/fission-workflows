#!/bin/sh
#
# Build command run by default.
#

set -e

parse() {
    src=$1
    dst=$2
    echo "Parsing ${src} -> ${dst}"
    cat ${src}
    if ! wfcli parse ${src} > ${dst} ; then
        echo "Failed to parse ${SRC_PKG}"
        exit 12
    fi
    echo "Completed parsing ${src} -> ${dst}"
}

echo  "Building with $(wfcli version --client)..."
if [[ -f ${SRC_PKG} ]] ; then
    # Package is a single file
    echo "Package is a single file"
    parse ${SRC_PKG} ${DEPLOY_PKG}
elif [[ -d ${SRC_PKG} ]] ; then
    # Package is a directory
    echo "Package is a directory"
    mkdir -p ${DEPLOY_PKG}
    for wf in ${SRC_PKG}/*.yaml ; do
        dst=$(basename wf)
        parse ${wf} ${DEPLOY_PKG}/${dst}
    done
else
    echo "Invalid file type: '${SRC_PKG}'"
    exit 11
fi
echo "Build succeeded."
