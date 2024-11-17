import deleteLogo from "./assets/delete.svg";
import editLogo from "./assets/edit.svg";
import { TodoData } from "./models/types";
export function Todo(props: {
  data: TodoData;
  setIsCompleted: (todo: TodoData) => void;
  deleteTodoData: (id: string) => void;
}) {
  const { data, setIsCompleted, deleteTodoData } = props;
  return (
    <div className="bg-zinc-800 shadow-md shadow-black p-4 m-4 flex justify-between w-1/5">
      <p
        className={"text-white text-lg " + (data.isCompleted && "line-through")}
      >
        {data.title}
      </p>
      <input
        type="checkbox"
        checked={data.isCompleted}
        onChange={() => setIsCompleted(data)}
      />
      <img src={editLogo} width="12" />
      <img onClick={() => deleteTodoData(data.id)} src={deleteLogo} width="12" />
    </div>
  );
}
