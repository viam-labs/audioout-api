package main

import (
	"context"
	"os"

	"github.com/edaniels/golog"
	"go.viam.com/rdk/components/audioinput"
	"go.viam.com/rdk/robot/client"
	"go.viam.com/rdk/utils"
	"go.viam.com/utils/rpc"
)

func main() {
	logger := golog.NewDevelopmentLogger("client")
	robot, err := client.New(
		context.Background(),
		os.Getenv("ROBOT_ADDRESS"),
		logger,
		client.WithDialOptions(rpc.WithCredentials(rpc.Credentials{
			Type:    utils.CredentialsTypeRobotLocationSecret,
			Payload: os.Getenv("ROBOT_SECRET"),
		})),
	)
	if err != nil {
		logger.Fatal(err)
	}

	defer robot.Close(context.Background())
	logger.Info("Resources:")
	logger.Info(robot.ResourceNames())

	audio, err := audioinput.FromRobot(robot, "ao")
	audio.Play("test/munch_2.wav")

}
