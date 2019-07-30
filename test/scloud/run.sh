set -x
ls $PWD

python -m unittest discover . -v 2>&1
