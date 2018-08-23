FROM ubuntu:latest
LABEL maintainer "Muhammad Falak R Wani <falakreyaz@gmail.com>"
# get dependencies
RUN apt-get update
RUN apt-get install -y build-essential git-core m4 scons nano zlib1g zlib1g-dev libprotobuf-dev protobuf-compiler libprotoc-dev libgoogle-perftools-dev swig python-dev python
# checkout repo with mercurial
RUN useradd -u 1000 -ms /bin/bash gem5 && chown -R gem5:gem5 /home/gem5
WORKDIR /home/gem5
# build it
WORKDIR /home/gem5
USER gem5
RUN git clone --depth 1 https://github.com/gem5/gem5.git
RUN cd gem5 && scons -j$(nproc) --ignore-style build/X86/gem5.opt
USER root
RUN apt-get clean
RUN apt-get purge --auto-remove -y \
	&& rm -rf /var/lib/apt/lists/* \
	&& rm -rf /src/*.deb
USER gem5
ENTRYPOINT bash
