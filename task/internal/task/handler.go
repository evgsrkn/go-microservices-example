package task

import (
	"go-microservices-example/task/pkg/pb"
	"task/internal/task/model"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type (
	taskHandler struct {
		service IService
		log     *zap.Logger
		server  *grpc.Server
	}

	IHandler interface {
		GetAllTasks(_ *pb.Empty) (*pb.TasksList, error)
		GetTaskById(task *pb.TaskWithID) (*pb.Task, error)
		UpdateTask(task *pb.Task) (*pb.Task, error)
		DeleteTask(task *pb.TaskWithID) (*pb.Empty, error)
		CreateTask(task *pb.Task) (*pb.Empty, error)
	}
)

func NewHandler(service IService, log *zap.Logger, server *grpc.Server) IHandler {
	return &taskHandler{service, log, server}
}

func (h *taskHandler) GetAllTasks(_ *pb.Empty) (*pb.TasksList, error) {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	var pbtasks pb.TasksList
	for _, task := range tasks {
		pbtasks.Tasks = append(pbtasks.Tasks, &pb.Task{
			Id:          int64(task.ID),
			Name:        task.Name,
			Description: task.Description,
			ProjectId:   int64(task.Project_id),
			UserId:      int64(task.User_id),
			Status:      task.Status,
		})
	}

	return &pbtasks, nil
}

func (h *taskHandler) GetTaskById(request *pb.TaskWithID) (*pb.Task, error) {
	task, err := h.service.GetTaskById(int(request.Id))
	if err != nil {
		return nil, err
	}

	return &pb.Task{
		Id:          int64(task.ID),
		Name:        task.Name,
		Description: task.Description,
		ProjectId:   int64(task.Project_id),
		UserId:      int64(task.User_id),
		Status:      task.Status,
	}, nil
}

func (h *taskHandler) UpdateTask(request *pb.Task) (*pb.Task, error) {
	task, err := h.service.UpdateTask(&model.Task{
		ID:          int(request.Id),
		Name:        request.Name,
		Description: request.Description,
		Project_id:  int(request.ProjectId),
		User_id:     int(request.UserId),
		Status:      request.Status,
	})
	if err != nil {
		return nil, err
	}

	return &pb.Task{
		Id:          int64(task.ID),
		Name:        request.Name,
		Description: task.Description,
		ProjectId:   int64(task.Project_id),
		UserId:      int64(task.User_id),
		Status:      task.Status,
	}, nil
}

func (h *taskHandler) DeleteTask(task *pb.TaskWithID) (*pb.Empty, error) {
	if err := h.service.DeleteTask(int(task.Id)); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (h *taskHandler) CreateTask(request *pb.Task) (*pb.Empty, error) {
	err := h.service.CreateTask(&model.Task{
		Name:        request.Name,
		Description: request.Description,
		Project_id:  int(request.ProjectId),
		User_id:     int(request.UserId),
		Status:      request.Status,
	})

	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
