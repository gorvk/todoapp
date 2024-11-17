import { useEffect, useState } from "react";
import { Todo } from "./components/todo/Todo";
import { TodoData } from "./models/types";
import { getAllTodos, updateTodo, deleteTodo, addTodo } from "./svc/todo";
import add from "./assets/add.svg";
import { TodoForm } from "./components/todo/TodoForm";

function App() {
  const [todoData, setTodoData] = useState<TodoData[]>([]);
  const [isAddEnabled, setIsAddEnabled] = useState<boolean>(false);
  useEffect(() => {
    getAllTodosData();
  }, []);

  async function getAllTodosData() {
    const data = await getAllTodos();
    setTodoData(data);
  }

  async function setIsCompleted(todo: TodoData) {
    const isCompleted = !todo.isCompleted;
    const toggledTodo = { ...todo, isCompleted };
    const mappedData = todoData.map((data) => {
      if (data.id === toggledTodo.id) {
        data = toggledTodo;
      }
      return data;
    });
    await updateTodo(toggledTodo);
    setTodoData(mappedData);
  }

  async function deleteTodoData(id: string) {
    await deleteTodo(id);
    await getAllTodosData();
  }

  async function editTodoData(
    checkRef: React.RefObject<HTMLInputElement>,
    data?: TodoData
  ) {
    if (checkRef.current && data) {
      const title = checkRef.current.value;
      const editedTodo = { ...data, title };
      await updateTodo(editedTodo);
      await getAllTodosData();
    }
  }

  async function Add(checkRef: React.RefObject<HTMLInputElement>) {
    if (checkRef.current) {
      const title = checkRef.current.value;
      const data = { title } as TodoData;
      await addTodo(data);
      await getAllTodosData();
    }
  }

  return (
    <>
      {todoData.map((data) => (
        <Todo
          key={data.id}
          data={data}
          setIsCompleted={setIsCompleted}
          deleteTodoData={deleteTodoData}
          editTodoData={editTodoData}
        ></Todo>
      ))}
      {isAddEnabled ? (
        <TodoForm
          closeHandler={() => setIsAddEnabled(false)}
          data={{} as TodoData}
          submitHandler={Add}
        />
      ) : (
        <img
          className="m-4"
          src={add}
          width="12"
          onClick={() => setIsAddEnabled(true)}
        />
      )}
    </>
  );
}

export default App;
