#!/bin/sh

set -x

cp agent /usr/local/bin/

pkill mcd_agent

cd /usr/local/bin/

cp agent mcd_agent

nohup mcd_agent serve -d &

