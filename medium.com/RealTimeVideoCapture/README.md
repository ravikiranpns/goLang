Realtime Video Capture with Go
Using Go and the Video for Linux (V4L2) API to stream videos from camera hardware natively without using CGo.

This write up explores how to access video device drivers using the Video for Linux API (or V4L2) and the Go standard libraries, with no Cgo, to build realtime video capture and streaming programs.

How this got started
As with many endeavors in 2020, this got started because of COVID, specifically due to the shortage of webcams. Eventually, my search for a solution got turned into a hobby activity for building a webcam using a Raspberry Pi and the HQ Camera. That took me through a long journey of researching Linux drivers, learning about the V4L2 API, and understanding the USB Video Class gadgets (UVC gadgets).

Eventually I settled on understanding how the V4L2 API works and how to use it to do simple video capture on the Raspberry Pi. Naturally, I wondered if that could be done using Go. Fortunately, all of the system calls needed to access the device driver for video capture are well-implemented as part of the Go standard library 🎉 This means that Go programs can stream video from camera devices, directly, without the need to make calls to existing C libraries using Cgo.

The V4L2 API
It is important to have some background information about the Video for Linux API (if you are familiar with V4L2, skip to the next section).It is a large corpus of Kernel code that exposes an API for userspace applications. The API can be used for device control, video stream capture, video stream output, and other media functionalities as documented here.

Using the V4L2 API consist of making several several system calls to interact with the device driver. At first the application will open the device to get a file descriptor. Then the application will make a series of ioctl calls to send encoded command codes, along with struct values to collect the results, to configure the device so that its video data can be captured.

V4L2 in Go
Because this approach does not use any bindings to C libraries, all constant values, types, functionalities needed to use the V4L2 API, must be recreated in Go. For instance, the following snippet shows a Go struct, and its C equivalent, that is used to send a V4L2 command (see <linux/videodev2.h> header file).

// Go
type RequestBuffers struct {
  Count        uint32
  StreamType   uint32
  Memory       uint32
  Capabilities uint32
  Reserved     [1]uint32
}
// C
struct v4l2_requestbuffers {
  __u32 count;
  __u32 type;
  __u32 memory;
  __u32 capabilities;
  __u32 reserved[1];
};
Capturing a single video frame in Go
To illustrate how a Go program can capture image data from a video camera, this section provides a walkthrough of a simple program that captures a single frame and saves it to a file. Each step is discussed separately to show the mechanism of using V4L2 with Go.

The source snippets shown are abbreviated and omit error-handling. The full source is at https://github.com/vladimirvivien/go4vl/examples/raw_v4l2

Encoding the V4L2 commands
Each V4L2 command is represented by an encoded uint32 value stores a command ID, the type of the command, the argument information. Before we can send any command, we need to implement Go code to help with encoding those V4L2 request values as shown below.

// v4l2 command encoding
const (
  iocOpNone  = 0
  iocOpWrite = 1
  iocOpRead  = 2
  iocTypeBits   = 8
  iocNumberBits = 8
  iocSizeBits   = 14
  iocOpBits     = 2
  numberPos = 0
  typePos   = numberPos + iocNumberBits
  sizePos   = typePos + iocTypeBits
  opPos     = sizePos + iocSizeBits
)
// ioctl command encoding
func ioEnc(iocMode, iocType, number, size uintptr) uintptr {
  return (iocMode << opPos) |
    (iocType << typePos) |
    (number << numberPos) |
    (size << sizePos)
}
// convenience funcs
func ioEncR(iocType, number, size uintptr) uintptr {
  return ioEnc(iocOpRead, iocType, number, size)
}
func ioEncW(iocType, number, size uintptr) uintptr {
  return ioEnc(iocOpWrite, iocType, number, size)
}
func ioEncRW(iocType, number, size uintptr) uintptr {
  return ioEnc(iocOpRead|iocOpWrite, iocType, number, size)
}
// fourcc implements C v4l2_fourcc 
func fourcc(a, b, c, d uint32) uint32 {
  return (a | b<<8) | c<<16 | d<<24
}
You can find more detail about V4L2 command encoding and the section in the C code where this comes from here and the C implementation of the four-character pixel format encoding is here.

