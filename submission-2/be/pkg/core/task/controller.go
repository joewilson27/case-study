package task

import (
	"fmt"
	"strconv"

	"be/common/model"
	"be/common/response"
	"be/config/db"

	"github.com/gofiber/fiber/v2"
)

func AddTask(c *fiber.Ctx) error {
	var newTask TaskAdd

	if err := c.BodyParser(&newTask); err != nil {
		dataReturnError := model.AppResponse{
			Data: nil,
			Meta: model.Meta{Message: err.Error()},
		}
		return c.Status(fiber.StatusBadRequest).JSON(dataReturnError)
	}

	svc := Service{
		DB: db.PG,
	}

	resp, errAdd := svc.Add(newTask)
	if errAdd != nil {
		dataReturnError := model.AppResponse{
			Data: nil,
			Meta: model.Meta{Message: response.DataAddedFailed + ": " + errAdd.Error()},
		}
		return c.Status(fiber.StatusBadRequest).JSON(dataReturnError)
	}

	return c.Status(fiber.StatusAccepted).JSON(model.AppResponse{
		Data: resp,
		Meta: model.Meta{Message: response.DataAddedSuccessfully},
	})
}

func GetTasks(c *fiber.Ctx) error {

	svc := Service{
		DB:  db.PG,
		Ctx: c,
	}

	result, err := svc.GetTasks()
	if err != nil {
		dataReturnError := model.AppResponse{
			Data: nil,
			Meta: model.Meta{Message: response.GetDataFailed + " : " + err.Error()},
		}
		return c.Status(fiber.StatusBadRequest).JSON(dataReturnError)
	}

	return c.Status(fiber.StatusOK).JSON(model.AppResponse{
		Data: result,
		Meta: model.Meta{Message: response.GetDataSuccessfully},
	})
}

func GetTaskById(c *fiber.Ctx) error {

	svc := Service{
		DB:  db.PG,
		Ctx: c,
	}

	result, err := svc.GetTaskById()
	if err != nil {
		dataReturnError := model.AppResponse{
			Data: nil,
			Meta: model.Meta{Message: response.DataNotFound + " " + err.Error()},
		}

		return c.Status(fiber.StatusNotFound).JSON(dataReturnError)
	}

	return c.Status(fiber.StatusOK).JSON(model.AppResponse{
		Data: result,
		Meta: model.Meta{Message: response.StatusSuccess},
	})
}

func DeleteTask(c *fiber.Ctx) error {
	idp := c.Params("id")
	id, _ := strconv.Atoi(idp)

	svc := Service{
		DB: db.PG,
	}
	fmt.Println(idp)
	if idp == "" {
		dataReturnError := model.AppResponse{
			Meta: model.Meta{Message: response.FailToDelete},
		}
		return c.Status(fiber.StatusBadRequest).JSON(dataReturnError)
	}

	if err := svc.DeleteByID(uint(id)); err != nil {
		dataReturnError := model.AppResponse{
			Meta: model.Meta{Message: response.FailToDelete + " " + err.Error()},
		}
		return c.Status(fiber.StatusBadRequest).JSON(dataReturnError)
	}

	return c.Status(fiber.StatusOK).JSON(model.AppResponse{
		Meta: model.Meta{Message: response.SuccessDeleted},
	})
}

func UpdateTask(c *fiber.Ctx) error {
	var dataInput TaskAdd
	idp := c.Params("id")
	id, _ := strconv.Atoi(idp)

	if idp == "" {
		dataReturnError := model.AppResponse{
			Meta: model.Meta{Message: response.FailedToUpdate},
		}
		return c.Status(fiber.StatusBadRequest).JSON(dataReturnError)
	}

	if err := c.BodyParser(&dataInput); err != nil {
		dataReturnError := model.AppResponse{
			Meta: model.Meta{Message: response.ErrorInput + " " + err.Error()},
		}
		return c.Status(fiber.StatusBadRequest).JSON(dataReturnError)
	}

	svc := Service{
		DB:  db.PG,
		Ctx: c,
	}

	if err := svc.UpdateByID(uint(id), dataInput); err != nil {
		dataReturnError := model.AppResponse{
			Meta: model.Meta{Message: response.FailedToUpdate + " " + err.Error()},
		}
		return c.Status(fiber.StatusBadRequest).JSON(dataReturnError)
	}

	return c.Status(fiber.StatusOK).JSON(model.AppResponse{
		Data: dataInput,
		Meta: model.Meta{Message: response.SuccessUpdate},
	})
}

func GetTaskComplete(c *fiber.Ctx) error {

	svc := Service{
		DB:  db.PG,
		Ctx: c,
	}

	result, err := svc.GetTasksComplete()
	if err != nil {
		dataReturnError := model.AppResponse{
			Data: nil,
			Meta: model.Meta{Message: response.GetDataFailed + " : " + err.Error()},
		}
		return c.Status(fiber.StatusBadRequest).JSON(dataReturnError)
	}

	return c.Status(fiber.StatusOK).JSON(model.AppResponse{
		Data: result,
		Meta: model.Meta{Message: response.GetDataSuccessfully},
	})

}
