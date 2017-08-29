 docker run --rm -it \
          -v /run/user/1000/pulse/native:/home/chrome/pulse \ // HL
          -v /tmp/.X11-unix:/tmp/.X11-unix \
          -e DISPLAY=$DISPLAY --privileged \
          -v /dev/shm:/dev/shm --name chrome \
          mfrw/chrome
