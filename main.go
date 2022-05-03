package main

/*
#cgo CFLAGS: -I ./include

//points to the right platform version of tflite libs
#cgo arm LDFLAGS: -L arm
#cgo darwin LDFLAGS: -L macosx
#cgo x86_64 LDFLAGS: -L x86_64

#cgo LDFLAGS: -ltensorflowlite_c

//Raspberry Pi needs to include libatomic when linking w/ tflite
#cgo arm LDFLAGS: -latomic

#include "tensorflow/lite/c/c_api.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"sync"
	"unsafe"
)

type TFGan struct {
	modelName   *C.char
	model       *C.TfLiteModel
	options     *C.TfLiteInterpreterOptions
	runner      *C.TfLiteInterpreter
	input       *C.TfLiteTensor
	inputBuffer []float32
	output      *C.TfLiteTensor
	mutex       sync.Mutex
}

func makeTFGan(modelName string) *TFGan {
	version := C.TfLiteVersion()
	fmt.Printf("Tensorflow Version: %v\n", C.GoString(version))
	name := C.CString(modelName)
	model := C.TfLiteModelCreateFromFile(name)
	if model == nil {
		fmt.Printf("failed to create model from - %v\n", C.GoString(name))
		return nil
	}
	options := C.TfLiteInterpreterOptionsCreate()
	if options == nil {
		fmt.Printf("failed to create options for %v\n", modelName)
		return nil
	}
	C.TfLiteInterpreterOptionsSetNumThreads(options, C.int32_t(4))

	runner := C.TfLiteInterpreterCreate(model, options)
	if runner == nil {
		fmt.Printf("failed to create interperter for %v\n", modelName)
		return nil
	}
	C.TfLiteInterpreterAllocateTensors(runner)
	input := C.TfLiteInterpreterGetInputTensor(runner, 0)
	if input == nil {
		fmt.Printf("input tensor is empty\n")
		return nil
	}
	output := C.TfLiteInterpreterGetOutputTensor(runner, 0)
	if output == nil {
		fmt.Printf("putput tensor is empty\n")
		return nil
	}

	return &TFGan{
		modelName:   name,
		model:       model,
		options:     options,
		runner:      runner,
		input:       input,
		output:      output,
		inputBuffer: []float32{},
	}
}

func (gan *TFGan) free() {
	C.TfLiteInterpreterDelete(gan.runner)
	C.TfLiteModelDelete(gan.model)
	C.free(unsafe.Pointer(gan.modelName))
}

func main() {

	gan := makeTFGan("model/tfliteModel.tflite")
	if gan == nil {
		fmt.Printf("failed \n")
		return
	}

	// input your data
	// x_sample = scaler.transform([[-12, -1]])
	// x_sample = scaler.transform([[14, -1]])
	// x_sample = scaler.transform([[10, -5]])
	// x_sample = scaler.transform([[6, -2.5]])
	
	// # result nearest to 1 is our array index
	// 0 => 0
	// 1.03 => 1
	// 0.23 => 2
	// 0.43 => 3
	// 0.03 => 4

	// inputData := []float32{ -0.70504665, -0.47355217 }
	// inputData := []float32{ 0.74449456, -0.47355217 }
	// inputData := []float32{ 0.52148825, -0.8003357 }
	inputData := []float32{ 0.2984819, -0.596096 }

	ptr1 := C.TfLiteTensorData(gan.input)
	if ptr1 == nil {
		fmt.Errorf("bad tensor")
	}

	n1 := uint(C.TfLiteTensorByteSize(gan.input)) / 4
	to := (*((*[1<<29 - 1]float32)(ptr1)))[:n1]
	copy(to, inputData)

	if C.TfLiteInterpreterInvoke(gan.runner) != C.kTfLiteOk {
		fmt.Printf("failed to run\n")
	}

	ptr := C.TfLiteTensorData(gan.output)
	if ptr == nil {
		fmt.Errorf("bad tensor")
	}

	n := uint(C.TfLiteTensorByteSize(gan.output)) / 4
	result := (*((*[1<<29 - 1]float32)(ptr)))[:n]

	fmt.Println(result)

	// # result nearest to 1 is our array index
	// 0 => 0
	// 1.03 => 1
	// 0.23 => 2
	// 0.43 => 3
	// 0.03 => 4

}
