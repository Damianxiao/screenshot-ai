@echo off
echo Starting Go backend...
start "" /B go run main.go


echo Starting React frontend...
cd .\frontend
start "" /B npm start

echo Both services are running...
pause
