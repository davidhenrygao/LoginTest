@echo off

set pg=ProtoGen\protogen.exe

for /R proto %%f in (*.proto) do (
		echo generate csharp file: %%~nf.cs
		%pg% -i:%%f -o:csharp/%%~nf.cs -ns:proto
		)

pause
