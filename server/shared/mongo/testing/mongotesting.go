package mongotesting

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

const (
	image         = "mongo:4.4"
	containerPort = "27017/tcp"
)

// RunWithMongoInDocker runs the tests with
// a mongodb instance in a docker container.
func RunWithMongoInDocker(m *testing.M, mongoURI *string) int {
	c, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	defer c.Close()
	fmt.Println(c.ClientVersion())
	ctx := context.Background()
	resp, err := c.ContainerCreate(ctx,
		&container.Config{
			Image: image,
			ExposedPorts: nat.PortSet{
				containerPort: {},
			},
		}, &container.HostConfig{
			PortBindings: nat.PortMap{
				containerPort: []nat.PortBinding{
					{
						HostIP:   "127.0.0.1",
						HostPort: "0",
					},
				},
			},
		},
		nil, nil, "")
	if err != nil {
		panic(err)
	}
	containerID := resp.ID
	defer func() {
		err = c.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{Force: true})
		if err != nil {
			//log.Fatalf("error removing container: %v\n", err)
			panic(err)
		}
	}()
	err = c.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("container started")
	time.Sleep(5 * time.Second)
	inspRes, err := c.ContainerInspect(ctx, containerID)
	if err != nil {
		panic(err)
	}
	hostPort := inspRes.NetworkSettings.Ports[containerPort][0]
	*mongoURI = fmt.Sprintf("mongodb://%s:%s/?readPreference=primary&ssl=false", hostPort.HostIP, hostPort.HostPort)
	return m.Run()
}
