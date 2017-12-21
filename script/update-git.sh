#!/bin/bash

ssh root@120.78.175.235 > /dev/null 2>&1 << eeooff
cd /root/github/insurance-file
git pull
cd /root/github/insurance-backend
git pull
exit