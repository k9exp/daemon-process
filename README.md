### A daemon process

Following [4 Steps for Daemonizie Go Application](https://ieftimov.com/posts/four-steps-daemonize-your-golang-programs)

```bash
go run main.go run.go

# this will print the PID (process id)
```

## Send Signals

```bash
kill -SIGTERM PID
kill -SIGINT PID
kill -SIGHUP PID
kill -SIGKILL PID
```