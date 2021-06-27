# Build docker
```
STEP 1: docker build . -t order
STEP 2: Docker run -d -p 3004:3004 --name order -e VARIABLE_ENVIRONMENT order
STEP 3: dock ps
```