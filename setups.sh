export SRC_PATH
SRC_PATH=$PWD
echo "SRC_PATH is $SRC_PATH"
echo "++++++++++++++++++++++++++++++++++++++++++"
./account/service_setup.sh
./auth/service_setup.sh