Another convenience function that needs to be implemented is the following which wraps the system call for IO control (or ioctl) system calls.

import(
  sys "golang.org/x/sys/unix"
)
// wrapper for ioctl system call
func ioctl(fd, req, arg uintptr) (err error) {
  errno := sys.Syscall(sys.SYS_IOCTL, fd, req, arg)
  if errno != 0 {
    err = errno
    return
  }
  return nil
}
Open the device
To send any system command to the device, it must be opened first by getting a file descriptor to the device. The descriptor value will be used for all subsequent system calls.

func main(){
  devFile, err := os.OpenFile(devName, sys.O_RDWR|sys.O_NONBLOCK, 0)
  if err != nil { ... }
  // file descriptor fd used in all subsequent calls
  fd := devFile.Fd()
  ...
}
Set the pixel format
Before the device can be used, the application must set the format and the size of the frames that will be captured. To do this, two types, Format and PixFormat, are implemented in Go to represent their C counterparts, as shown below.

// Format represents C type v4l2_format
type Format struct {
  StreamType uint32
  fmt        [200]byte // union
}
// PixFormat represents C type v4l2_pix_format
type PixFormat struct {
 Width        uint32
 Height       uint32
 PixelFormat  uint32
 Field        uint32
 BytesPerLine uint32
 SizeImage    uint32
 Colorspace   uint32
 Priv         uint32
 Flags        uint32
 YcbcrEnc     uint32
 Quantization uint32
 XferFunc     uint32
}
// values for pix format v4l2_fields
const (
  FieldAny  uint32 = iota 
  FieldNone               
)
Next, we can use the previous types to set the format using function setFormat .

func setFormat(fd uintptr, pixFmt PixFormat) error {
  format := Format{StreamType: BufTypeVideoCapture}
  
  // a bit of C union type magic with unsafe.Pointer
  *(*PixFormat)(unsafe.Pointer(&format.fmt[0])) = pixFmt
  // encode command VIDIC_S_FMT
  vidiocSetFormat := ioEncRW(
    'V', 5, uintptr(unsafe.Sizeof(Format{})),
  )
  // send command
  return ioctl(fd,vidiocSetFormat,uintptr(unsafe.Pointer(&format)))
}
Each V4L2 command has numerical value. In the previous snippet, note how the command encoding function ioEncRW is used to encode and store the command value, used to set the pixel format, in vidioSetFormat. Once encoded, the command is sent to the driver using the ioctl call.

That model (command encoding, ioctl command dispatch, collect result in a structure) is followed for most V4L2 ioctl calls covered in the rest of this write up.

Device buffer allocation
The next step is to request the driver to allocate a number of internal buffers before they can be mapped into the application’s address space for use. To do this, we must first implement the Go types that will be used as both argument and to store the result for the command, as shown below.

// RequestBuffers represents C type v4l2_requestbuffers 
type RequestBuffers struct {
  Count        uint32
  StreamType   uint32
  Memory       uint32
  Capabilities uint32
  Reserved     [1]uint32
}
// constant representing memory map IO stream mode
const StreamMemoryTypeMMAP uint32 = 1    // see V4L2_MEMORY_MMAP 
Next, let us define function that uses the RequestBuffers type to send the V4L2 command to allocate count device buffers.

func reqBuffers(fd uintptr, count uint32) error {
  reqbuf := RequestBuffers{
    StreamType: BufTypeVideoCapture,
    Count:      count,
    Memory:     StreamMemoryTypeMMAP,
  }
 
  vidiocReqBufs := ioEncRW(
    'V', 8, uintptr(unsafe.Sizeof(RequestBuffers{})),
  )
  return ioctl(fd, vidiocReqBufs, uintptr(unsafe.Pointer(&reqbuf)))
}
Map the device address space
V4L2 supports the ability to map the device’s address space, in the application, as a way to stream data. This allows the application to directly access the device buffers using fast, no-copy memory operations. As before, we, first, need to define Go types to send the V4L2 commands.

