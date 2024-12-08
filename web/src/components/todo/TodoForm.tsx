import { TodoData } from "../../models/types";
import cancel from "../../assets/cancel.svg";
import done from "../../assets/done.svg";
import { useRef } from "react";

export function TodoForm(props: {
  data: TodoData;
  closeHandler: () => void;
  submitHandler: (
    checkRef: React.RefObject<HTMLInputElement>,
    data?: TodoData,
  ) => Promise<void>;
}) {
  const { data, closeHandler, submitHandler } = props;
  const checkRef = useRef<HTMLInputElement>(null);

  async function editClickHandler() {
    await submitHandler(checkRef, data);
    closeHandler();
  }

  return (
    <div className="flex bg-zinc-800 shadow-md shadow-black p-4 m-4 w-1/5">
      <input
        ref={checkRef}
        className="w-full"
        type="text"
        defaultValue={data.title}
      />
      <div className="flex ml-4 gap-4">
        <img onClick={editClickHandler} src={done} width="12" />
        <img onClick={() => closeHandler()} src={cancel} width="12" />
      </div>
    </div>
  );
}
