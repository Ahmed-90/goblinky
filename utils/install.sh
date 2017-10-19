#!/bin/sh


mv ./goblinky /usr/local/bin/goblinky
mv ./init.sh /etc/init.d/goblinky

apt-get update
apt-get install wiringpi

update-rc.d goblinky defaults