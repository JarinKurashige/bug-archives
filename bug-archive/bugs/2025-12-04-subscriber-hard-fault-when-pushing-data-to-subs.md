# Subscriber hard fault when pushing data to subs

## Summary
When a service would prepare to send data to queues, the service ended up sending to a dead queue

## Symptoms
assert failed: prvNotifyQueueSetContainer queue.c:3362 (pxQueueSetContainer->uxMessagesWaiting < pxQueueSetContainer->uxLength)                                                                                                                                

## Root Cause
Application task did not properly remove itself from the subscriber list, so the service thought that we still has an active subscriber.

Service was working properly

## Fix
Properly removed application from subscriber list to prevent service from writing to a dead queue

## How To Detect This Bug In the Future
Confirm that we are not overwritting messages to services and closing out our applications properly

## Tags
#freertos #queue
