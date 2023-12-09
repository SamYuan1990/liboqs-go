FROM ubuntu:latest

# Install dependencies
RUN apt-get -y update && \
    apt-get install -y build-essential git cmake libssl-dev wget

RUN uname -m
RUN echo https://go.dev/dl/go1.21.5.linux-$(uname -m).tar.gz
RUN wget https://go.dev/dl/go1.21.5.linux-$(uname -m).tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.5.linux-$(uname -m).tar.gz

#https://go.dev/dl/go1.21.5.linux-s390x.tar.gz
# Get liboqs
RUN git clone --depth 1 --branch main https://github.com/open-quantum-safe/liboqs

# Install liboqs
RUN cmake -S liboqs -B liboqs/build -DBUILD_SHARED_LIBS=ON && \
    cmake --build liboqs/build --parallel 4 && \
    cmake --build liboqs/build --target install

# Enable a normal user
RUN useradd -m -c "Open Quantum Safe" oqs
USER oqs
WORKDIR /home/oqs

# Get liboqs-go
RUN git clone --depth 1 --branch main https://github.com/open-quantum-safe/liboqs-go.git

# Configure liboqs-go
ENV PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/home/oqs/liboqs-go/.config
ENV LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/lib
ENV PATH=$PATH:/usr/local/go/bin