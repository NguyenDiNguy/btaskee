#!/bin/bash
services=(identity booking)

deploy_local() {
    if [ "$1" == "" ]; then
        for value in "${services[@]}"
        do
            echo "......apply $value......"
            kubectl apply -f ./services/$value/k8s.yml 
        done           
    else
        echo "......apply $1......"
        kubectl apply -f ./services/$1/k8s.yml    
    fi
}

########################################
deploy_local "$1"
