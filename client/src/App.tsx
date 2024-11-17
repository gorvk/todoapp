import { useEffect, useState } from "react";
import { Todo } from "./Todo";
import { TodoData } from "./models/types";
import { getAllTodos, updateTodo, deleteTodo } from "./svc/todo";

function App() {
  const [todoData, setTodoData] = useState<TodoData[]>([]);
  useEffect(() => {
    getAllTodosData();
  }, []);

  async function getAllTodosData() {
    const data = await getAllTodos();
    setTodoData(data);
  }

  async function setIsCompleted(todo: TodoData) {
    todo.isCompleted = !todo.isCompleted;
    const mappedData = todoData.map((data) => {
      if (data.id === todo.id) {
        data = todo;
      }
      return data;
    });
    await updateTodo(todo);
    setTodoData(mappedData);
  }

  async function deleteTodoData(id: string) {
    await deleteTodo(id);
    await getAllTodosData();
  }

  return (
    <>
      {todoData.map((data) => (
        <Todo
          key={data.id}
          data={data}
          setIsCompleted={setIsCompleted}
          deleteTodoData={deleteTodoData}
        ></Todo>
      ))}
    </>
  );
}

export default App;
