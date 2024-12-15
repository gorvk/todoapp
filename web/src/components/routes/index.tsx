import { Routes, Route, Navigate } from "react-router-dom";
import { Dashboard } from "../../pages/Dashboard";
import { Welcome } from "../../pages/Welcome";
import { useSelector } from "react-redux";
import { AppState } from "../../state/store";
import { Auth } from "../../pages/Auth";

export const AppRoutes = () => {
  const currentUser = useSelector((state: AppState) => state.currentUser);
  return (
    <Routes>
      <Route path="/" element={<Welcome />} />
      <Route
        path="/auth"
        element={!currentUser ? <Auth /> : <Navigate to="/dashboard" />}
      />
      <Route
        path="/dashboard"
        element={currentUser ? <Dashboard /> : <Navigate to="/auth" />}
      />
    </Routes>
  );
};
