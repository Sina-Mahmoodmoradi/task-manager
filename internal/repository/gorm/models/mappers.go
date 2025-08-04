package models

import "github.com/Sina-Mahmoodmoradi/task-manager/internal/domain"

// ----------- User Mappers ------------

func ToDomainUser(u UserModel) domain.User {
    tasks := make([]domain.Task, len(u.Tasks))
    for i, t := range u.Tasks {
        tasks[i] = ToDomainTask(t)
    }

    return domain.User{
        ID:       u.ID,
        Username: u.Username,
        Password: u.Password,
        Tasks:    tasks,
    }
}

func FromDomainUser(u domain.User) UserModel {
    tasks := make([]TaskModel, len(u.Tasks))
    for i, t := range u.Tasks {
        tasks[i] = FromDomainTask(t)
    }

    return UserModel{
        Username: u.Username,
        Password: u.Password,
        Tasks:    tasks,
    }
}

// ----------- Task Mappers ------------

func ToDomainTask(t TaskModel) domain.Task {
    return domain.Task{
        ID:     t.ID,
        Title:  t.Title,
        Done:   t.Done,
        UserID: t.UserID,
    }
}

func FromDomainTask(t domain.Task) TaskModel {
    return TaskModel{
        Title:  t.Title,
        Done:   t.Done,
        UserID: t.UserID,
    }
}
