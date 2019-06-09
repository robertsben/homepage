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

* Create a helm chart for app with ingress service
* Deploy app to GKE via helm in circle
* Figure out k8s gcr secret
* Figure out a cert for GKE
* Figure out getting domain pointed at ingress
