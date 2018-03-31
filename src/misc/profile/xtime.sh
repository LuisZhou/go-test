#!/bin/sh
/usr/bin/time -f '%Uu %Ss %er %MkB %C' "$@"
# shows user time, system time, real time, and maximum memory usage