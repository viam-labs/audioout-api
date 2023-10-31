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
	fmt.Println("err", err)
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	audio.Play(context.Background(), path+"/test/munch_2.wav", 10, 0, 0, true)
}
