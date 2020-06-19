## homepage

A homepage website built using nothing but standard library tools.


## Usage

You must provide self signed certs for localhost in the `certs` directory, 
under the names `homepage.crt` and `homepage.key`.

    docker-compose up --build python_app
    curl -G https://localhost:7001

    docker-compse up --build go_app
    curl -G https://localhost:7002

@NOTE: that `curl` doesn't like self signed certs by default, you'll need to tell it about your
certs, or use `brew install curl` version of curl.


## @TODO

### General
* Add the /go route
* Unify the data and templating
* 

### Go
* ~~Write a quick HTTP server~~
* ~~Figure out HTML templating~~
* ~~Figure out driving template data from JSON files~~
* ~~Add a data file reader based on route~~
* ~~Figure out how to return different content-type based on the accept header~~
* ~~Add templating for plain text~~
* Add markdown templating?
* ~~Refactor a bit to allow for more than just GET responses? - rendering temlpates is limiting~~
* ~~Just let the json responses be json responses - don't template those~~
* ~~Make it HTTPS~~
* Write a socket server
* Parse HTTP response
* ~~Build router~~
* ~~Build templating engine~~

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
