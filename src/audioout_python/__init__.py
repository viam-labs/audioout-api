"""
This file registers the model with the Python SDK.
"""

from viam.resource.registry import Registry, ResourceRegistration

from .api import AudiooutClient, AudiooutRPCService, Audioout

Registry.register_api(
    ResourceRegistration(
        Audioout,
        AudiooutRPCService,
        lambda name, channel: AudiooutClient(name, channel),
    )
)
