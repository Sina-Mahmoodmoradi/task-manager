package models

import "github.com/Sina-Mahmoodmoradi/task-manager/internal/domain"

// ----------- User Mappers ------------

func ToDomainUser(u User) domain.User {
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

func FromDomainUser(u domain.User) User {
    tasks := make([]Task, len(u.Tasks))
    for i, t := range u.Tasks {
        tasks[i] = FromDomainTask(t)
    }

    return User{
        Username: u.Username,
        Password: u.Password,
        Tasks:    tasks,
    }
}

// ----------- Task Mappers ------------

func ToDomainTask(t Task) domain.Task {
    return domain.Task{
        ID:     t.ID,
        Title:  t.Title,
        Done:   t.Done,
        UserID: t.UserID,
    }
}

func FromDomainTask(t domain.Task) Task {
    return Task{
        Title:  t.Title,
        Done:   t.Done,
        UserID: t.UserID,
    }
}
