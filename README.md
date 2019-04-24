# Google App Engine Go 1.11 Runtime Doesn't Cache Dependencies

The documentation of Google App Engine [mentions](https://cloud.google.com/appengine/docs/standard/go111/specifying-dependencies) that the Go 1.11 runtime, in the standard environment, caches fetched dependencies to reduce build times.

But it looks like the cache mentioned in the documentation is not used when using Go modules (with a `go.mod` file).

This repository is used to reproduce the issue following this methodology:

1. Host the source code of a dummy dependency at `github.com/ngrilly/test-golang-cache/hello`. The import path of the dependency is `test-golang-cache.appspot.com/hello`.
2. Host a static web page containing a `go-import` meta tag referencing the GitHub repository at `test-golang-cache.appspot.com/hello`.
3. Deploy to App Engine a dummy web server that depends on `test-golang-cache.appspot.com/hello`.
4. Check the log of requests served by `test-golang-cache.appspot.com` to verify if the dependency was fetched or not.
5. Repeat steps 3 and 4 several times to check that the cache works or not.

## Steps to reproduce

Checkout the projet outside of your $GOPATH (working outside your $GOPATH is important to enable Go modules):

    $ git clone https://github.com/ngrilly/test-golang-cache.git

Set your default project:

    $ gcloud config set project test-golang-cache

Deploy the static website containing the go-import meta tag:
    
    $ gcloud app deploy static

Make sure the server runs locally:

    $ cd server
    $ go run .

Deploy to Google App Engine:

    $ gcloud app deploy

In the Google Cloud console, go to [Cloud Build > History](https://console.cloud.google.com/cloud-build/builds), open the last build, and check the log.

Now go to [Logging > Logs](https://console.cloud.google.com/logs/viewer) and check if there is a request for `/hello?go-get=1`.

## Conclusion

Each time we deploy with `gcloud app deploy`, App Engine downloads the dependency test-golang-cache.appspot.com/hello.
It looks like the dependency cache mentioned in the [documentation](XXX) is not used.