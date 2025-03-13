############################################    
push_local() {
    echo "......build image...... $1" 
    docker build . -t btaskee/"$1" -f ./services/"$1"/Dockerfile

    echo "......push image...... $1"
    docker save btaskee/"$1"
}

build_local() {
    docker context use default
    if [ "$1" == "" ]; then
        for value in "${services[@]}"
        do
            push_local $value
        done           
    else
        push_local $1
    fi
}

services=(identity booking)

build_local "$1"       

echo "......clean up build stage image......"
docker image prune -f	
