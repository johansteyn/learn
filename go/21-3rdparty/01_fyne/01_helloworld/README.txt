Based on: 
  https://www.youtube.com/watch?v=-v1vz_NcWng
  https://www.youtube.com/watch?v=kQ3-jfEQqho

Not mentioned in videos... needed to run:

  % go mod init helloworld
  % go mod tidy

Then, to package:

  % go get fyne.io/fyne/cmd/fyne

Hmmm... doesn't work with Go version 1.17
And couldn't get it to work with 1.16 either...

This seemed to finally work (with 1.16):

  % go get fyne.io/fyne/v2/cmd/fyne

It puts a "fyne" executable in: ~/go/bin

Then I was able to package it with:

  % ~/go/bin/fyne package -icon helloworld.png

Nice! It creates a helloworld.app package that I can click and run :)

