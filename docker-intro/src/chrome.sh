 docker run --rm -it \
          -v /tmp/.X11-unix:/tmp/.X11-unix \ // HL
          -e DISPLAY=$DISPLAY --privileged \
          -v /dev/shm:/dev/shm --name chrome \
          mfrw/chrome
