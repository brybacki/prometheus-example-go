
# # prometheus-example-go - Example go app with prometheus metrics

App starts http server with a few urls:
- `http://localhost:8123` - root url, shows Hell, means app is working. 

- `http://localhost:8123/update/${value}` - a way to set example_gauge value. Last part of path the value.  

- `http://localhost:8123/metrics` - prometheus metrics, shows one gauge, and requests counter
with labels per url/path.

# Building and running


```bash
cd src

GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o=./bin/main .

./bin/main
```