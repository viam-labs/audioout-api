# audioout-api

Proto API and grpc bindings for audioout

*audioout-api* provides Proto API and grpc bindings for audio output capabilities

## API

The audioout resource implements the following API:

### play(file_path=*string*, loop_count=*int*(0), maxtime_ms=*int*(0), fadein_ms=*int*(0), block=*bool*)

The *play()* command takes:

* file_path: The audio file on device to play
* loop_count: How many times to play the audio file.  0 means once, -1 will loop infinitely (until stop() is called). Default 0.
* maxtime_ms: How long to play the audio for.  0 means no maxtime. Note that some file types like .wav do not support time indexing so this will fail. Default 0.
* fadein_ms: If non-zero, will make the sound start playing at 0 volume and fade up to full volume over the time given. The sample may end before the fade-in is complete.  Default 0.
* block: If False, will play sound async.  If true, will not return until sound is complete.  Default False.

This method returns a string response, which is the file_path that was passed in to the *play()* request.

### stop()

The *stop()* command will stop sound playback on any active channels.

If successful, will return the string "OK".

## Using audioout-api with the Python SDK

Because this module uses a custom protobuf-based API, you must include this project in your client code.  One way to do this is to include it in your requirements.txt as follows:

```
audioout_api @ git+https://github.com/viam-labs/audioout-api.git@main
```

You can now import and use it in your code as follows:

```
from audioout_python import Audioout
ao = Audioout.from_robot(robot, name="audioout")
ao.play(...)
```

See client.py for an example.

## Using audioout with the Golang SDK

Because this module uses a custom protobuf-based API, you must import and use in your client code as follows:

``` go
import audioout "github.com/viam-labs/audioout-api/src/audioout_go"

audio, err := audioout.FromRobot(robot, "ao")
fmt.Println("err", err)
audio.Play(context.Background(), "test/munch_2.wav", 10, 0, 0, true)
```

See client.go for an example.

## Building

To rebuild the GRPC bindings, run:

``` bash
make generate
```

Then, in `src/audioout_python/grpc/audioout_grpc.py change:

``` python
import audioout_pb2
```

to:

``` python
from . import audioout_pb2
```

Then, update the version in pyproject.toml
