package mgmt

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/gen2brain/malgo"
)

var lastFrameProcessed int64 = 0
type CaptureDevice struct {
	Name string
	ID   string
}

type Audio struct {
	Context *malgo.AllocatedContext
	ActiveDevice *malgo.Device
}
func (a *Audio) SetDevice(id string) {
	cbs := malgo.DeviceCallbacks{
		Data: a.AudioReceived,
		Stop: a.DeviceStopped,
	}
	var idPtr unsafe.Pointer = nil
	devs, _ := a.Context.Context.Devices(malgo.Capture)
	for _, dev := range devs {
		if dev.String() == id {
			idPtr = dev.ID.Pointer()
		}
	}
	deviceConfig := malgo.DefaultDeviceConfig(malgo.Capture)
	// if id != "default" { deviceConfig.Capture.DeviceID = idPtr } else { deviceConfig.Capture.DeviceID = nil }
	deviceConfig.Capture.DeviceID = idPtr
	deviceConfig.Capture.Format = malgo.FormatS16
	deviceConfig.Capture.Channels = 1
	deviceConfig.Playback.Format = malgo.FormatS16
	deviceConfig.Playback.Channels = 1
	deviceConfig.SampleRate = 44100
	deviceConfig.Alsa.NoMMap = 1
	newdev, err := malgo.InitDevice(a.Context.Context, deviceConfig, cbs)
	if err != nil {
		fmt.Println("Error initializing device")
		fmt.Println(err)
	}
	a.ActiveDevice = newdev
	
	if a.ActiveDevice.Start() != nil {
		fmt.Println("Error starting device")
		fmt.Println(err)
	}
}
func (a *Audio) AudioReceived(pSample2, pSample []byte, framecount uint32) {
	timeStamp := time.Now().UnixMicro()
	if timeStamp - lastFrameProcessed > 16000 {
		fmt.Println("Data received", framecount, timeStamp)
		lastFrameProcessed = timeStamp
	}
}
func (a *Audio) DeviceStopped() {
	fmt.Println("Device stopped")
}