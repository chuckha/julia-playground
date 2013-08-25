# Julialang on Ubuntu
FROM ubuntu:12.04
ADD src/debs /root/tmp
RUN apt-get update -qq -y
RUN apt-get install build-essential -y
RUN apt-get install python-software-properties -y
RUN add-apt-repository ppa:staticfloat/julia-deps -y
RUN add-apt-repository ppa:staticfloat/julianightlies -y
RUN apt-get update -qq -y
RUN apt-get install g++ git make patchelf gfortran llvm-3.2-dev libsuitesparse-dev libncurses5-dev libopenblas-dev liblapack-dev libarpack2-dev libfftw3-dev libgmp-dev libpcre3-dev libreadline-dev libdouble-conversion-dev libopenlibm-dev librmath-dev libuv-julia-dev libmpfr-dev -y
RUN cd /root/tmp && dpkg -i libunwind7_0.99-0.3ubuntu1_amd64.deb  
RUN cd /root/tmp && dpkg -i libunwind7-dev_0.99-0.3ubuntu1_amd64.deb
RUN apt-get install julia -y
RUN rm -rf /root/tmp
RUN useradd runner -d /home/runner -p ojOJI#@OIJOIEJFOIJEwojo@Ij@OEIJF
ADD src/filewriter /home/runner/
USER runner
CMD ["./home/runner/filewriter"]
