import {
  BrowserRouter,
} from "react-router-dom";
import { AppRoutes } from "./components/routes";

function App() {
  return (
    <BrowserRouter>
      <AppRoutes></AppRoutes>
    </BrowserRouter>
  );
}

export default App;
