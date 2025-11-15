# subdetect
A video processing tool to detect burned in subtitles

Currently there are a lot of false positives.

# TODOs:
* time filtering to remove single frame noise
* masking out the top part of the screen to speed things up
* language detection
* revisit false positives issue

# Libraries
* https://github.com/otiai10/gosseract for text recognition
 ( requires Tesseract to be installed )
* https://gocv.io/ for general computer vision tasks

## Notes
Running on a mac with homebrew, the LIBRARY_PATH and CPATH need to be set:
```bash
export LIBRARY_PATH="/opt/homebrew/lib"
export CPATH="/opt/homebrew/include"
```
