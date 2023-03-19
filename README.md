# Go 1.20 Performance Regression - Minimal Example for Reproduction

Change into one of the directories.

Run with

```
docker build -t test .
docker run test
```

This will output how long the go programm took to run.

I consistently see ~0.7s for the Go 1.19 version and >3s for Go 1.20.