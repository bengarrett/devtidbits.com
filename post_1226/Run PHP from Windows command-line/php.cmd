::File: c:\terminal\cmd-scripts\php.cmd
@echo off
@setlocal enableextensions enabledelayedexpansion
set php=c:\terminal\php\php.exe
IF [%1]==[] (
	%php% --version
) ELSE (
	%php% %*
)
endlocal