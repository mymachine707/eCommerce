rm -rf ./protogen
mkdir protogen

protoc --go_out=./protogen \
    --go-grpc_out=./protogen \
    ./protos/*.proto


rm -rf ./eCommerce_apiGETWAY/protogen
cp -r protogen ./eCommerce_apiGETWAY/

rm -rf ./category_service/protogen
cp -r protogen ./category_service/

rm -rf ./order_service/protogen
cp -r protogen ./order_service/

rm -rf ./authorization_client
cp -r protogen ./authorization_client/