### A daemon process

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