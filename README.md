## homepage

A homepage website built using nothing but standard library tools.


## Usage

You must provide self signed certs for localhost in the `certs` directory, 
under the names `homepage.crt` and `homepage.key`.

    docker-compose up --build python_app
    curl -G https://localhost:7001

@NOTE: that `curl` doesn't like self signed certs by default, you'll need to tell it about your
certs, or use `brew install curl` version of curl.


## @TODO

### Go
* Write a socket server
* Parse HTTP response
* Build router
* Build templating engine

### Python
* Break down the templating
** Should do partials
** JSON response should just be the data
** Figure out a better way of printing lists

### Deployment
* Get a k3s cluster working
* Deploy apps to k3s
* Write a websocket (or something) link from VPS to cluster
* Find an ingress controller for cluster (nginx L4?)
