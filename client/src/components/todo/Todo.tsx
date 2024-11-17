import { useState } from "react";
import trash from "../../assets/trash.svg";
import edit from "../../assets/edit.svg";
import { TodoData } from "../../models/types";
import { TodoForm } from "./TodoForm";
export function Todo(props: {
  data: TodoData;
  setIsCompleted: (todo: TodoData) => Promise<void>;
  deleteTodoData: (id: string) => Promise<void>;
  editTodoData: (
    checkRef: React.RefObject<HTMLInputElement>,
    data?: TodoData,
  ) => Promise<void>;
}) {
  const { data, setIsCompleted, deleteTodoData, editTodoData } = props;
  const [isEditEnabled, setIsEditEnabled] = useState<boolean>(false);

  return (
    <>
      {isEditEnabled ? (
        <TodoForm
          data={data}
          closeHandler={() => setIsEditEnabled(false)}
          submitHandler={editTodoData}
        ></TodoForm>
      ) : (
        <div className="bg-zinc-800 shadow-md shadow-black p-4 m-4 flex justify-between w-1/5">
          <div className="flex w-1/2">
            <p
              className={
                "text-white text-lg " + (data.isCompleted && "line-through")
              }
            >
              {data.title}
            </p>
          </div>
          <div className="flex w-1/2 ml-8 justify-between">
            <input
              type="checkbox"
              defaultChecked={data.isCompleted}
              onChange={() => setIsCompleted(data)}
            />
            <img onClick={() => setIsEditEnabled(true)} src={edit} width="12" />
            <img
              onClick={() => deleteTodoData(data.id)}
              src={trash}
              width="12"
            />
          </div>
        </div>
      )}
    </>
  );
}
