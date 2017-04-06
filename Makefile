# The MIT License
#
# Copyright (c) 2016 Sebastian Florek
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in
# all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
# THE SOFTWARE.

GO := $(shell command -v go 2> /dev/null)
XGO := $(shell command -v xgo 2> /dev/null)
DOCKER := $(shell command -v docker 2> /dev/null)

DOCKER_HUB = floreks

prepare:
	@echo "Creating build dir..."
	@mkdir -p ./build

check_docker:
ifndef DOCKER
	$(error "Could not find docker.")
endif

check_go:
ifndef GO
	$(error "Could not find GO compiler.")
endif

check_xgo:
ifndef XGO
	$(error "Could not find XGO compiler.")
endif

all: prepare check_go check_xgo amd64 arm

amd64: prepare check_go
	@echo "Building app for amd64 arch"
	@go build -o build/sensor-mqtt-client

arm: prepare check_xgo
	@echo "Building app for arm arch"
	@xgo -targets linux/arm-6 -out build/sensor-mqtt-client .
	@mv build/sensor-mqtt-client-linux-arm-6 build/sensor-mqtt-client-arm

clean:
	echo "Cleaning up..."
	rm -rf ./build

build_docker_arm: check_xgo
	docker build -t $(DOCKER_HUB)/sensor-mqtt-client-arm .

deploy_arm: build_docker_arm
	docker push $(DOCKER_HUB)/sensor-mqtt-client-arm