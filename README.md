# TuPNG
## Introduction

TuPNG is a quick, smart Image compression tool by Go.It supports batch compressing images including jpg, jpeg, png, and gif. 

## Usage

You can compile *tupng.go* to generate an executable file, or you can run tupng.exe directly. 

TuPNG supports parameters as follows:

```go
-rt // the files' rootpath, default rootpath is "./image"
-op // the files' output path, default output path is "./output"
-wd // set output images' width in pxs, default width is the original image width
-qu // set output images' quality within [0,100], default quality is 70
-ni // whether output images keep original file names, 0 means no, and 1 means yes
```

