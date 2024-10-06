package todo

import "time"

type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TodoList struct {
	Todos []Todo `json:"todos"`
}

func NewTodoList() *TodoList {
	tdl := &TodoList{
		Todos: []Todo{},
	}
	tdl.LoadTodos()
	return tdl
}

func NewTodo(title string, description string) *Todo {
	return &Todo{
		ID:          -1,
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
func (tl *TodoList) Add(todo Todo) {
	todo.ID = len(tl.Todos) + 1
	tl.Todos = append(tl.Todos, todo)
}

func (tl *TodoList) Remove(id int) {
	for i, todo := range tl.Todos {
		if todo.ID == id {
			tl.Todos = append(tl.Todos[:i], tl.Todos[i+1:]...)
			break
		}
	}
	tl.SaveTodo()
}

func (tl *TodoList) Update(todo Todo) {
	for i, t := range tl.Todos {
		if t.ID == todo.ID {
			tl.Todos[i] = todo
			break
		}
	}
	tl.SaveTodo()
}

func (tl *TodoList) Get(id int) *Todo {
	for _, t := range tl.Todos {
		if t.ID == id {
			return &t
		}
	}
	return nil
}

func (tl *TodoList) GetAllTitles() []string {
	var todos []string
	for _, t := range tl.Todos {
		todos = append(todos, t.Title)
	}
	return todos
}