// BufferInfo implements C type v4l2_buffer
type BufferInfo struct {
  Index      uint32
  StreamType uint32
  BytesUsed  uint32
  ...
  Memory     uint32
  m          [unsafe.Sizeof(&BufferService{})]byte // C union type
  Length     uint32
  ...
}
// BufferService represents the embedded union m
// in v4l2_buffer C type.
type BufferService struct {
  Offset  uint32
  UserPtr uintptr
  Planes  uintptr
  FD      int32
}
The types, in the snippet above, are used to request device buffer information. That information can then be used to establish the memory map properly, which is what is implemented in the following function, mmapBuffer.

func mmapBuffer(fd uintptr, idx uint32) ([]byte, error) {
  buf := BufferInfo{
    StreamType: BufTypeVideoCapture,
    Memory:     StreamMemoryTypeMMAP,
    Index:      idx,
  }
  // send the V4L2 VIDIOC_QUERYBUF command via ioctl
  vidiocQueryBuf := ioEncRW(
    'V', 9, uintptr(unsafe.Sizeof(BufferInfo{})),
  ) 
  err := ioctl(fd, vidiocQueryBuf, uintptr(unsafe.Pointer(&buf)))
  if err != nil {return nil, err}
  // cast embedded union BufferInfo.m to BufferService
  bufSvc := *(*BufferService)(unsafe.Pointer(&buf.m[0]))
  // map the device memory to a []byte 
  mbuf, err := sys.Mmap(
    int(fd), 
    int64(bufSvc.Offset), 
    int(buf.Length), 
    sys.PROT_READ|sys.PROT_WRITE, sys.MAP_SHARED,
  )
  if err != nil {return nil, err}
  return mbuf, nil
}
Queue the device buffers
Before a buffer can be filled with device data, it must be queued in the devices incoming buffer queue. This is done in the following, queueBuffer, function with an ioctl call.

func queueBuffer(fd uintptr, idx uint32) error {
  buf := BufferInfo{
    StreamType: BufTypeVideoCapture,
    Memory:     StreamMemoryTypeMMAP,
    Index:      idx,
  }
  // send VIDIOC_QBUF command via ioctl
  vidiocQueueBuf := ioEncRW(
    'V', 15, uintptr(unsafe.Sizeof(BufferInfo{})),
  )
 
  return ioctl(fd, vidiocQueueBuf, uintptr(unsafe.Pointer(&buf)))   
}
Turn on the stream
Now, the application is ready to request that the device turn on its stream so that the buffers can be filled. The following function, startStreaming, sends ioctl to the driver to open its stream.

func startStreaming(fd uintptr) error {
  bufType := BufTypeVideoCapture
  // send VIDIOC_STREAM on command
  vidiocStreamOn := ioEncW(
    'V', 18, uintptr(unsafe.Sizeof(int32(0))),
  )
  return ioctl(fd, vidiocStreamOn,uintptr(unsafe.Pointer(&bufType)))
}
Wait for device read-ready signal
When the device starts to stream, it will fills its buffers as each frame becomes available. The application can use the select system call to monitor when the device’s file descriptor becomes ready for the next read operation as is implemented in the following function, waitForDeviceReady.

func waitForDeviceReady(fd uintptr) error {
  timeval := sys.NsecToTimeval((2 * time.Second).Nanoseconds())
  var fdsRead sys.FdSet
  fdsRead.Set(int(fd))
  for {
    n, err := sys.Select(int(fd+1), &fdsRead, nil, nil, &timeval)
    switch n {
    case -1:
      if err == sys.EINTR { continue }
      return err
    case 0:
      return fmt.Errorf("wait for device ready: timeout")
    default:
      return nil
    }
  }
}
Dequeue the device buffers
Once the application is notified that the device’s file descriptor is ready for a read (see above), it sends a request to the driver to dequeue the filled buffers so that they can be accessed using the mapped byte slice.

