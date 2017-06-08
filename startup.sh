#!/bin/bash
nohup ./EmailService  &> email.log &
echo $! > email.pid
