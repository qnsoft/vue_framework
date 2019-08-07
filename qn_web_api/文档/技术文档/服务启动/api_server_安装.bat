@echo off
%CD%\Qn_cms_server.exe install
sc.exe start Qn_cms_server

echo -----------------------------
echo   Qn_cms_server服务安装成功！
echo -----------------------------
pause

