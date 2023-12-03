package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

type Position struct {
    X float32 `json:"x"`
    Y float32 `json:"y"`
}

type Agent struct {
    Name string `json:"name"`
    Coordinates []Position `json:"coordinates"`
}

func getSolvedAgents(c *gin.Context) {
    agents := readCoordsFile("multi_agent.txt");
    c.IndentedJSON(http.StatusOK, agents)
}

// ideally this would be a tcp connection to the simulator
// where the mapper is that would provide the coordinates
// but we did not have enough time :(
func readCoordsFile(fileName string) []Agent{
    var agents = []Agent {}

    // initialize the three agents
    agentMap := make(map[string] []Position)
    agentMap["agent_1"] = []Position {}
    agentMap["agent_2"] = []Position {}
    agentMap["agent_3"] = []Position {}


    // Open the text file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)

        // empty array
        return agents;
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate over each line in the file
	for scanner.Scan() {
        line := scanner.Text()

        split := strings.Split(line, ",")
        coordX, _ := strconv.ParseFloat(split[0], 32);
        coordY, _ := strconv.ParseFloat(split[1], 32);
        name := split[2];

        pos := Position {
            X: float32(coordX),
            Y: float32(coordY),
        }

        agentMap[name] = append(agentMap[name], pos)
    }

    // Check for errors during scanning
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
        return agents
    }

    // create the agents
    // this code bad
    // its 3 am I want ton sleep
    // dont @ me
    agent1 := Agent {
        Name: "agent_1",
        Coordinates: agentMap["agent_1"],
    }
    agent2 := Agent {
        Name: "agent_2",
        Coordinates: agentMap["agent_2"],
    }
    agent3 := Agent {
        Name: "agent_3",
        Coordinates: agentMap["agent_3"],
    }

    agents = append(agents, agent1)
    agents = append(agents, agent2)
    agents = append(agents, agent3)

    return agents
}

func main() {
    router := gin.Default()
    router.Use(cors.Default())
    router.GET("/agents", getSolvedAgents)

    router.Run("localhost:7070")
}

