::File: c:\terminal\cmd-scripts\edit.cmd
@echo off
@setlocal enableextensions enabledelayedexpansion
set text_editor="C:\Program Files (x86)\Notepad++\notepad++.exe"
%text_editor% %*
endlocal