# bTaskee

## Test Flow
- make proto
- make setup
- make base
- make build
- make deploy
- make test task=$Test
    * TestCreateUser (Create 1000 user prepare for quality test)
    * TestBasicFlow (CreateTask - AcceptTask - ConfirmTask)
    * TestMainFlow (CreateTask - 4 AcceptTask - ConfirmTask)
    * TestAcceptedLate (CreateTask - AcceptTask - ConfirmTask - AcceptTask)
    * TestQuality (Full MainFlow with 1000/5 loop = 200 task - 1 asker and 4 tasker with each task)
    
# 0. Prerequirement

- Install docker

    https://docs.docker.com/engine/install/

- Install kubectl

    https://kubernetes.io/docs/tasks/tools/

- Install istio

    https://github.com/istio/istio/releases

# 1. Usage

## Proto
- make proto

## MongoDB
- docker run -d --name mongodb-container -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=password mongo

## Redis
- docker run --name redis -p 6379:6379 -d redis


