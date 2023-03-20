# Go 1.20 Performance Regression - Minimal Example for Reproduction

Change into one of the directories, `19` contains a Dockerfile with Go 1.19, `20` contains a Dockerfile with Go 1.20.

Run with

```
docker build -t test .
docker run test
```

This will output how long the go program took to run.

I consistently see ~0.7s for the Go 1.19 version and >3s for Go 1.20.

The problem can be mitigated by making making `appuser` the owner of the directory that holds the code.
```
RUN chown -R appuser:appuser app
```