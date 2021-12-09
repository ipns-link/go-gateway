For local development  

0. sudo chown -R ubuntu go-gateway/
1. Build the dev container with tag name ipns-gateway-local
    docker build --target dev . -t ipns-gateway-local
2. Run the image in interactive terminal mode "-it" , with volume mount "-v ${PWD}:/work", with entry point as shell terminal "sh"
    docker run -it -v ${PWD}:/work -p 8888:8888 ipns-gateway-local sh
3. mkdir app  
4. cd app
5. go mod init github.com/ipns-link/go-gateway/app
6. create a app.go

For build
1. Build the build container with tag name ipns-gateway-build
    sudo docker build --target build . -t ipns-gateway-build

For deployment
1. Build the build container with tag name ipns-gateway
    sudo docker build --target runtime . -t ipns-gateway
    sudo docker build . -t ipns-gateway

attach to running docker
docker exec -it "name of the container" sh
