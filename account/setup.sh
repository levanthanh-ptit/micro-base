export SRC_PATH=$PWD
cd "$SRC_PATH/account"
make proto
make build
cd "$SRC_PATH"