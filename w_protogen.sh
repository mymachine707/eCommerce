rm -rf ./protogen
mkdir protogen

protoc --go_out=./protogen \
    --go-grpc_out=./protogen \
    ./protos/*.proto


rm -rf ./getway_service/protogen
cp -r protogen ./getway_service/

rm -rf ./category_service/protogen
cp -r protogen ./category_service/

rm -rf ./order_service/protogen
cp -r protogen ./order_service/

rm -rf ./authorization_service/protogen
cp -r protogen ./authorization_service/