package todos

type Service interface {
	FindAll() ([]Todo, error)
	FindByID(id int) (Todo, error)
	Save(input CreateTodosInput) (Todo, error)
	Update(input UpdateTodosInput) (Todo, error)
	Done(input UpdateTodosInput) (Todo, error)
	Delete(id int) (Todo, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Todo, error) {
	todos, err := s.repository.FindAll()
	if err != nil {
		return todos, err
	}

	return todos, nil
}

func (s *service) Save(input CreateTodosInput) (Todo, error) {
	todo := Todo{}
	todo.Task = input.Task
	todo.Asignee = input.Asignee
	todo.Deadline = input.Deadline
	todo.Done = input.Done

	newTodo, err := s.repository.Save(todo)
	if err != nil {
		return newTodo, err
	}

	return newTodo, nil
}

func (s *service) Update(input UpdateTodosInput) (Todo, error) {
	todo, err := s.repository.FindByID(input.ID)
	if err != nil {
		return todo, err
	}
	todo.Task = input.Task
	todo.Asignee = input.Asignee
	todo.Deadline = input.Deadline

	updatedTodo, err := s.repository.Update(todo)
	if err != nil {
		return updatedTodo, err
	}

	return updatedTodo, nil
}

func (s *service) Done(input UpdateTodosInput) (Todo, error) {
	todo, err := s.repository.FindByID(input.ID)
	if err != nil {
		return todo, err
	}
	todo.Task = input.Task
	todo.Asignee = input.Asignee
	todo.Deadline = input.Deadline
	todo.Done = input.Done

	updatedTodo, err := s.repository.Update(todo)
	if err != nil {
		return updatedTodo, err
	}

	return updatedTodo, nil
}

func (s *service) Delete(id int) (Todo, error) {
	deleteTodo, err := s.repository.Delete(id)
	if err != nil {
		return deleteTodo, err
	}

	return deleteTodo, nil
}

func (s *service) FindByID(id int) (Todo, error) {
	findTodo, err := s.repository.FindByID(id)
	if err != nil {
		return findTodo, err
	}

	return findTodo, nil
}
