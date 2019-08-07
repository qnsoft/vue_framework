@echo off

rem 当前bat的作用
echo ==================begin========================

cls 

SET Qn_cms_PATH=%~d0

SET Qn_cms_DIR=%~dp0

color 0a 

TITLE Qn_cms 管理程序 Power By Ants (http://www.cnwtn.com)

CLS 

 
ECHO. 

ECHO. * * Qn_cms 管理程序 Power By Ants (http://www.cnwtn.com)  *  

ECHO. * update by 辉哥 2013-03-13 *  

ECHO. 

:MENU 

ECHO. * Qn_cms 进程list *  

tasklist|findstr /i "qn_cms.exe"

ECHO. 

    ECHO.  [1] 启动Qn_cms 

    ECHO.  [2] 关闭Qn_cms  

    ECHO.  [3] 重启Qn_cms 

    ECHO.  [4] 退 出 

ECHO. 

 
ECHO.请输入选择项目的序号:

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

    ECHO.关闭qn_cms...... 

    taskkill /F /IM qn_cms.exe > nul

    ECHO.OK,关闭所有qn_cms 进程

    goto :eof

 

:startQn_cms

    ECHO. 

    ECHO.启动qn_cms.exe...... 

    IF NOT EXIST "%Qn_cms_DIR%qn_cms.exe" ECHO "%Qn_cms_DIR%qn_cms.exe"不存在 

 

    Qn_cms_PATH% 

 

    cd "%Qn_cms_DIR%" 

 

    IF EXIST "%Qn_cms_DIR%qn_cms.exe" (

        echo "start '' qn_cms.exe"

        start "" qn_cms.exe

    )

    ECHO.OK

    goto :eof