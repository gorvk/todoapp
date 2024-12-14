import { Routes, Route, Outlet, Navigate } from "react-router-dom";
import { Auth } from "../../pages/Auth";
import { Dashboard } from "../../pages/Dashboard";
import { Welcome } from "../../pages/Welcome";
import { useSelector } from "react-redux";
import { AppState } from "../../state/store";

export const AppRoutes = () => {
  const currentUser = useSelector((state: AppState) => state.currentUser);
  return (
    <Routes>
      <Route path="/auth" element={<Auth isCallback={false} />} />
      <Route path="/callback" element={<Auth isCallback={true} />} />
      <Route path="/" element={<Welcome />} />
      <Route element={currentUser ? <Outlet /> : <Navigate to="/auth" />}>
        <Route path="/dashboard" element={<Dashboard />} />
      </Route>
    </Routes>
  );
};
