mkdir D:\ProgramData\Ghost
mkdir D:\ProgramData\Ghost\bin
mkdir D:\ProgramData\Ghost\lib
copy gl.exe D:\ProgramData\Ghost
copy jre.zip D:\ProgramData\Ghost
xcopy /E bin D:\ProgramData\Ghost\bin
xcopy /E lib D:\ProgramData\Ghost\lib
copy nowVer.txt D:\ProgramData\Ghost
copy launcher.ini D:\ProgramData\Ghost
copy ghostjc.jar D:\ProgramData\Ghost
copy ghostjc.ini D:\ProgramData\Ghost
copy lgl.bat D:\ProgramData\Ghost
greg.reg
D:
cd D:\ProgramData\Ghost
gl.exe
pause