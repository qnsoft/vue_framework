@echo off

rem ��ǰbat������
echo ==================begin========================

cls 

SET Qn_cms_PATH=%~d0

SET Qn_cms_DIR=%~dp0

color 0a 

TITLE Qn_cms ������� Power By Ants (http://www.cnwtn.com)

CLS 

 
ECHO. 

ECHO. * * Qn_cms ������� Power By Ants (http://www.cnwtn.com)  *  

ECHO. * update by �Ը� 2013-03-13 *  

ECHO. 

:MENU 

ECHO. * Qn_cms ����list *  

tasklist|findstr /i "qn_cms.exe"

ECHO. 

    ECHO.  [1] ����Qn_cms 

    ECHO.  [2] �ر�Qn_cms  

    ECHO.  [3] ����Qn_cms 

    ECHO.  [4] �� �� 

ECHO. 

 
ECHO.������ѡ����Ŀ�����:

set /p ID=

    IF "%id%"=="1" GOTO start 

    IF "%id%"=="2" GOTO stop 

    IF "%id%"=="3" GOTO restart 

    IF "%id%"=="4" EXIT

PAUSE 

 

:start 

    call :startQn_cms

    GOTO MENU

 

:stop 

    call :shutdownQn_cms

    GOTO MENU

 

:restart 

    call :shutdownQn_cms

    call :startQn_cms

    GOTO MENU

 

:shutdownQn_cms

    ECHO. 

    ECHO.�ر�qn_cms...... 

    taskkill /F /IM qn_cms.exe > nul

    ECHO.OK,�ر�����qn_cms ����

    goto :eof

 

:startQn_cms

    ECHO. 

    ECHO.����qn_cms.exe...... 

    IF NOT EXIST "%Qn_cms_DIR%qn_cms.exe" ECHO "%Qn_cms_DIR%qn_cms.exe"������ 

 

    Qn_cms_PATH% 

 

    cd "%Qn_cms_DIR%" 

 

    IF EXIST "%Qn_cms_DIR%qn_cms.exe" (

        echo "start '' qn_cms.exe"

        start "" qn_cms.exe

    )

    ECHO.OK

    goto :eof