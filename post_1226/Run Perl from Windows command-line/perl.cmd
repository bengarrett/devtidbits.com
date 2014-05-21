::File: c:\terminal\cmd-scripts\perl.cmd
@echo off
@setlocal enableextensions enabledelayedexpansion
set perl=c:\terminal\perl\portableshell.bat
IF [%1]==[] (
	%perl% -v
) ELSE (
	%perl% %*
)
endlocal