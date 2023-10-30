"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import google.protobuf.descriptor
import google.protobuf.message
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

@typing_extensions.final
class PlayRequest(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    NAME_FIELD_NUMBER: builtins.int
    FILE_PATH_FIELD_NUMBER: builtins.int
    LOOP_COUNT_FIELD_NUMBER: builtins.int
    MAXTIME_MS_FIELD_NUMBER: builtins.int
    FADEIN_MS_FIELD_NUMBER: builtins.int
    BLOCK_FIELD_NUMBER: builtins.int
    name: builtins.str
    file_path: builtins.str
    loop_count: builtins.int
    maxtime_ms: builtins.int
    fadein_ms: builtins.int
    block: builtins.bool
    def __init__(
        self,
        *,
        name: builtins.str = ...,
        file_path: builtins.str = ...,
        loop_count: builtins.int = ...,
        maxtime_ms: builtins.int = ...,
        fadein_ms: builtins.int = ...,
        block: builtins.bool = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["block", b"block", "fadein_ms", b"fadein_ms", "file_path", b"file_path", "loop_count", b"loop_count", "maxtime_ms", b"maxtime_ms", "name", b"name"]) -> None: ...

global___PlayRequest = PlayRequest

@typing_extensions.final
class PlayResponse(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    TEXT_FIELD_NUMBER: builtins.int
    text: builtins.str
    def __init__(
        self,
        *,
        text: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["text", b"text"]) -> None: ...

global___PlayResponse = PlayResponse

@typing_extensions.final
class StopRequest(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    NAME_FIELD_NUMBER: builtins.int
    name: builtins.str
    def __init__(
        self,
        *,
        name: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["name", b"name"]) -> None: ...

global___StopRequest = StopRequest

@typing_extensions.final
class StopResponse(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    TEXT_FIELD_NUMBER: builtins.int
    text: builtins.str
    def __init__(
        self,
        *,
        text: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["text", b"text"]) -> None: ...

global___StopResponse = StopResponse