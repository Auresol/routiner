#! /bin/bash

rm $HOME/Documents/code/Routiner/server/mock.db
cp $HOME/Documents/code/Routiner/server/mock/mock_1.db mock.db
go run main.go mock
