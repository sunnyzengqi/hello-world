#!/usr/bin/env bash

function Usage() {
	echo "Usage:
		1.control.sh start
		2. control.sh stop
		3. control.sh forcestop  #this option will killall the uranus on the processes"
}

module=nova-toc-backend2
app=bin/$module
logfile=./logs/app.log
conf=cfg/cfg.toml

function Start() {
	#exec "./$app"

	nohup ./${app} -c ${conf} >>${logfile} 2>&1 &
}

function Stop() {
	killall $module
	if [ $? -eq 0 ]; then
		echo -e "stop all nova-toc-backend2 ok"
	else
		echo -e "stop failed"
	fi
}
if [ $# == 0 ]; then
	Usage;
fi

OPT=$1

if [ "$OPT" = "start" ]; then
	Start
elif [ "$OPT" = "stop" ]; then
	Stop
fi

