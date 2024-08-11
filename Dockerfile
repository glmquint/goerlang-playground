FROM erlang:25.3.2.6-slim

RUN apt-get update && apt-get install -y tmux ca-certificates
COPY . /app
WORKDIR /app
RUN tar -C /usr/local -xzf go1.22.6.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"
RUN echo "export PATH=/usr/local/go/bin:\${PATH}" >> ~/.profile
RUN go get github.com/ergo-services/ergo
