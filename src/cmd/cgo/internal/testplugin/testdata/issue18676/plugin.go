package main

import "C"

import "testplugin/issue18676/dynamodbstreamsevt"

func F(evt *dynamodbstreamsevt.Event) {}
