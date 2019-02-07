#!/usr/bin/env python3
from ctypes import *
from omxplayer.player import OMXPlayer
from pathlib import Path
from time import sleep
import pyaudio, re
import sys, os
from contextlib import contextmanager
import numpy as np

def Monitor(p):
    CHUNK = 512
    FORMAT = pyaudio.paInt16
    CHANNELS = 1
    RATE = 48000
    RECORD_SECONDS = 5
    WAVE_OUTPUT_FILENAME = "cache.wav"
    if p==None:
        p = pyaudio.PyAudio()
    stream = p.open(format=FORMAT,
                    channels=CHANNELS,
                    rate=RATE,
                    input=True,
                    frames_per_buffer=CHUNK)
    print("开始缓存录音")
    frames = []
    while (True):
        for i in range(0, 100):
            data = stream.read(CHUNK)
            frames.append(data)
        audio_data = np.fromstring(data, dtype=np.short)
        large_sample_count = np.sum( audio_data > 800 )
        temp = np.max(audio_data)
        if temp > 800 :
            print("\r检测到信号", temp, end='')
        else:
            print("\rNot检测到信号", end='')
    stream.stop_stream()
    stream.close()
    p.terminate()


ERROR_HANDLER_FUNC = CFUNCTYPE(None, c_char_p, c_int, c_char_p, c_int, c_char_p)

def py_error_handler(filename, line, function, err, fmt):
    pass

c_error_handler = ERROR_HANDLER_FUNC(py_error_handler)

@contextmanager
def noalsaerr():
    asound = cdll.LoadLibrary('libasound.so')
    asound.snd_lib_error_set_handler(c_error_handler)
    yield
    asound.snd_lib_error_set_handler(None)

class warnMessage():
    warn=[]
    error=[]
    def __init__(self):
        pass

    @property
    def total(self):
        return len(self.warn)+len(self.error)

    @property
    def warned(self):
        return (len(self.warn)+len(self.error))!=0

    def add_error(self, message):
            self.error += [message]

    def add_warn(self, message):
        self.warn+=[message]

    def add(self, level, message=''):
        level=level.lower()
        if level=='warn':
            self.add_warn(message)
        elif level=='error':
            self.add_error(message)

    def show_raw(self):
        print(self.total, self.warn, self.error)


class hideStdoutErr:
    def __init__(self, out='', err=''):
        self.custom_out=out
        self.custom_err=err
    def __enter__(self):
        self._original_stdout = sys.stdout
        self._original_stderr = sys.stderr
        if self.custom_out=='':
            sys.stdout = open(os.devnull, 'w')
            sys.stderr = open(os.devnull, 'w')
        else:
            sys.stdout = self.custom_out
            sys.stderr = self.custom_err

    def __exit__(self, exc_type, exc_val, exc_tb):
        if self.custom_out == '':
            sys.stdout.close()
        if self.custom_err == '':
            sys.stderr.close()
        sys.stdout = self._original_stdout
        sys.stderr = self._original_stderr
