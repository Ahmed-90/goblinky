#!/bin/sh


mv goblinky /usr/local/bin/goblinky
mv init.sh /etc/init.d/goblinky

update-rc.d goblinky defaults