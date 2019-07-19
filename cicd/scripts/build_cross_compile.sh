#!/bin/bash -e

if [[ -z "${SCLOUD_SRC_PATH}" ]] ; then
    echo "SCLOUD_SRC_PATH must be set, exiting ..."
    exit 1
fi

if [[ -n "${SIGN_PACKAGES}" ]]
then
    echo "\$SIGN_PACKAGES is set to \"${SIGN_PACKAGES}\", .asc signature files WILL be generated .."
else
    echo "\$SIGN_PACKAGES is set to \"${SIGN_PACKAGES}\", .asc signature files WILL NOT be generated .."
fi

BUILD_TARGETS_ARCH=( 386 amd64 )
BUILD_TARGETS_OS=( darwin linux windows )
TARGET_ROOT_DIR=bin/cross-compiled
ARCHIVE_DIR=${TARGET_ROOT_DIR}/archive
TAG=$(git describe --abbrev=0 --tags)

rm -fr ${TARGET_ROOT_DIR}
mkdir -p ${ARCHIVE_DIR}

function get_suffix {
    if [[ 'windows' == $1 ]]
    then
        return '.exe'
    else
        return ''
    fi
}

for os in ${BUILD_TARGETS_OS[@]}
do
    for arch in ${BUILD_TARGETS_ARCH[@]}
    do
        if [[ 'windows' == ${os} ]]
        then
            program_name='scloud.exe'
        else
            program_name='scloud'
        fi
        target_dir=${TARGET_ROOT_DIR}/${os}_${arch}
        mkdir -p ${target_dir}
        target_file=${target_dir}/${program_name}
        archive_file=${PWD}/${ARCHIVE_DIR}/scloud_${TAG}_${os}_${arch}
        echo "Building ${target_file}";
        # The -s flag strips debug symbols from Linux, -w from DWARF (darwin). This reduces binary size by about half.
        env GOOS=${os} GOARCH=${arch} go build -ldflags "-s -w" -a -o ${target_file} ${SCLOUD_SRC_PATH}/cli
        if [[ 'windows' == ${os} ]]
        then
            pushd ${target_dir}
            zip ${archive_file}.zip ${program_name}
            if [[ -n "${SIGN_PACKAGES}" ]]
            then
                echo "Generating signature file for ${archive_file}.zip .."
                gpg2 --armor --detach-sign ${archive_file}.zip
            fi
            popd
        else
            tar -C ${target_dir} -czvf ${archive_file}.tar.gz ${program_name}
            if [[ -n "${SIGN_PACKAGES}" ]]
            then
                echo "Generating signature file for ${archive_file}.zip .."
                gpg2 --armor --detach-sign ${archive_file}.tar.gz
            fi
        fi
    done
done

echo "Package archives created: "
echo ""
archives=$(ls ${ARCHIVE_DIR})
echo "${archives}"