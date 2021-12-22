#!/bin/bash

lastDay=`ls -1trd day* | sed -e 's/day//' | sort -n | tail -1`

let nextDay=lastDay+1

ld=day${lastDay}
nd=day${nextDay}

mkdir $nd
cp $ld/$ld.go $nd/$nd.go

sed -e "s/$ld/$nd/" $ld/go.mod > $nd/go.mod

touch $nd/test$nextDay.txt
touch $nd/input$nextDay.txt