func dequeueBuffer(fd uintptr) (uint32, error) {
  buf := BufferInfo{
    StreamType: BufTypeVideoCapture,
    Memory:     StreamMemoryTypeMMAP,
  }
  // send VIDIOC_DQBUF command
  vidiocDequeueBuf := ioEncRW(
    'V', 17, uintptr(unsafe.Sizeof(BufferInfo{})),
  )
  err := ioctl(fd, vidiocDequeueBuf, uintptr(unsafe.Pointer(&buf)))   
  err != nil {return 0, err}
  return buf.BytesUsed, nil
}
In the snippet above, The ioctl call fills the BufferInfo variable buf with several pieces of information including the number of bytes used by the device to fill its buffers (see buf.BytesUsed). The application can use that information to accurately read the image data from the mapped buffer.

Turnoff the stream
Lastly, when the application is done, streaming image data, it should send an ioctl request to the driver to stop the stream so that the device can cleanly release its resources as is shown in the following snippet.

func stopStreaming(fd uintptr) error {
  bufType := BufTypeVideoCapture
  
  // send VIDIOC_STREAMOFF command
  vidiocStreamOff := ioEncW(
    'V', 19, uintptr(unsafe.Sizeof(int32(0))),
  ) 
  return ioctl(
    fd, vidiocStreamOff, uintptr(unsafe.Pointer(&bufType)))
}
Putting it all together
The functions, defined in the previous section, can be put together in a simple example to capture a single frame and save the frame buffer into a JPEG file, as shown in the following source snippet (error-handling is omitted, see full source here).

func main() {
  devName := "/dev/video0"
 
  // open the device and get a file descriptor
  devFile, _ := os.OpenFile(devName, sys.O_RDWR|sys.O_NONBLOCK, 0)
  defer devFile.Close()
  fd := devFile.Fd()
  // Set the pixel format
  setFormat(fd, PixFormat{
    Width:       640,
    Height:      480,
    PixelFormat: PixelFmtMJPEG,
    Field:       FieldNone,
  })
  // request and allocate 3 device buffers
  reqBuffers(fd, 3); err != nil {log.Fatal(err)}
  // memory map a device buffer, at index 2, to byte slice data
  data, _ := mmapBuffer(fd, 2)
  
  // queue selected device buffer prior to starting the stream
  queueBuffer(fd, 2)
  // start the device stream
  startStreaming(fd)
  // wait for the device to be ready for read
  waitForDeviceReady(fd)
  // now dequeue the device buffer and get bytesUsed
  bytesUsed, _ := dequeueBuffer(fd)
  // save mapped buffer bytes to file
  jpgFile, _ := os.Create("capture.jpg")
  defer jpgFile.Close()
  jpgFile.Write(data[:bytesUsed])
  // release streaming resources
  stopStreaming(fd)
}
For continuous streaming, some the steps above can be put in a queue/wait for read/dequeue arrangement and wrapped within a loop, that runs until a stop condition, to grab multiple video frames.

Go4vl: Go package for video in Linux
The walkthrough, in the previous section, is intended to be informational to show the multitude of steps that are required to use V4L2 from Go. A better way, however, would be to have idiomatic Go constructs that hides those details. That is why I started the Go4vl project to provide a Go API that enables programmers to work with video data by hiding the complexities of the V4L2 API.

See project https://github.com/vladimirvivien/go4vl

Example: capturing multiple frames
The following snippet (error-handling omitted for brevity) is an example that shows the Go4vl used to capture 10 consecutive frames. Notice that the package provides a channel so that the program can stream the image frames.

See full source here.

