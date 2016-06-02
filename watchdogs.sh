#!/bin/bash
now=`date '+%Y-%m-%d %H:%M:%S'`

grepFlag='molicms'
thisLog='/usr/local/molicms/watchlog'
 
baseDir="/usr/local/molicms"
sleepTime=60
 
if [ ! -f "$baseDir/run.sh" ]; then
    echo "$baseDir/run.sh missing, check again" > "$thisLog"
    exit
fi
 
#user=`whoami`
user="root"
if [ "$user" != "root" ]; then
    echo "this tool must run as *root*"
    exit
fi
 
while [ 0 -lt 1 ]
do
    now=`date '+%Y-%m-%d %H:%M:%S'`
    ret=`ps aux | grep "$grepFlag" | grep -v grep | wc -l`
    if [ $ret -eq 0 ]; then
        cd $baseDir
        echo "$now process not exists ,restart process now... " > "$thisLog"
        $baseDir/run.sh restart
        echo "$now restart done ..... "  > "$thisLog"
        cd $curDir
    else
        echo "$now process exists , sleep $sleepTime seconds " > "$thisLog"
    fi
    sleep $sleepTime
done