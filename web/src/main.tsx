import { createRoot } from "react-dom/client";
import App from "./App.tsx";
import { Provider } from "react-redux";
import { store } from "./state/store.ts";

createRoot(document.getElementById("root") as HTMLElement).render(
  <Provider store={store}>
      <App />
  </Provider>
);
