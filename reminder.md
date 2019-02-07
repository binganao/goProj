不使用红外的原因是如果摄像头附近有热光源，会导致识别误差偏大，红外识别的兼容性低，对安装所处环境的要求较高。


Alsa: http://www.alsa-project.org/main/index.php/Main_Page
omxplayer -o omxplayer ~/a.wav  -o alsa:hw:1,0 # For usb sound card output