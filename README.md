# Go 1.20 Performance Regression - Minimal Example for Reproduction

Change into one of the directories.

Run with

```
docker build -t test .
docker run test
```

This will output how long the go program took to run.

I consistently see ~0.7s for the Go 1.19 version and >3s for Go 1.20.

The problem disappears when not running with "appuser", i.e. commenting out line 3 and 10 in the Dockerfile of the Go 1.20 version.