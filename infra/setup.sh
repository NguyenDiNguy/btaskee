istio_local() {
    echo "......setup istio profile : kiali, jaeger, prometheus, grafana......"
    istioctl install --set profile=demo -y

    echo "......envoy injection......"
    kubectl label namespace default istio-injection=enabled
    
    echo "......apply route......"
    kubectl apply -f ./infra/istio/route.yml

    echo "......please, manual setup database after afew time......"

    docker run -d --name mongodb-container -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=password mongo

    docker run --name redis -p 6379:6379 -d redis
}

############################################

istio_local       

