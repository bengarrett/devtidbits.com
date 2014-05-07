::File: C:\terminal\cmd-scripts\shell-colour.cmd
@echo off
echo Welcome to %computername%
echo %username%'s user directory %userprofile%
echo.
prompt $E[0;32;40m%username%$E[1;30;40m@$E[0;32;40m%userdomain%$E[1;30;40m:$E[1;34;40m$P$E[1;30;40m$$$S$E[0m
C:\terminal\ansicon\x64\ansicon.exe