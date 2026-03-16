@echo off
echo Setting up folder protection...

REM Create the Service directory structure
if not exist "C:\Service" mkdir "C:\Service"
if not exist "C:\Service\application" mkdir "C:\Service\application"
if not exist "C:\Service\backups" mkdir "C:\Service\backups"
if not exist "C:\Service\logs" mkdir "C:\Service\logs"
if not exist "C:\Service\scripts" mkdir "C:\Service\scripts"
if not exist "C:\Service\service" mkdir "C:\Service\service"

REM Set folder permissions to prevent deletion
REM Remove delete permissions for Users group
icacls "C:\Service\application" /deny Users:D /t
icacls "C:\Service\application" /deny Users:DC /t

REM Grant read, write, and execute permissions but not delete
icacls "C:\Service\application" /grant Users:F /t
icacls "C:\Service\application" /deny Users:D /t

REM Hide sensitive files
attrib +h "C:\Service\application\.env"
attrib +h "C:\Service\application\storage\logs"

REM Set read-only on critical files
attrib +r "C:\Service\application\config\app.php"
attrib +r "C:\Service\application\config\database.php"

echo Folder protection applied successfully!
echo Application folder: C:\Service\application
echo - Users cannot delete the folder
echo - Sensitive files are hidden
echo - Critical config files are read-only
pause