IDENT=org.golang.todo.osx
ACTIVITY=org.golang.app.GoNativeActivity
PROJECT_DIR=github.com/arzonus/snake

build:
	gomobile build -target=android ${PROJECT_DIR}/cmd/osx
install:
	gomobile install -target=android ${PROJECT_DIR}/cmd/osx
logs:
	adb logcat ${IDENT}:I

start:
	adb shell am start -n ${IDENT}/${ACTIVITY}

stop: force-stop

force-stop:
	adb shell am force-stop ${IDENT}

restart: force-stop start

run: install restart logs