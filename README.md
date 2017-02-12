CourseCorrect
=============

An intelligent class scheduling assistant.

Prerequisites
----------------------

This project requires mingw-w64 on 64 bit Windows systems

Installing and Running
----------------------

    mkdir $GOPATH/src/27thisotope.ddns.net/mibzman/
    cd $GOPATH/src/27thisotope.ddns.net/mibzman/
    git clone http://27thisotope.ddns.net/mibzman/CourseCorrect-Student.git

	#install glide dependency management
    curl https://glide.sh/get | sh

	#install the dependencies
    glide install

    cd CourseCorrect-Student/frontend
    npm install
    npm start
    cd ..
    go build && ./CourseCorrect-Student

