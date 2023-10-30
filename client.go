package main

import (
	"context"
	"fmt"
	"os"

	"github.com/edaniels/golog"
	audioout "github.com/viam-labs/audioout-api/src/audioout_go"
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

	audio, err := audioout.FromRobot(robot, "ao")
	audio.Play(context.Background(), "test/munch_2.wav", 0, 10, 0, true)
	fmt.Println("err", err)
}
