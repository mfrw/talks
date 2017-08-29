#!/bin/sh

sudo chroot busy-box /bin/sh
# busy-box => path to busy box install
# /bin/sh  => Command to execute in new root (will run relative to new root)
