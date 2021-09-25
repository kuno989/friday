################################
# friday web & engine server 빌드
################################
FROM debian:stretch-slim
LABEL email="kuno@twelvelabs.io"
LABEL name="Kuno Kim"
LABEL version="1.0"
LABEL description="Friday antivirus analyzer"

RUN echo "Installing ..." \
    && apt-get update \
    && apt-get install -qq --no-install-recommends ca-certificates

##### Install dependencies #####
RUN echo "Installing Dependencies" \
    && buildDeps="automake unzip wget libtool make gcc pkg-config" \
    && apt-get install -y $buildDeps --no-install-recommends

####### Yara 설치 #######
ENV YARA_VERSION        4.1.1
ENV YARA_ARCHIVE        $YARA_VERSION.tar.gz
ENV YARA_DOWNLOAD_URL   https://github.com/VirusTotal/yara/archive/v$YARA_ARCHIVE
ENV YARA_REPO_REPO https://github.com/Yara-Rules/rules.git

RUN echo "Installing Yara..." \
    && apt-get install libssl-dev libglib2.0-0 build-essential autoconf libtool -y \
	&& wget $YARA_DOWNLOAD_URL \
    && tar zxvf v$YARA_ARCHIVE \
    && cd ./yara-$YARA_VERSION \
    && ./bootstrap.sh \
    && ./configure \
    && make \
    && make install

####### Capstone 설치 #######
RUN echo "Installing Capstone..." \
    && apt-get install libcapstone-dev -y

######## Installing TRiD ########
RUN echo "Installing TRiD..." \
    && wget http://mark0.net/download/trid_linux_64.zip -O /tmp/trid_linux_64.zip \
    && wget http://mark0.net/download/triddefs.zip -O /tmp/triddefs.zip \
    && cd /tmp \
    && unzip trid_linux_64.zip \
    && unzip triddefs.zip \
    && chmod +x trid \
    && mv trid /usr/bin/ \
    && mv triddefs.trd /usr/bin/

####### Installing Exiftool #######
RUN echo "Installing Exiftool..." \
    && wget https://exiftool.org/Image-ExifTool-12.26.tar.gz \
    && gzip -dc Image-ExifTool-12.26.tar.gz | tar -xf - \
    && cd Image-ExifTool-12.26 \
    && perl Makefile.PL \
    && make test \
    && make install

####### Installing File #######
RUN echo "Installing File..." \
    && apt-get install -qq file

####### Installing DiE #######
ENV DIE_VERSION     2.05
ENV DIE_URL         https://github.com/horsicq/DIE-engine/releases/download/$DIE_VERSION/die_lin64_portable_$DIE_VERSION.tar.gz
ENV DIE_ZIP         /tmp/die_lin64_portable_$DIE_VERSION.tar.gz
ENV DIE_DIR         /opt/die/

RUN echo "Installing DiE..." \
	&& apt-get install libglib2.0-0 -y \
	&& wget $DIE_URL -O $DIE_ZIP \
	&& tar zxvf $DIE_ZIP -C /tmp \
	&& mv /tmp/die_lin64_portable/ $DIE_DIR

RUN echo "Installing DiE..." \
    && apt-get install libglib2.0-0 -y \
    && wget $DIE_URL -O $DIE_ZIP \
    && tar zxvf $DIE_ZIP -C /tmp \
    && mv /tmp/die_lin64_portable/ $DIE_DIR


####### friday 설치 #######
ENV GO_VERSION 1.17.1
ENV GO111MODULE on
COPY . /go/src/github.com/kuno989/friday
RUN echo "Install Go..." \
     && ARCH="$(dpkg --print-architecture)" \
     && wget https://golang.org/dl/go$GO_VERSION.linux-$ARCH.tar.gz -O /tmp/go.tar.gz \
     && tar -C /usr/local -xzf /tmp/go.tar.gz \
     && export PATH=$PATH:/usr/local/go/bin \
     && echo "building friday..." \
     && export GOPATH=/go \
     && go version \
     && cd /go/src/github.com/kuno989/friday \
     && go build -ldflags "-X main.Version=$(cat VERSION) -X main.BuildTime=$(date -u +%Y%m%d)" -o /bin/friday \
     && echo "===> 삭제중..." \
     && apt-get purge -y --auto-remove $buildDeps && apt-get clean \
     && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* /go /usr/local/go

WORKDIR /malware
#삭제 예정
EXPOSE 8888
ENTRYPOINT ["/bin/friday","server"]