#!/bin/bash

function gophersay {
# base64 encoded then gzipped golang gopher ascii art
# from here: https://gist.github.com/belbomemo/b5e7dad10fa567a5fe8a
gopherart="H4sIAJ73h1cAA32OSxLCIBBE95yiy40VCoNHYOsZggyruMnGPebs9pCPSSztED49bxpw+SEDIOo0
z4j8nbAyqhRpBYuIiHPibBJq8nQnOduHjSfriHhaHv5cq6kFjTk5hSfXHqYDQo+C0DRaKVj3A3AF
25kTszbGsV6loYj1Lr8809TOVTIIx8a4Tczr43Q1QO57iMzWYWjZBevhyGDk983s2o4q03v+QpUx
b+lcNw2vAQAA
"

# Go Proverbs http://go-proverbs.github.io/
say[0]="Don't communicate by sharing memory, share memory by communicating."
say[1]="Concurrency is not parallelism"
say[2]="Channels orchestrate; mutexes serialize."
say[3]="The bigger the interface, the weaker the abstraction."
say[4]="Make the zero value useful."
say[5]="interface{} says nothing."
say[6]="Gofmt's style is no one's favorite, yet gofmt is everyone's favorite."
say[7]="A little copying is better than a little dependency."
say[8]="Syscall must always be guarded with build tags."
say[9]="Cgo must always be guarded with build tags."
say[10]="Cgo is not Go."
say[11]="With the unsafe package there are no guarantees."
say[12]="Clear is better than clever."
say[13]="Reflection is never clear."
say[14]="Errors are values."
say[15]="Don't just check errors, handle them gracefully."
say[16]="Design the architecture, name the components, document the details."
say[17]="Documentation is for users."
say[18]="Don't panic."

# random
rand=$[$RANDOM % 19]

# Echo random saying and art
echo ' ------------------------'
echo " ${say["$rand"]}"
echo "$gopherart" | base64 --decode | gzip -d

}

gophersay