import "github.com/vladimirvivien/go4vl/v4l2" 
func main() {
  devName := "/dev/video0"
 
  // open device
  device, _ := v4l2.Open(devName)
  defer device.Close()
  // configure pix format
  device.SetPixFormat(v4l2.PixFormat{
    Width:       640,
    Height:      480,
    PixelFormat: v4l2.PixelFmtMJPEG,
    Field:       v4l2.FieldNone,
  })
  // start the stream with 10 device buffers
  device.StartStream(10)
  ctx, cancel := context.WithCancel(context.TODO())
  // setup capture at 15 fps, then get the channel to receive frames
  frameChan, _ := device.Capture(ctx, 15)
  // loop and save 10 frames to files
  totalFrames := 10
  count := 0
  for frame := range frameChan {
    fileName := fmt.Sprintf("capture_%d.jpg", count) 
    // create and write file
    file, _ := os.Create(fileName)
    file.Write(frame) 
    file.Close()
    count++
    if count >= totalFrames { break }
  }
  cancel() // signal capture to stop
  device.StopStream(); err != nil {
}
Example: realtime webcam streaming
The following abbreviated snippet (detail and error-handling omitted) shows how Go4vl can be used to create a web server that captures and streams video, from a webcam feed, unto a web page in real time. The server uses the multipart/x-mixed-replace protocol to continuously send a stream of JPEG images to the browser from the same image source.

See full source here.

import "github.com/vladimirvivien/go4vl/v4l2"
var frames <-chan []byte // channel video frames
func main() {
  port := ":9090"
  devName := "/dev/video0"
  device, _ := v4l2.Open(devName)
  defer device.Close()
 
  // set device format
  device.SetPixFormat(v4l2.PixFormat{
    Width:       640,
    Height:      480,
    PixelFormat: v4l2.PixelFmtMJPEG,
    Field:       v4l2.FieldNone,
  })
  // start stream with 15 buffers
  device.StartStream(15)
  // start capture at 30 fps
  ctx, cancel := context.WithCancel(context.TODO())
  f, _ := device.Capture(ctx, 30)
  frames = f // make frames available.
  // setup http service
  http.HandleFunc("/webcam", servePage) // serve page
  http.HandleFunc("/stream", serveVideoStream) // video feed
  http.ListenAndServe(port, nil)
}
// serves html page built from template, detail not shown
func servePage(w http.ResponseWriter, r *http.Request) {
  ...
  pd := map[string]string{"path":"/stream"}
  t, _ := template.ParseFiles("webcam.html")
  t.Execute(w, pd) 
  ...
}
// streams jpg frames using multipart/x-mixed-replace
// which uses a series of logical boundaries to send multiple frames
func serveVideoStream(w http.ResponseWriter, req *http.Request) {
  const boundaryName = "Yt08gcU534c0p4Jqj0p0"
  // send multi-part header
  w.Header().Set(
    "Content-Type", 
    fmt.Sprintf("multipart/x-mixed-replace; boundary=%s", 
    boundaryName))
  // write frame boundaries to response
  for frame := range frames{
    io.WriteString(w, fmt.Sprintf("--%s\n", boundaryName))
    io.WriteString(w, "Content-Type: image/jpeg\n")
    w.Write(frame)
    io.WriteString(w, "\n")
  }
}
Conclusion
Using Go to do video capture works and it works well. This writeup attempts to provide all the background information needed to understand how to create Go code that can stream video data from hardware on Linux using V4L2. While the V4L2 API is complex at first, once it is understood it can be used in a large number of applications to capture or output media. Fortunately, if you work in Go, you can take advantage of that API to build great programs.

References
Linux media infrastructure userspace API
Video for Linux API (V4L2)
V4L2 header files
An image capture example in C
Capturing a webcam stream using V4L2 (C walkthrough)
Another V4L2 webcam package in Go
Streaming images from Raspberry Pi (multipart/x-mixed-replace header example)
Using multipart/x-mixed-replace for HTTP streaming video (Python)
183


1


183


1




More from Learning the Go Programming Language
Follow
Short and insightful posts for newcomers learning the Go programming language

Vladimir Vivien
Vladimir Vivien

·Jul 20, 2019

Working with Compressed Tar Files in Go
This post shows how to use the archive and the compress packages to create code that can programmatically build or extract compressed files from tar-encoded archive files. …