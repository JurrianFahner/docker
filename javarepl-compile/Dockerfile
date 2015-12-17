# JavaRepl
#
# VERSION 0.2

FROM ubuntu:wily

RUN apt-get -y update 
RUN apt-get -y upgrade

RUN apt-get -y install mercurial openjdk-8-jre-headless openjdk-8-jdk ant zip make build-essential libX11-dev libxext-dev libxrender-dev libxtst-dev libxt-dev libcups2-dev libfreetype6-dev libasound2-dev

# cd /opt/
WORKDIR /opt

# hg clone http://hg.openjdk.java.net/kulla/dev jshell
RUN hg clone http://hg.openjdk.java.net/kulla/dev jshell

# cd jshell/
WORKDIR /opt/jshell

# bash ./get_source.sh 
RUN sh ./get_source.sh

# bash ./configure 
RUN bash ./configure 


# make all
RUN make all

CMD /opt/jshell/build/linux-x86_64-normal-server-release/images/jdk/bin/jshell 
