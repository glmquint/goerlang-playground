## Erlang demo scenario ##

This example demonstrates how Ergo's node interop with the Erlang node.

Here is output of this demonstration
![Screenshot from 2022-07-20 16-06-27](https://user-images.githubusercontent.com/118860/180004548-5916ecdd-f78a-4cae-bca7-3956bd710b0e.png)

### Usage

```bash
docker build -f Dockerfile -t goerlang-playground .
docker run -v .:/app --rm -it goerlang-playground bash
```

for a more specific use case:
```bash
docker.exe run -v .:/app --rm -it goerlang-playground bash -c "tmux new -d 'go run .' && tmux split-window -h 'erl -name erl-demo@127.0.0.1 -setcookie 123' && tmux attach"
```
