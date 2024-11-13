package random

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func randomHandler(c *fiber.Ctx) error {
	var data map[string]any
	var data_array []map[string]any
	c.BodyParser(&data)
	var options Options = Default

	// Get config from headers
	h := c.GetReqHeaders()

	// Config loop
	if len(h["Loop"]) != 0 {
		var s = h["Loop"]
		options.loop, _ = strconv.Atoi(s[0])
		delete(data, "options")
	}

	// Commands
	for i := 0; i < options.loop; i++ {
		for k, v := range data {
			switch v {
			case "rand:name":
				data[k] = rName()
			case "rand:phone":
				data[k] = rPhone()
			case "rand:username":
				data[k] = rUsername()
			case "rand:email":
				data[k] = rEmail()
			case "rand:score":
				data[k] = rScore()
			case "rand:priority":
				data[k] = rPriority()
			}
		}
		if options.loop > 1 {
			data_array = append(data_array, data)
		}
	}

	if options.loop > 1 {
		return c.JSON(data_array)
	} else {
		return c.JSON(data)
	}
}
