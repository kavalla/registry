Docker Registry for DomeOS
===

This is an implementation of docker registry customized for DomeOS, providing backend storage on [Sohu Cloud Storage](https://cs.sohu.com/ "Sohu Cloud Storage"). Note that this project is based on github docker/distribution project, so all the configurations provided officially can be applied to our registry.

## Running in a container

It's recommended that the registry run in a docker container:

```bash
sudo docker run --restart=always -d -p <your port>:5000 \
-e REGISTRY_STORAGE_SOHUSTORAGE_ACCESSKEY=<your access key> \
-e REGISTRY_STORAGE_SOHUSTORAGE_SECRETKEY=<your secret key> \
-e REGISTRY_STORAGE_SOHUSTORAGE_REGION=<your region> \
-e REGISTRY_STORAGE_SOHUSTORAGE_BUCKET=<your bucket> \
--name private-registry \
domeos/docker-registry-driver-sohustorage:1.0
```

Regions should be one of "bjcnc", "bjctc", "shctc" and "gzctc". Bucket must be existed and accessable.
