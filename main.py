#!/usr/bin/env python3
from audio import *
import wave

CHUNK=512
RECORD_SECONDS=2
WAVE_OUTPUT_FILENAME = r"\home\pi\Music\Oldboy.wav"
class device_info():
    format = pyaudio.paInt16
    in_channels = 1
    out_channels = 1
    rate = 44100
    index = 2
    def __init__(self, device_info):
        if device_info!=None:
            self.update(device_info)

    def update(self, di):
        if di.get('maxInputChannels'):
           self.in_channels=di.get('maxInputChannels') if di.get('maxInputChannels') else self.in_channels
           self.out_channels=di.get('maxOutputChannels') if di.get('maxOutputChannels') else self.out_channels
           self.rate=int(di.get('defaultSampleRate')) if di.get('defaultSampleRate') else self.rate
           self.index=di.get('index') if di.get('index') else self.index


with noalsaerr():
    p = pyaudio.PyAudio()
audio_usb_index=-1
warn=warnMessage()

for ii in range(p.get_device_count()):
    audio_part_name=p.get_device_info_by_index(ii).get('name')
    if 'USB' in audio_part_name or 'usb' in audio_part_name:
        if audio_usb_index!=-1:
            warn.add_warn('USB Index confilct')
        else:
            audio_usb_index=ii
            serStr=re.findall(r'hw:[\d,]+', audio_part_name)[0]
            audio_info=device_info(p.get_device_info_by_index(ii))

if warn.warned:
    print('Something happened!')
stream = p.open(input_device_index=audio_info.index,
                format=audio_info.format,
                channels=audio_info.in_channels,
                rate=audio_info.rate,
                input=True,
                frames_per_buffer=CHUNK)

print("开始录音,请说话......")

frames = []

for i in range(0, int(audio_info.rate / CHUNK * RECORD_SECONDS)):
    data = stream.read(CHUNK)
    frames.append(data)

print("录音结束,请闭嘴!")

stream.stop_stream()
stream.close()
p.terminate()

wf = wave.open(WAVE_OUTPUT_FILENAME, 'wb')
wf.setnchannels(audio_info.in_channels)
wf.setsampwidth(p.get_sample_size(audio_info.format))
wf.setframerate(audio_info.rate)
wf.writeframes(b''.join(frames))
wf.close()


'''
audio_usb_index=-1
for ii in range(p.get_device_count()):
    audio_part_name=p.get_device_info_by_index(ii).get('name')
    if 'USB' in audio_part_name or 'usb' in audio_part_name:
        if audio_usb_index!=-1:
        else:
            audio_usb_index=ii
            serStr=re.findall(r'hw:[\d,]+', audio_part_name)[0]

VIDEO_PATH = Path("/home/pi/py-work/gatekeeper/d.mp3")

player = OMXPlayer(VIDEO_PATH, args='-o alsa:'+serStr)

print('?')
sleep(15)

player.quit()
'''