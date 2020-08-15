if test -z "$SRC_PATH"
then
  echo "SRC_PATH: is undefined"
  SRC_PATH=$PWD
  echo "new SRC_PATH: $SRC_PATH"
fi

## Define service name local scope
SERVICE_NAME=account

cd "$SRC_PATH/$SERVICE_NAME"
make proto
make build
cd "$SRC_PATH"
micro run $SERVICE_NAME
micro status
echo "++++++++++++++++++++++++++++++++++++++++++"