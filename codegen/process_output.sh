#!/bin/sh
script_dir=$(dirname "$0")
for service in "action identity" ; do
    src_dir="$script_dir/generated/$service-go"
    dest_dir="$script_dir/../services/$service/generated"
    echo "Replacing services/$service/generated with newly-generated code..."
    if [[ ! -d   ]] ; then
        echo "$script_dir/generated/$service-go not found, skipping ..."
        continue
    fi
    if [[ -d $dest_dir ]] ; then
        rm -rf $dest_dir
    fi
    cp -R $src_dir $dest_dir
    # TODO: remove this in codegen if possible
    rm -rf $dest_dir/api
done