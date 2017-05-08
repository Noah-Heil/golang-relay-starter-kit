FROM node:6

RUN apt-get update && apt-get install -y --no-install-recommends \
		vim \
                wget \
                libxml2-dev \
                libxslt1-dev \
                libpq-dev \
                supervisor \
                ca-certificates \
                net-tools \
		postgresql-client \
		g++ \
		gcc \
		libc6-dev \
		make \
		curl \
		git \
		mercurial \
		binutils \
		bison \
		build-essential \
		pkg-config \
		&& rm -rf /var/lib/apt/lists/*

ENV GOLANG_VERSION 1.8.1
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 a579ab19d5237e263254f1eac5352efcf1d70b9dacadb6d6bb12b0911ede8994

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
	&& echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - \
	&& tar -C /usr/local -xzf golang.tar.gz \
	&& rm golang.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

RUN mkdir -p "$GOPATH/src/github.com/Noah-Heil/golang-relay-starter-kit/"

WORKDIR $GOPATH/src/github.com/Noah-Heil/golang-relay-starter-kit/

COPY go-wrapper /usr/local/bin/
 
CMD tail -f /dev/null
