# harbor replication

Currently the replication on [harbor] is failing in the lasts versions, due to changes in the API. This is a dirty-and fast script to replicate. Make public so you can avoid the pain

Currently only the replication of repositories, images and tags are replicated, if you need to copy charts and/or labels, feel free to improbe or fork it :)

## usage

- Edit cmd/main.go and change the constants to point to your needs:

    - **harborFrom** : endpoint for the harbor repo you want to copy from (example harbor.mydomain.net)
    - **harborUserFrom**: Harbor admin user (or an user with capacity to read and pull all images from the harbor repo you want to copy from)
    - **harborPasswordFrom** : Password for the above user
    - **harborTo** : endpoint for the harbor repo you want to copy to (example new-harbor.mydomain.net)
    - **harborUserTo**: Harbor admin user (or an user with capacity to write new repositories and push images into your destiny harbor repo
    - **harborPasswordTo** : Password for the above user

- run from command line:
```
$ go run cmd/main.go
```