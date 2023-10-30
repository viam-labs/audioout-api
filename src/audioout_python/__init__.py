"""
This file registers the model with the Python SDK.
"""

from viam.resource.registry import Registry, ResourceCreatorRegistration, ResourceRegistration

from .api import AudiooutClient, AudiooutRPCService, Audioout

Registry.register_subtype(ResourceRegistration(Audioout, AudiooutRPCService, lambda name, channel: AudiooutClient(name, channel)))
