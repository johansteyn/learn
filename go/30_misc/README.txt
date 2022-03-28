Looking for a way to handle user pressing CTRL-C...

What is the difference between SIGTERM and SIGINT?

https://en.wikipedia.org/wiki/Signal_(IPC)
https://www.gnu.org/software/libc/manual/html_node/Termination-Signals.html

SIGTERM comes from the "kill" shell command.
SIGINT (interrupt) is sent when the user types the INTR character, usually CTRL-C
SIGTSTP (terminal stop) is sent when CTRL-Z is pressed
SIGQUIT is sent when CTRL-\ is pressed

https://gobyexample.com/signals
https://golangcode.com/handle-ctrl-c-exit-in-terminal


