docker build -t emulator . 
docker run --ipc=host -v /tmp:/tmp emulator