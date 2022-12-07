protoc --go_out=./protogen \
    --go-grpc_out=./protogen \
    ./protos/*.proto

rm -rf ./api_getway/protogen
cp -r protogen ./api_getway/

rm -rf ./category_service/protogen
cp -r protogen ./category_service/

rm -rf ./order_service/protogen
cp -r protogen ./order_service/