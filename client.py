# test, will play sample audio file looped 10 times
import asyncio
import os

from src.audioout_python import Audioout

from viam import logging
from viam.robot.client import RobotClient
from viam.rpc.dial import Credentials, DialOptions

# these must be set, you can get them from your robot's 'CODE SAMPLE' tab
robot_address = os.getenv('ROBOT_ADDRESS') or ''
robot_api_key = os.getenv('ROBOT_API_KEY') or ''
robot_api_key_id = os.getenv('ROBOT_API_KEY_ID') or ''

async def connect():
    opts = RobotClient.Options.with_api_key(
      api_key=robot_api_key,
      api_key_id=robot_api_key_id
    )
    return await RobotClient.at_address(robot_address, opts)


async def main():
    robot = await connect()

    #print("Resources:")
    #print(robot.resource_names)

    ao = Audioout.from_robot(robot, name="audio")

    text = await ao.play(os.getcwd() + "/test/munch_2.wav", 10, 0, 0)
    print(f"Played '{text}'")
    
    await robot.close()


if __name__ == "__main__":
        asyncio.run(main())