# Machine Learning Distance Line Protection Zone

Implementing Machine Learning Multi-class Classification Algorithms for obtaining the Micom P543 distance relay protection curve in transmission lines with Deep Neural Network and Random Forest

For gathering data requirements, first the distance function of Micom P543 relay was tested with Vebko AMT105 relay tester and the results were given as input to the Deep Neural Network and Random Forest to get the characteristic distance curve

## Features
  - Using <b>Tensorflow</b> to build a Multi Classification Algorithm with Deep Neural Network model
  - Using <b>Scikit-Learn</b> to build a a Multi Classification Algorithm with Random Forest
  - Deep Neural Network Accuracy = 98% 
  - Random Forest Accuracy = 96%
  - Using <a href="https://www.se.com/uk/en/product-range/60747-micom-p54x/#overview" target="_blank"><b>Schneider Electric Micom P543 Relay</b></a> testing by <a href="https://vebko.org/en/Default.aspx" target="_blank"><b>Vebko AMT105</b></a> relay tester for create dataset
  - Converting the <b>Tensorflow</b> model to tflite for running on Embedded Board ARM Architecture
  - Using <b>Golang</b> TFLite to be able to easily run tflite model
  - Running on <a href="https://www.xilinx.com/products/silicon-devices/soc/zynq-7000.html" target="_blank"><b>Xilinx Zynq-7020</b></a> Embedded Board
  - Usable via <b>Docker</b> file
  
## Installation

#### First you need install TensorFlow for C

1) Install bazel
```bash
curl https://bazel.build/bazel-release.pub.gpg | sudo apt-key add -
echo "deb [arch=amd64] https://storage.googleapis.com/bazel-apt stable jdk1.8" | sudo tee /etc/apt/sources.list.d/bazel.list
sudo apt update && sudo apt install bazel
sudo apt install openjdk-11-jdk
```

2) Build tensorflowlite c lib from source
```bash
cd ~/workspace
git clone https://github.com/tensorflow/tensorflow.git && cd tensorflow
./configure
bazel build --config opt --config monolithic --define tflite_with_xnnpack=false //tensorflow/lite:libtensorflowlite.so
bazel build --config opt --config monolithic --define tflite_with_xnnpack=false //tensorflow/lite/c:libtensorflowlite_c.so

# Check status
file bazel-bin/tensorflow/lite/c/libtensorflowlite_c.so
# ELF 64-bit LSB shared object, x86-64
```

3) Build go-tflite
```bash
export CGO_LDFLAGS=-L$HOME/workspace/tensorflow/bazel-bin/tensorflow/lite/c
export CGO_CFLAGS=-I$HOME/workspace/tensorflow/
```

## Build

For Linux/MacOs amd64:

```bash
  export CGO_LDFLAGS=-L$HOME/workspace/tensorflow/bazel-bin/tensorflow/lite/c

  go build main.go
```

For xilinx Zynq-7020 (ARM-based computers):

```bash
  sudo apt-get install gcc-arm-linux-gnueabihf

  export CGO_LDFLAGS=-L$HOME/workspace/tensorflow/bazel-bin/tensorflow/lite/c
  
  CGO_ENABLED=1 GOOS=linux GOARCH=arm CC=arm-linux-gnueabihf-gcc go build -o main
```

## Running

This running for ubuntu/MacOs amd64:

```bash
  ./main
```

This running for xilinx Zynq-7020 (ARM-based computers):

```bash
  export LD_LIBRARY_PATH=./arm
  
  ./main
```

## Running with Docker

First of all, clone and the repo then run
```bash
  docker build -t dnn .
```

After pulling and building the image, You can get the result like this

```bash
  docker run --rm -t distance ./main
```

Or you can go to the container for running it manually like this

```bash
  docker run -it distance
```

## More Info

#### Micom P543 Relay testing by vebko AMT105
![Graph](https://github.com/taherfattahi/dnn-distance-line-protection-zone/blob/master/images/micom_p543_relay_testing_by_vebko.jpg)

#### Distance Line Protection Zone in the AMPro software
![Graph](https://github.com/taherfattahi/dnn-distance-line-protection-zone/blob/master/images/distance_line_protection_zone.png)

#### Graph of the Deep Neural Network
![Graph](https://github.com/taherfattahi/dnn-distance-line-protection-zone/blob/master/images/graph.png)

#### Model Accuracy Plot
![Graph](https://github.com/taherfattahi/dnn-distance-line-protection-zone/blob/master/images/model_accuracy.png)

#### Model Loss Plot
![Graph](https://github.com/taherfattahi/dnn-distance-line-protection-zone/blob/master/images/model_loss.png)


## Note:
 If you had issue and got `standard_init_linux.go:211: exec user process caused "exec format error` error, try [this solution](https://www.stereolabs.com/docs/docker/building-arm-container-on-x86/).

## Collaborators

- [Nima Akbarzade](https://www.github.com/iw4p) - [Mohsen Shahsavari](https://github.com/mohsenshahsavari)


## License

[MIT](https://choosealicense.com/licenses/mit/)

